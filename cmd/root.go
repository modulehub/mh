package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mh",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// if cmd.CalledAs() != "init" {
		// 	if err := viper.ReadInConfig(); err != nil {
		// 		fmt.Println("Run init first")
		// 		os.Exit(1)
		// 	}
		// }
	},
	Version: version, //this field has to be set in order to have --version working
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mh.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("organization", "o", "", "Help message for organization")

	err := viper.BindPFlag("organization", rootCmd.PersistentFlags().Lookup("organization"))
	if err != nil {
		log.Fatal(err)
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// stub init
	configHome, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetDefault("app_url", "https://app.modulehub.io")
	viper.SetDefault("api_url", "https://api.v2.modulehub.io/api")
	viper.SetDefault("registry_url", "https://registry.v2.modulehub.io")

	configName := ".mh"
	configType := "yaml"
	configPath := filepath.Join(configHome, ".modulehub")
	configFile := filepath.Join(configPath, configName+"."+configType)

	err = os.MkdirAll(configPath, 0755)
	if err != nil { // handle failed create
		log.Info(err)
	} else {
		log.Info("config file exists")
	}
	_, err = os.OpenFile(configFile, os.O_CREATE, 0644)
	if err != nil { // handle failed create
		log.Info(err)
	} else {
		log.Info("config file exists")
	}
	// ----
	viper.SetEnvPrefix("mh") // will be uppercased automatically

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	err = viper.ReadInConfig()
	if err != nil { // handle failed create
		log.Info(err)
	}

	viper.SetConfigName(".mh.local")
	viper.AddConfigPath(".")
	err = viper.MergeInConfig()
	if err != nil { // handle failed create
		log.Info(err)
	}

	// Find home directory.
	// Search config in home directory with name ".mh" (without extension).
	viper.AutomaticEnv() // read in environment variables that match

	// log.Info(viper.AllSettings())
}
