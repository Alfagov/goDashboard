package logger

import "go.uber.org/zap"

var L *zap.Logger

func InitLogger() error {
	config := zap.NewDevelopmentConfig()

	logger, err := config.Build()
	if err != nil {
		return err
	}

	L = logger

	return nil
}
