package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  "",
	Run:   userCreate,
}

func userCreate(_ *cobra.Command, _ []string) {
	fmt.Println("tbd")
}
