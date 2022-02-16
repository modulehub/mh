package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stateInitCmd represents the showConfig command
var stateInitCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("states init called")
		var err error
		var otpt []byte
		registryURL := viper.GetString("registry_url")
		email := viper.GetString("email")
		key := viper.GetString("apikey")
		org := viper.GetString("organization")
		pwd, _ := os.Getwd()
		var sid string
		if len(args) == 1 && args[0] != "" {
			sid = args[0]
		} else if _, errs := os.Stat(pwd + "/.mhrc"); errs == nil {
			logrus.Info(".mhrc found")
			file, erro := os.Open(".mhrc")
			if erro != nil {
				logrus.Fatal(err)
			}
			b, errb := ioutil.ReadAll(file)
			if errb != nil {
				logrus.Fatal(err)
			}
			sid = string(b)
		}
		logrus.Info(sid)
		url := fmt.Sprintf("%s/%s/remote_states/%s", registryURL, org, sid)
		tfcmd := exec.Command("terraform", "init", `-backend-config=address=`+url+``,
			`-backend-config=lock_address=`+url+`/lock`,
			`-backend-config=unlock_address=`+url+`/unlock`,
			`-backend-config=username=`+email+``,
			`-backend-config=password=`+key+``,
			`-backend-config=lock_method=POST`,
			`-backend-config=unlock_method=POST`)
		logrus.Info(tfcmd.String())
		switch runtime.GOOS {
		case "darwin":
			otpt, err = tfcmd.CombinedOutput()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err != nil {
			logrus.Info(string(otpt))
			logrus.Fatal(err)
		} else {
			logrus.Info(string(otpt))
		}
	},
}

func init() {
	stateCmd.AddCommand(stateInitCmd)
}
