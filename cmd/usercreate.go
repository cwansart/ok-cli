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
	Use:   "create --email <email> --username <username> --password <password>",
	Short: "", // TODO: add Short and Long
	Long:  "",
	Run:   userCreate,
}

var (
	email    string
	username string
	password string
)

type createUserOption struct {
	Email              string `json:"email"`
	Username           string `json:"username"`
	MustChangePassword bool   `json:"must_change_password"`
	Password           string `json:"password"`
}

func init() {
	userCreateCmd.Flags().StringVarP(&email, "email", "e", "", "User's email address")
	userCreateCmd.Flags().StringVarP(&username, "username", "u", "", "User's login name")
	userCreateCmd.Flags().StringVarP(&password, "password", "p", "", "User's password")
	_ = userCreateCmd.MarkFlagRequired("email")
	_ = userCreateCmd.MarkFlagRequired("username")
	_ = userCreateCmd.MarkFlagRequired("password")
}

func userCreate(_ *cobra.Command, _ []string) {
	s := make(chan bool, 2)
	go createOnGitea(s)
	go createOnJenkins(s)

	g, j := <-s
	// @Max: does this work?
	if !g || !j {
		// TODO: undo changes if one fails
	}
}

func createOnGitea(s chan bool) {
	u := createUserOption{
		Email:              email,
		Username:           username,
		MustChangePassword: false,
		Password:           password,
	}

	// TODO: change to decoder?
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("An error occurred during marshalling:  %s\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", giteaUserCreateUrl(), bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("An error occurred during request creation: %s\n", err)
		return
	}

	// TODO: perhaps using a token would be better? See /admin​/users​/{username}​/keys route
	req.SetBasicAuth(os.Getenv(usernameKey), os.Getenv(passwordKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("An error occurred during request: %s\n", err)
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body) // TODO: error handling

	switch resp.StatusCode {
	case 201:
		fmt.Printf("User has been created. %s\n", string(respBody))
		s <- true
	case 403:
		fmt.Printf("Request failed (403). You are not authorized for this action. %s\n", string(respBody))
		s <- false
	case 422:
		// api validation error; this should not happen except there are changes in the api
		fmt.Printf("Request failed (422). The Gitea api may have changed. %s\n", string(respBody))
	default:
		fmt.Printf("Request failed (%d). %s\n", resp.StatusCode, string(respBody))
	}
}

func createOnJenkins(s chan bool) {
	// TODO: to be implemented
	s <- true
}

func giteaUserCreateUrl() string {
	return cleanUrl(giteaUrlKey, "/api/v1/admin/users")
}
