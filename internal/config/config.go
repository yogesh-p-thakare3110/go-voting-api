package config

import "os"

type Config struct {
	Environment      string
	ConnectionString string
}

func getConfigValue(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}

	return defaultValue
}

func NewConfig() *Config {
	return &Config{
		Environment:      getConfigValue("EVN", "local"),
		ConnectionString: getConfigValue("CONN_STRING", "root:India2000@tcp(localhost:3306)/yogeshdb"),
	}
}
