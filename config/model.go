package config

type Config struct {
	DashboardConfig dashboardConfig `yaml:"dashboard" ,json:"dashboard" ,mapstructure:"dashboard"`
	ServicesConfig  []serviceConfig `yaml:"services" ,json:"services" ,mapstructure:"services"`
	LoggerConfig    loggerConfig    `yaml:"logger" ,json:"logger" ,mapstructure:"logger"`
}

type dashboardConfig struct {
	Host     string `yaml:"host" ,json:"host" ,mapstructure:"host"`
	Port     string `yaml:"port" ,json:"port" ,mapstructure:"port"`
	SSL      bool   `yaml:"ssl" ,json:"ssl" ,mapstructure:"ssl"`
	CertPath string `yaml:"certPath" ,json:"certPath" ,mapstructure:"certPath"`
	KeyPath  string `yaml:"keyPath" ,json:"keyPath" ,mapstructure:"keyPath"`
}

type serviceConfig struct {
	Name        string                 `yaml:"name" ,json:"name" ,mapstructure:"name"`
	Description string                 `yaml:"description" ,json:"description" ,mapstructure:"description"`
	Options     map[string]interface{} `yaml:"options" ,json:"options" ,mapstructure:"options"`
}

type loggerConfig struct {
	Level string `yaml:"level" ,json:"level" ,mapstructure:"level"`
}

type DashboardSettings struct {
	IndexPage string
}
