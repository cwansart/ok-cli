package cmd

import (
	"fmt"
	neturl "net/url"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// Is there a better way to do that in Go?
const (
	usernameKey   = "OK_USERNAME"
	passwordKey   = "OK_PASSWORD"
	jenkinsUrlKey = "OK_JENKINS_URL"
	giteaUrlKey   = "OK_GITEA_URL"
)

var rootCmd = &cobra.Command{
	Use:   "ok <command> <action> [parameters]",
	Short: "ok is a management tool for projects",
	Long: "\nok manages users and projects in the infrastructure set-up via env variables. " +
		"Required environment variables" +
		"\n\tOK_JENKINS_URL\t- URL for the Jenkins instance" +
		"\n\tOK_GITEA_URL\t- URL for the Gitea Instance" +
		"\n\tOK_USER_NAME\t- Gitea's admin username" +
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

	checkEnv()
}

// Checks if environment variables for the external server, user name and password are set.
func checkEnv() {
	check := func(key string) {
		if _, ok := os.LookupEnv(key); !ok {
			fmt.Printf("%s is not set but is required to work.\n", key)
			os.Exit(1)
		}
	}

	check(jenkinsUrlKey)
	check(giteaUrlKey)
}

// Creates a clean url without any trailing slashes or special characters.
func cleanUrl(remoteKey string, remotePath string) string {
	// TODO: differentiate between Gitea and Jenkins backend
	rawURL := os.Getenv(remoteKey) + remotePath
	url, err := neturl.Parse(rawURL)

	// TODO: proper error handling
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
		os.Exit(1)
	}

	// TODO: add https support and disable http
	if url.Scheme != "http" {
		fmt.Printf("Invalid server type %s\n, only http is supported.", url.Scheme)
	}

	// alternatively use StringBuilder
	return fmt.Sprintf("%s://%s:%s%s", url.Scheme, url.Hostname(), url.Port(), path.Clean(url.EscapedPath()))
}
