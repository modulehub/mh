package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/modulehub/mh/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//StateListResponse holds list of states
type StateListResponse struct {
	Data []State
}

//State type
type State struct {
	ID string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("create called")
		// Create a new HTTP client with a default timeout
		//
		client := util.GetClient()
		res, err := client.Get("/organizations/"+viper.GetString("organization")+"/states", nil)
		if err != nil {
			panic(err)
		}

		var states StateListResponse
		if e := json.NewDecoder(res.Body).Decode(&states); e != nil {
			logrus.Fatal(e)
		}
		for _, state := range states.Data {
			fmt.Println(state.ID)
		}
	},
}

func init() {
	stateCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
