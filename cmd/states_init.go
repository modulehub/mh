package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("states init called")
		var err error
		var otpt []byte
		registryURL := viper.GetString("registry_url")
		email := viper.GetString("email")
		key := viper.GetString("apikey")
		trCmd := "terraform"
		params := []string{"init", "-backend-config=\"address=" + registryURL + "/modulehub/remote_states/${_REMOTE_STATE_ID}\"",
			"-backend-config=\"lock_address=" + registryURL + "/modulehub/remote_states/${_REMOTE_STATE_ID}/lock\"",
			"-backend-config=\"unlock_address=" + registryURL + "/modulehub/remote_states/${_REMOTE_STATE_ID}/unlock\"",
			"-backend-config=\"username=" + email + "\"",
			"-backend-config=\"password=" + key + "\"",
			"-backend-config=\"lock_method=POST\"",
			"-backend-config=\"unlock_method=POST\"",
		}
		prms := strings.Join(params, " ")
		log.Info(trCmd + " " + prms)
		switch runtime.GOOS {
		case "darwin":
			otpt, err = exec.Command(trCmd, prms).Output()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err != nil {
			log.Info(string(otpt))
			log.Fatal(err)
		} else {
			log.Info(string(otpt))
		}
	},
}

func init() {
	stateCmd.AddCommand(stateInitCmd)
}
