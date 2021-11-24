package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/modulehub/mh/utility"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"html/template"
)

//StateResponse holds the state response
type StateResponse struct {
	Data struct {
		ID string `json:"id"`
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		// Create a new HTTP client with a default timeout
		//
		client := utility.GetClient()
		res, err := client.Post("/organizations/"+viper.GetString("organization")+"/states", nil)
		if err != nil {
			panic(err)
		}

		// Heimdall returns the standard *http.Response object
		var state StateResponse
		if err := json.NewDecoder(res.Body).Decode(&state); err != nil {
			log.Fatal(err)
		}

		log.Info(fmt.Sprintf("state id: %s created.", state.Data.ID))

		tmpl := `
Usage:

terraform {
	backend "http" {
		username = "{{ .username }}"
		password = "{{ .key }}"
		address = "https://registry.v2.modulehub.io/remote_states/{{ .ID }}"
		lock_address = "https://registry.v2.modulehub.io/remote_states/{{ .ID }}/lock"
		unlock_address = "https://registry.v2.modulehub.io/remote_states/{{ .ID }}/unlock"
		lock_method = "POST"
		unlock_method = "POST"
	}
}
		`

		t, err := template.New("todos").Parse(tmpl)
		if err != nil {
			panic(err)
		}
		err = t.Execute(os.Stdout, map[string]string{"ID": state.Data.ID, "key": viper.GetString("apikey"), "username": viper.GetString("email")})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	stateCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
