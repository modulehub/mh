package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/modulehub/mh/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		log.Info("prerun")
		if !viper.GetBool("force") && viper.Get("apikey") != nil {
			cmd.SilenceUsage = true
			return errors.New("Configuration exists, use --force to override")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("cmd init")
		validate := func(input string) error {
			if len(input) < 5 {
				return errors.New("invalid value")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "MH account email",
			Validate: util.ValidateEmailFunc,
		}

		result, errm := prompt.Run()

		if errm != nil {
			log.Printf("Prompt failed %v\n", errm)
			return
		}

		log.Printf("Your mail: %q\n", result)
		viper.Set("email", result)

		promptKey := promptui.Prompt{
			Label:    "API key",
			Validate: validate,
		}

		key, errk := promptKey.Run()
		if errk != nil {
			log.Printf("Prompt failed %v\n", errk)
			return
		}
		viper.Set("apikey", key)

		promptOrg := promptui.Prompt{
			Label:    "Organization",
			Validate: validate,
		}

		org, erro := promptOrg.Run()
		if erro != nil {
			log.Printf("Prompt failed %v\n", erro)
			return
		}
		viper.Set("organization", org)

		if err := viper.WriteConfig(); err != nil {
			log.Info(err)
		}

		fmt.Println("You're good to go!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("force", "f", false, "Force")
}
