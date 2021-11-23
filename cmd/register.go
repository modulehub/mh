/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/modulehub/mh/utility"
	log "github.com/sirupsen/logrus"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RegisterResponse struct {
	Data struct {
		Email string `json:"email"`
	}
	Key string `json:"key"`
	Org string `jsong:"org"`
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			if ok := utility.ValidateEmail(input); ok == false {
				return errors.New("Invalid email")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Github account email",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}

		log.Printf("Your mail: %q\n", result)

		// Create a new HTTP client with a default timeout
		//
		client := utility.GetClient()
		postBody, _ := json.Marshal(map[string]string{
			"email": result,
		})
		responseBody := bytes.NewBuffer(postBody) // Use the clients GET method to create and execute the request
		res, err := client.Post("/users?type=cli", responseBody)
		if err != nil {
			panic(err)
		}

		// Heimdall returns the standard *http.Response object
		var key RegisterResponse

		if err := json.NewDecoder(res.Body).Decode(&key); err != nil {
			log.Println(err)
		}
		log.Println(key)

		viper.Set("email", key.Data.Email)
		viper.Set("apikey", key.Key)
		viper.Set("organization", key.Org)

		if err := viper.WriteConfig(); err != nil {
			log.Println(err)
		}

		// log.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
