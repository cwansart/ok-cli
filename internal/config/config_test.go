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
	setEnv(t, "OK_USERNAME", "admin")
	setEnv(t, "OK_PASSWORD", "admin")
	setEnv(t, "OK_JENKINS_URL", "localhost:8080")
	setEnv(t, "OK_GITEA_URL", "localhost:8081")

	c := NewConfig()

	assert.Equal(t, "admin", c.Username)
	assert.Equal(t, "admin", c.Password)
	assert.Equal(t, "localhost:8080", c.JenkinsURL)
	assert.Equal(t, "localhost:8081", c.GiteaURL)
}

func setEnv(t *testing.T, key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatalf("could not set env var: %s", err)
	}
}
