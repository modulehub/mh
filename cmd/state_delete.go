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
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gojek/heimdall/httpclient"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		validate := func(input string) error {
			if input != "yes" {
				return errors.New("Need explicit \"yes\"")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Do you realy want to delete state_id for org?",
			Validate: validate,
		}

		_, err := prompt.Run()

		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}

		bkey := []byte(viper.GetString("email") + ":" + viper.GetString("apikey"))
		key := base64.StdEncoding.EncodeToString(bkey)
		// Create a new HTTP client with a default timeout
		timeout := 1000 * time.Millisecond
		client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
		headers := http.Header{}
		headers.Set("Authorization", "Basic "+key)

		res, err := client.Delete("http://localhost:81/api/organizations/modulehub/states/288319c1-3ce7-4bf3-910b-50a75faa7f64", headers)
		if err != nil {
			panic(err)
		}
		// Heimdall returns the standard *http.Response object
		body, err := ioutil.ReadAll(res.Body)
		log.Println(string(body))
	},
}

func init() {
	stateCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
