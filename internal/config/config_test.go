package config

import (
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func TestGetEnvVar(t *testing.T) {
	if err := os.Setenv("KEY", "value"); err != nil {
		t.Fatalf("Could not set env var: %s", err)
	}
	envVar := getEnvVar("KEY")
	assert.Equal(t, "value", envVar)

	envVar = getEnvVar("NOT_EXISTING_KEY")
	assert.Equal(t, "", envVar)
}

func TestNewConfig(t *testing.T) {
	setEnv(t, "OK_JENKINS_USERNAME", "admin")
	setEnv(t, "OK_JENKINS_PASSWORD", "admin")
	setEnv(t, "OK_JENKINS_URL", "http://localhost:8080")
	setEnv(t, "OK_JENKINS_CLI_PATH", "/home/user/bin/jenkins-cli.jar")
	setEnv(t, "OK_GITEA_USERNAME", "admin")
	setEnv(t, "OK_GITEA_PASSWORD", "admin")
	setEnv(t, "OK_GITEA_URL", "http://localhost:8081")

	c := NewConfig()

	assert.Equal(t, "admin", c.Jenkins.Username)
	assert.Equal(t, "admin", c.Jenkins.Password)
	assert.Equal(t, "http://localhost:8080", c.Jenkins.Url)
	assert.Equal(t, "/home/user/bin/jenkins-cli.jar", c.Jenkins.CliPath)
	assert.Equal(t, "admin", c.Gitea.Username)
	assert.Equal(t, "admin", c.Gitea.Password)
	assert.Equal(t, "http://localhost:8080", c.Gitea.Url)
}

func setEnv(t *testing.T, key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatalf("could not set env var: %s", err)
	}
}
