package cmd

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
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

		pwd, _ := os.Getwd()
		var sid string
		var err error
		if len(args) == 1 && args[0] != "" {
			sid = args[0]
		} else if _, errs := os.Stat(pwd + "/.mhrc"); errs == nil {
			log.Info(".mhrc found")
			file, erro := os.Open(".mhrc")
			if erro != nil {
				log.Fatal(err)
			}
			b, errb := ioutil.ReadAll(file)
			if errb != nil {
				log.Fatal(err)
			}
			sid = string(b)
		}
		log.Info(sid)

		t, err := template.New("state").Parse(stateTpl)
		if err != nil {
			panic(err)
		}
		err = t.Execute(os.Stdout, map[string]string{"ID": sid, "key": viper.GetString("APIKey"), "username": viper.GetString("email")})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	stateCmd.AddCommand(showConfigCmd)
}
