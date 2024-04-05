package deploy

import "github.com/mitchellh/mapstructure"

type CustomConfig struct{}

// Return Config from stack attributes
func ConfigFromAttributes(attributes map[string]interface{}) (*CustomConfig, error) {
	config := &CustomConfig{}
	err := mapstructure.Decode(attributes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
