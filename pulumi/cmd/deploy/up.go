package main

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nitrictech/nitric-provider-template/pulumi/pkg/collection"
	"github.com/nitrictech/nitric-provider-template/pulumi/pkg/secret"
	"github.com/nitrictech/nitric/cloud/aws/deploy/api"
	"github.com/nitrictech/nitric/cloud/aws/deploy/bucket"
	"github.com/nitrictech/nitric/cloud/aws/deploy/exec"
	"github.com/nitrictech/nitric/cloud/aws/deploy/policy"
	"github.com/nitrictech/nitric/cloud/aws/deploy/schedule"
	"github.com/nitrictech/nitric/cloud/aws/deploy/topic"
	"github.com/nitrictech/nitric/cloud/aws/deploy/websocket"
	"github.com/nitrictech/nitric/cloud/common/deploy/image"
	"github.com/nitrictech/nitric/cloud/common/deploy/resources"
	"github.com/nitrictech/nitric/cloud/common/deploy/telemetry"
	v1 "github.com/nitrictech/nitric/core/pkg/api/nitric/v1"
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optrefresh"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Embeds the runtime directly into the deploytime binary
// This way the versions will always match as they're always built and versioned together (as a single artifact)
// This should also help with docker build speeds as the runtime has already been "downloaded"
//
//go:embed runtime
var runtime []byte

func (*DeployServer) Up(request *deploypb.DeployUpRequest, stream deploypb.DeployService_UpServer) error {
	attributes := request.Attributes.AsMap()

	pulumiStack, err := auto.UpsertStackInlineSource(context.TODO(), "TODO", "TODO", func(ctx *pulumi.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				stack := string(debug.Stack())
				err = fmt.Errorf("recovered panic: %+v\n Stack: %s", r, stack)
			}
		}()

		// Deploy all secrets
		secrets := map[string]*secret.Secret{}
		for _, res := range request.Spec.Resources {
			switch c := res.Config.(type) {
			case *deploypb.Resource_Secret:
				secrets[res.Name], err = secret.NewSecret(ctx, res.Name, c.Secret)
				if err != nil {
					return err
				}
			}
		}

		// Deploy all collections
		collections := map[string]*collection.Collection{}
		for _, res := range request.Spec.Resources {
			switch c := res.Config.(type) {
			case *deploypb.Resource_Collection:
				collections[res.Name], err = collection.NewCollection(ctx, res.Name, c.Collection)
				if err != nil {
					return err
				}
			}
		}

		// Deploy all queues
		queues := map[string]*queue.Queue{}
		for _, res := range request.Spec.Resources {
			switch q := res.Config.(type) {
			case *deploypb.Resource_Queue:
				queues[res.Name], err = queue.NewQueue(ctx, res.Name, q.Queue)
				if err != nil {
					return err
				}
			}
		}

		// Deploy all execution units
		execs := map[string]*exec.LambdaExecUnit{}
		for _, res := range request.Spec.Resources {
			switch eu := res.Config.(type) {
			case *deploy.Resource_ExecutionUnit:
				repo, err := ecr.NewRepository(ctx, res.Name, &ecr.RepositoryArgs{
					ForceDelete: pulumi.BoolPtr(true),
					Tags:        pulumi.ToStringMap(common.Tags(stackID, res.Name, resources.ExecutionUnit)),
				})
				if err != nil {
					return err
				}

				if eu.ExecutionUnit.GetImage() == nil {
					return fmt.Errorf("aws provider can only deploy execution with an image source")
				}

				if eu.ExecutionUnit.GetImage().GetUri() == "" {
					return fmt.Errorf("aws provider can only deploy execution with an image source")
				}

				if eu.ExecutionUnit.Type == "" {
					eu.ExecutionUnit.Type = "default"
				}

				typeConfig, hasConfig := config.Config[eu.ExecutionUnit.Type]
				if !hasConfig {
					return fmt.Errorf("could not find config for type %s in %+v", eu.ExecutionUnit.Type, config.Config)
				}

				image, err := image.NewImage(ctx, res.Name, &image.ImageArgs{
					SourceImage:   eu.ExecutionUnit.GetImage().GetUri(),
					RepositoryUrl: repo.RepositoryUrl,
					Server:        pulumi.String(authToken.ProxyEndpoint),
					Username:      pulumi.String(authToken.UserName),
					Password:      pulumi.String(authToken.Password),
					Runtime:       runtime,
					Telemetry: &telemetry.TelemetryConfigArgs{
						TraceSampling: typeConfig.Telemetry,
						TraceName:     "awsxray",
						MetricName:    "awsemf",
						Extensions:    []string{},
					},
				}, pulumi.DependsOn([]pulumi.Resource{repo}))
				if err != nil {
					return err
				}

				if typeConfig.Lambda != nil {
					envMap := pulumi.StringMap{}

					for k, v := range baseExecEnv {
						envMap[k] = v
					}

					for k, v := range eu.ExecutionUnit.Env {
						envMap[k] = pulumi.String(v)
					}

					execs[res.Name], err = exec.NewLambdaExecutionUnit(ctx, res.Name, &exec.LambdaExecUnitArgs{
						DockerImage: image,
						StackID:     stackID,
						Compute:     eu.ExecutionUnit,
						EnvMap:      envMap,
						Client:      lambdaClient,
						Config:      *typeConfig.Lambda,
					})
					execPrincipals[res.Name] = execs[res.Name].Role
				} else {
					return fmt.Errorf("no target execution unit specified for %s", res.Name)
				}

				if err != nil {
					return err
				}
			}
		}
		principals[v1.ResourceType_Function] = execPrincipals

		// Deploy all buckets
		buckets := map[string]*bucket.S3Bucket{}
		for _, res := range request.Spec.Resources {
			switch b := res.Config.(type) {
			case *deploy.Resource_Bucket:
				buckets[res.Name], err = bucket.NewS3Bucket(ctx, res.Name, &bucket.S3BucketArgs{
					// TODO: Calculate stack ID
					StackID: stackID,
					Bucket:  b.Bucket,
				})
				if err != nil {
					return err
				}

				if len(b.Bucket.Notifications) > 0 {
					_, err = bucket.NewS3Notification(ctx, fmt.Sprintf("notification-%s", res.Name), &bucket.S3NotificationArgs{
						StackID:      stackID,
						Location:     details.Region,
						Bucket:       buckets[res.Name],
						Functions:    execs,
						Notification: b.Bucket.Notifications,
					})
					if err != nil {
						return err
					}
				}
			}
		}

		// deploy API Gateways
		// gws := map[string]
		for _, res := range request.Spec.Resources {
			switch t := res.Config.(type) {
			case *deploy.Resource_Api:
				// Deserialize the OpenAPI document

				if t.Api.GetOpenapi() == "" {
					return fmt.Errorf("aws provider can only deploy OpenAPI specs")
				}

				doc := &openapi3.T{}
				err := doc.UnmarshalJSON([]byte(t.Api.GetOpenapi()))
				if err != nil {
					return fmt.Errorf("invalid document suppled for api: %s", res.Name)
				}

				config, _ := config.Apis[res.Name]

				_, err = api.NewAwsApiGateway(ctx, res.Name, &api.AwsApiGatewayArgs{
					LambdaFunctions: execs,
					StackID:         stackID,
					OpenAPISpec:     doc,
					Config:          config,
				})
				if err != nil {
					return err
				}
			}
		}

		// Add all HTTP proxies
		httpProxies := map[string]*api.AwsHttpProxy{}
		for _, res := range request.Spec.Resources {
			switch t := res.Config.(type) {
			case *deploy.Resource_Http:
				fun := execs[t.Http.Target.GetExecutionUnit()]

				httpProxies[res.Name], err = api.NewAwsHttpProxy(ctx, res.Name, &api.AwsHttpProxyArgs{
					StackID:        stackID,
					LambdaFunction: fun,
				})
				if err != nil {
					return err
				}
			}
		}

		// deploy websockets
		websockets := map[string]*websocket.AwsWebsocketApiGateway{}
		for _, res := range request.Spec.Resources {
			switch ws := res.Config.(type) {
			case *deploy.Resource_Websocket:
				websockets[res.Name], err = websocket.NewAwsWebsocketApiGateway(ctx, res.Name, &websocket.AwsWebsocketApiGatewayArgs{
					DefaultTarget:    execs[ws.Websocket.MessageTarget.GetExecutionUnit()],
					ConnectTarget:    execs[ws.Websocket.ConnectTarget.GetExecutionUnit()],
					DisconnectTarget: execs[ws.Websocket.DisconnectTarget.GetExecutionUnit()],
					StackID:          stackID,
				})
				if err != nil {
					return err
				}
			}
		}

		// Deploy all schedules
		schedules := map[string]*schedule.AwsEventbridgeSchedule{}
		for _, res := range request.Spec.Resources {
			switch t := res.Config.(type) {
			case *deploy.Resource_Schedule:
				// get the target of the schedule

				execUnitName := t.Schedule.Target.GetExecutionUnit()
				execUnit, ok := execs[execUnitName]
				if !ok {
					return fmt.Errorf("no execution unit with name %s", execUnitName)
				}

				// Create schedule targeting a given lambda
				schedules[res.Name], err = schedule.NewAwsEventbridgeSchedule(ctx, res.Name, &schedule.AwsEventbridgeScheduleArgs{
					Exec: execUnit,
					Cron: t.Schedule.Cron,
					Tz:   config.ScheduleTimezone,
				})
				if err != nil {
					return err
				}
			}
		}

		// Deploy all topics
		topics := map[string]*topic.SNSTopic{}
		for _, res := range request.Spec.Resources {
			switch t := res.Config.(type) {
			case *deploy.Resource_Topic:
				// Create topics
				topics[res.Name], err = topic.NewSNSTopic(ctx, res.Name, &topic.SNSTopicArgs{
					StackID: stackID,
					Topic:   t.Topic,
				})
				if err != nil {
					return err
				}

				// Create subscriptions for the topic
				for _, sub := range t.Topic.Subscriptions {
					subName := fmt.Sprintf("%s-%s-sub", sub.GetExecutionUnit(), res.Name)
					// Get the deployed execution unit
					unit, ok := execs[sub.GetExecutionUnit()]
					if !ok {
						return fmt.Errorf("invalid execution unit %s given for topic subscription", sub.GetExecutionUnit())
					}

					_, err = topic.NewSNSTopicSubscription(ctx, subName, &topic.SNSTopicSubscriptionArgs{
						Lambda: unit,
						Topic:  topics[res.Name],
					})
					if err != nil {
						return err
					}
				}
			}
		}

		// Create policies
		for _, res := range request.Spec.Resources {
			switch t := res.Config.(type) {
			case *deploy.Resource_Policy:
				_, err = policy.NewIAMPolicy(ctx, res.Name, &policy.PolicyArgs{
					Policy: t.Policy,
					Resources: &policy.StackResources{
						Buckets:     buckets,
						Topics:      topics,
						Queues:      queues,
						Collections: collections,
						Secrets:     secrets,
						Websockets:  websockets,
					},
					Principals: principals,
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	_ = pulumiStack.SetConfig(context.TODO(), "aws:region", auto.ConfigValue{Value: details.Region})
	_ = pulumiStack.SetConfig(context.TODO(), "supabase:version", auto.ConfigValue{Value: "0.0.1"})

	messageWriter := &pulumiutils.UpStreamMessageWriter{
		Stream: stream,
	}

	if config.Refresh {
		_ = stream.Send(&deploypb.DeployUpEvent{
			Content: &deploypb.DeployUpEvent_Message{
				Message: &deploypb.DeployEventMessage{
					Message: "refreshing pulumi stack",
				},
			},
		})
		// refresh the stack first
		_, err := pulumiStack.Refresh(context.TODO(), optrefresh.ProgressStreams(messageWriter))
		if err != nil {
			return err
		}
	}

	// Run the program
	res, err := pulumiStack.Up(context.TODO(), optup.ProgressStreams(messageWriter))
	if err != nil {
		return err
	}

	// Send terminal message
	err = stream.Send(pulumiutils.PulumiOutputsToResult(res.Outputs))

	return err
}
