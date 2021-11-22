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
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/github"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// listCmd represents the list command
var listGithubOrgsCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		if viper.GetString("token") == "" {
			log.Fatal("configure your PAT token before calling this command")
		}
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: viper.GetString("token")},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		// list all repositories for the authenticated user
		orgs, _, err := client.Organizations.List(ctx, "", nil)
		// log.Print(orgs)

		var organizations []string
		var organization *github.Organization
		for _, organization = range orgs {
			strPointerValue := *organization.Login
			organizations = append(organizations, strPointerValue)
		}
		log.Print(err)

		searcher := func(input string, index int) bool {
			org := organizations[index]
			name := strings.Replace(strings.ToLower(org), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		}

		prompt := promptui.Select{
			Label:    "Pick org",
			Items:    organizations,
			Size:     4,
			Searcher: searcher,
		}

		i, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose number %d: %s\n", i+1, organizations[i])
	},
}

func init() {
	organizationsCmd.AddCommand(listGithubOrgsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
