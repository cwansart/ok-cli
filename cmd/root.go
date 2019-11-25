package cmd

import (
	"fmt"
	neturl "net/url"
	"os"
	"path"
	"strings"

	// 'config' collides with name declared in this package
	okconfig "github.com/cwansart/ok-cli/internal/config"

	"github.com/spf13/cobra"
)

var config okconfig.Config

var rootCmd = &cobra.Command{
	Use:   "ok <command> <action> [parameters]",
	Short: "ok is a management tool for projects",
	Long: "\nok manages users and projects in the infrastructure set-up via env variables. " +
		"Required environment variables" +
		"\n\tOK_JENKINS_URL\t- URL for the Jenkins instance" +
		"\n\tOK_GITEA_URL\t- URL for the Gitea Instance" +
		"\n\tOK_USERNAME\t- Gitea's admin username" +
		"\n\tOK_PASSWORD\t- Gitea's admin password",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userCreateCmd)
	userCmd.AddCommand(userListCmd)

	config = okconfig.NewConfig()
}

// Creates a clean url without any trailing slashes or special characters.
func cleanUrl(remoteKey string, remotePath string) string {
	// TODO: differentiate between Gitea and Jenkins backend
	rawURL := fmt.Sprintf("%s%s", remoteKey, remotePath)

	url, err := neturl.Parse(rawURL)
	if err != nil {
		fmt.Errorf("Could not parse URL: %s\n", err)
	}

	// TODO: add https support and disable http
	if url.Scheme != "http" {
		fmt.Printf("Invalid server type %s\n, only http is supported.", url.Scheme)
	}

	// is correct, but doesnt look good
	var urlBuilder strings.Builder

	urlBuilder.WriteString(url.Scheme)
	urlBuilder.WriteString("://")
	urlBuilder.WriteString(url.Hostname())
	if len(url.Port()) != 0 {
		urlBuilder.WriteString(":")
		urlBuilder.WriteString(url.Port())
	}
	urlBuilder.WriteString(path.Clean(url.EscapedPath()))

	return urlBuilder.String()
}
