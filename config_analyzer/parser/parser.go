package parser

import (
	"encoding/json"
	"errors"

	"go.yaml.in/yaml/v3"
)

func Parse(data []byte) (map[string]interface{}, error) {

	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err == nil {
		return result, nil
	}

	if err := yaml.Unmarshal(data, &result); err == nil {
		return result, nil
	}

	return nil, errors.New("unsupported format:  not valid JSON or YAML")
}
