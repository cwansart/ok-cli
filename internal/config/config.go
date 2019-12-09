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
			// I am not sure if "OK_JENKINS_USERNAME" is sufficient. When we have normal user command which don't
			// require administative powers we may have overlapping names. For example, I have "jenkins" the admin
			// user on Gitea and a normal user "cwansart" on Gitea. How do I run `ok init <projectName>` as user
			// "cwansart" without the need to change those env vars every time?
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
