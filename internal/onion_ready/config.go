package onion_ready

import (
	types "hve/onion-simulate/internal/onion/types"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func ReadConfig(path string) (*types.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config types.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
