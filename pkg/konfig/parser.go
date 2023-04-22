package konfig

import (
	"encoding/json"
	"konfig/model"
)

func Parse(config string) (model.Config, error) {
	var jsonMap model.Config
	err := json.Unmarshal([]byte(config), &jsonMap)
	if err != nil {
		return model.Config{}, nil
	}
	return jsonMap, nil
}
