package cmd

import (
	"testing"

	"github.com/magiconair/properties/assert"

	okconfig "github.com/cwansart/ok-cli/internal/config"
)

func TestCleanURL(t *testing.T) {
	var testData = []struct {
		config okconfig.GiteaConfig
		path   string
		want   string
	}{
		{
			createTestConfig("", "", "http://localhost"),
			"/api/v1/admin/users",
			"http://localhost/api/v1/admin/users",
		},
		// add more tests here
	}

	for _, td := range testData {
		got := cleanUrl(td.config.Url, td.path)
		assert.Equal(t, got, td.want)
	}
}

func createTestConfig(username string, password string, url string) okconfig.GiteaConfig {
	return okconfig.GiteaConfig{
		Username: username,
		Password: password,
		Url:      url,
	}
}
