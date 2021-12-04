package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showConfigCmd represents the showConfig command
var showConfigCmd = &cobra.Command{
	Use:   "showConfig",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showConfig called")

		t, err := template.New("state").Parse(stateTpl)
		if err != nil {
			panic(err)
		}
		err = t.Execute(os.Stdout, map[string]string{"ID": args[0], "key": viper.GetString("APIKey"), "username": viper.GetString("email")})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	stateCmd.AddCommand(showConfigCmd)
}
