package utils

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	Config "gitee.com/under-my-umbrella/cloud/config"
)

// ReadConfig reads the configuration from a YAML file and returns a Config.Config
func ReadConfig() Config.Config {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Config.Config{} // Return an empty Config.Config object on error
	}

	var config Config.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return Config.Config{} // Return an empty Config.Config object on error
	}

	return config // Return the config object on success
}




