package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("open called")
		var err error
		url := viper.GetString("app_url")
		log.Info(url)
		if url != "" {
			switch runtime.GOOS {
			case "linux":
				err = exec.Command("xdg-open", url).Start()
			case "windows":
				err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			case "darwin":
				err = exec.Command("open", url).Start()
			default:
				err = fmt.Errorf("unsupported platform")
			}
			if err != nil {
				log.Fatal(err)
			} else {
				log.Println("opened")
			}
		} else {
			log.Println("url is empty")
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
