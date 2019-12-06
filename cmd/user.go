package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <action> [parameters]",
	Short: "Manage users",
	Long:  "Long", // TODO: add description
}

func init() {
	lookupKey(giteaUsernameKey)
	lookupKey(giteaPasswordKey)
	lookupKey(giteaUrlKey)
	lookupKey(jenkinsUrlKey)
	lookupKey(jenkinsUsernameKey)
	lookupKey(jenkinsPasswordKey)
	lookupKey(jenkinsCliPathKey)
}

func lookupKey(k string) {
	if _, ok := os.LookupEnv(k); !ok {
		fmt.Printf("%s is not set but is required to work.\n", k)
		os.Exit(1)
	}
}
