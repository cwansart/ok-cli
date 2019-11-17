package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <action> [parameters]",
	Short: "Short",
	Long:  "Long",
}

func init() {
	_, ok := os.LookupEnv(userNameKey)
	if !ok {
		fmt.Println(userNameKey + " is not set but is required to work.")
		os.Exit(1)
	}

	_, ok = os.LookupEnv(passwordKey)
	if !ok {
		fmt.Println(passwordKey + " is not set but is required to work.")
		os.Exit(1)
	}
}
