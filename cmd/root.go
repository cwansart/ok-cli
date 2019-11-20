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
	userNameKey   = "OK_USER_NAME"
	passwordKey   = "OK_PASSWORD"
	jenkinsUrlKey = "OK_JENKINS_URL"
	giteaUrlKey   = "OK_GITEA_URL"
)

var rootCmd = &cobra.Command{
	Use:   "ok <command> <action> [parameters]",
	Short: "ok is a management tool for projects",
	Long:  "ok manages users and projects in the infrastructure set-up via env variables.",
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
		_, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s is not set but is required to work.", key)
			os.Exit(1)
		}
	}

	check(jenkinsUrlKey)
	check(giteaUrlKey)
}

// Creates a clean url without any trailing slashes or special characters.
func cleanUrl(remoteKey string, remotePath string) string {
	// TODO: differentiate between Gitea and Jenkins backend
	rawurl := os.Getenv(remoteKey) + remotePath
	url, err := neturl.Parse(rawurl)

	// TODO: proper error handling
	if err != nil {
		fmt.Printf("An error occured: %s", err)
		os.Exit(1)
	}

	// TODO: add https support and disable http
	if url.Scheme != "http" {
		fmt.Printf("Invalid server type %s, only http is supported.", url.Scheme)
	}

	return fmt.Sprintf("%s://%s:%s%s", url.Scheme, url.Hostname(), url.Port(), path.Clean(url.EscapedPath()))
	// alternatively use StringBuilder
}
