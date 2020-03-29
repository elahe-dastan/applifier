package cmd

import (
	"fmt"
	"os"

	"github.com/elahe-dastan/applifier/cmd/client"
	"github.com/elahe-dastan/applifier/cmd/server"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "applifier",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}

	server.Register(rootCmd)
	client.Register(rootCmd)

	exitFailure := 1

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitFailure)
	}
}

//func init() {
//	cobra.OnInitialize(initConfig)
//
//	// Here you will define your flags and configuration settings.
//	// Cobra supports persistent flags, which, if defined here,
//	// will be global for your application.
//
//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.applifier.yaml)")
//
//	// Cobra also supports local flags, which will only run
//	// when this action is called directly.
//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}

// initConfig reads in config file and ENV variables if set.