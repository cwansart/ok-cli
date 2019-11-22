package cmd

import (
	okconfig "github.com/cwansart/ok-cli/internal/config"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCleanURL(t *testing.T) {
	var testData = []struct {
		config        okconfig.Config
		path          string
		correctResult string
	}{
		{
			createTestConfig("", "", "http://localhost", ""),
			"/api/v1/admin/users",
			"http://localhost/api/v1/admin/users",
		},
		// add more tests here
	}

	for _, tt := range testData {
		got := cleanUrl(tt.config.GiteaURL, tt.path)
		assert.Equal(t, got, tt.correctResult)
	}
}

func createTestConfig(username, password, giteaURL, jenkinsURL string) okconfig.Config {
	return okconfig.Config{
		Username:   username,
		Password:   password,
		JenkinsURL: jenkinsURL,
		GiteaURL:   giteaURL,
	}
}
