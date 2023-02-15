package config

import "os"

type Mysql struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
}

func GetEnvWithDefautl(variable string, def string) string {
	env := os.Getenv(variable)
	if len(env) != 0 {
		return env
	}
	return def
}
