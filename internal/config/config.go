package config

import (
	"fmt"
	"os"
)

// Config contains all the env var which the user set
type Config struct {
	Username   string
	Password   string
	JenkinsURL string
	GiteaURL   string
}

func NewConfig() Config {
	return Config{
		Username:   getEnvVar("OK_USERNAME"),
		Password:   getEnvVar("OK_PASSWORD"),
		JenkinsURL: getEnvVar("OK_JENKINS_URL"),
		GiteaURL:   getEnvVar("OK_GITEA_URL"),
	}
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		fmt.Errorf("Could not find env %s\n", key)
	}
	return val
}
