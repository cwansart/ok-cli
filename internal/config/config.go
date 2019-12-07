package config

import (
	"fmt"
	"os"
)

type Config struct {
	Jenkins JenkinsConfig
	Gitea   GiteaConfig
}

type JenkinsConfig struct {
	Username string
	Password string
	Url      string
	CliPath  string
}

type GiteaConfig struct {
	Username string
	Password string
	Url      string
}

func NewConfig() Config {
	return Config{
		Jenkins: JenkinsConfig{
			Username: getEnvVar("OK_JENKINS_USERNAME"),
			Password: getEnvVar("OK_JENKINS_PASSWORD"),
			Url:      getEnvVar("OK_JENKINS_URL"),
			CliPath:  getEnvVar("OK_JENKINS_CLI_PATH"),
		},
		Gitea: GiteaConfig{
			Username: getEnvVar("OK_GITEA_USERNAME"),
			Password: getEnvVar("OK_GITEA_PASSWORD"),
			Url:      getEnvVar("OK_GITEA_URL"),
		},
	}
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		_ = fmt.Errorf("Could not find env %s\n", key)
	}
	return val
}
