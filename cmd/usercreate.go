package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var userCreateCmd = &cobra.Command{
	Use:   "create --email <email> --loginName <loginName> --password <password>",
	Short: "", // TODO: add Short and Long
	Long:  "",
	Run:   userCreate,
}

var (
	email     string
	loginName string
	password  string
)

type createUserOption struct {
	Email              string `json:"email"`
	LoginName          string `json:"login_name"`
	MustChangePassword bool   `json:"must_change_password"`
	Password           string `json:"password"`
}

func init() {
	userCreateCmd.Flags().StringVarP(&email, "email", "e", "", "User's email address")
	userCreateCmd.Flags().StringVarP(&loginName, "loginName", "l", "", "User's login name")
	userCreateCmd.Flags().StringVarP(&password, "password", "p", "", "User's password")
	_ = userCreateCmd.MarkFlagRequired("email")
	_ = userCreateCmd.MarkFlagRequired("loginName")
	_ = userCreateCmd.MarkFlagRequired("password")
}

func userCreate(cmd *cobra.Command, _ []string) {
	u := createUserOption{
		Email:              email,
		LoginName:          loginName,
		MustChangePassword: false,
		Password:           password,
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("An error occurred during marshalling: ", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", userCreateUrl(), bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("An error occurred during  request creation: ", err)
		return
	}

	// TODO: perhaps using a token would be better? See /admin​/users​/{username}​/keys route
	req.SetBasicAuth(os.Getenv(userNameKey), os.Getenv(passwordKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("An error occurred during request: ", err)
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body) // TODO: error handling

	switch resp.StatusCode {
	case 201:
		fmt.Println("User has been created. ", string(respBody))
	case 403:
		fmt.Println("Request failed (403). You are not authorized for this action.", string(respBody))
	case 422:
		// api validation error; this should not happen except there are changes in the api
		fmt.Println("Request failed (422). The Gitea api may have changed.", string(respBody))
	}
}

func userCreateUrl() string {
	return cleanUrl(giteaUrlKey, "/api/v1/admin/users")
}
