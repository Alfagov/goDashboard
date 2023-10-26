package config

import (
	"encoding/json"
	"github.com/Alfagov/goDashboard/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

var C Config

func InitConfig() error {

	var t_config Config

	fileName := os.Getenv("GD_CONFIG_NAME")
	extension := os.Getenv("GD_CONFIG_TYPE")
	basePath := os.Getenv("GD_CONFIG_PATH")

	configPath := basePath + "/" + fileName + "." + extension

	file, err := os.ReadFile(configPath)
	if err != nil {
		logger.L.Error("Error opening config file", zap.Error(err))
		panic(err)
	}

	switch extension {
	case "yaml":
		err = yaml.Unmarshal(file, &t_config)
		if err != nil {
			logger.L.Error("Error unmarshaling yaml", zap.Error(err))
		}
	case "json":
		err = json.Unmarshal(file, &t_config)
		if err != nil {
			logger.L.Error("Error unmarshaling json", zap.Error(err))
		}
	}

	C = t_config

	return nil
}
