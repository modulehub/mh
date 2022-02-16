package cmd

import (
	"os"

	"github.com/modulehub/mh/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//RegisterResponse contains the result of user registration
type RegisterResponse struct {
	Data struct {
		Email string `json:"email"`
	}
	Key   string `json:"key"`
	Org   string `jsong:"org"`
	Error string
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// validate := func(input string) error {
		// 	if ok := util.ValidateEmail(input); !ok {
		// 		return errors.New("invalid email")
		// 	}
		// 	return nil
		// }

		// prompt := promptui.Prompt{
		// 	Label:    "Github account email",
		// 	Validate: validate,
		// }

		// result, err := prompt.Run()

		// if err != nil {
		// 	logrus.Printf("Prompt failed %v\n", err)
		// 	return
		// }

		// logrus.Printf("Your mail: %q\n", result)

		// // Create a new HTTP client with a default timeout
		// //
		// client := util.GetClient()
		// postBody, _ := json.Marshal(map[string]string{
		// 	"email": result,
		// })
		// responseBody := bytes.NewBuffer(postBody) // Use the clients GET method to create and execute the request
		// res, err := client.Post("/users?type=cli", responseBody)
		// if err != nil {
		// 	panic(err)
		// }
		// // Heimdall returns the standard *http.Response object

		// var key RegisterResponse

		// if err := json.NewDecoder(res.Body).Decode(&key); err != nil {
		// 	logrus.Info(err)
		// }

		// if res.StatusCode != 200 {
		// 	logrus.Info(res.StatusCode)
		// 	logrus.Error(key.Error) //TODO resolve the error in a nicer way..
		// 	os.Exit(1)
		// }

		// viper.Set("email", key.Data.Email)
		// viper.Set("apikey", key.Key)
		// viper.Set("organization", key.Org)

		// if err := viper.WriteConfig(); err != nil {
		// 	logrus.Info(err)
		// }

		// fmt.Println("Check your email for activation link")

		err := util.OpenURL(os.Getenv(("MH_APP_BASE_URL")))
		if err != nil {
			logrus.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
