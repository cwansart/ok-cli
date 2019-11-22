package cmd

import (
	okconfig "github.com/cwansart/ok-cli/internal/config"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCleanURL(t *testing.T) {
	var testData = []struct {
		config okconfig.Config
		path   string
		want   string
	}{
		{
			createTestConfig("", "", "http://localhost", ""),
			"/api/v1/admin/users",
			"http://localhost/api/v1/admin/users",
		},
		// add more tests here
	}

	for _, td := range testData {
		got := cleanUrl(td.config.GiteaURL, td.path)
		assert.Equal(t, got, td.want)
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
