package dashboard

import "go.uber.org/zap"

func initializeLogger() (*zap.Logger, error) {

	config := zap.NewDevelopmentConfig()

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
