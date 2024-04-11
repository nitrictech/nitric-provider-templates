# Terraform Golang Provider

This is a nitric deployment provider designed to allow teams to re-use their existing HCL modules and build a nitric compatible API
around them in order to generate complete deployments.

## How it works

This project used [cdktf]() to generate programmatic bindings for existing HCL modules for golang.

This allows nitric applications to be defined deployed using leveraging HCL stacks.