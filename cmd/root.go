package cmd

import (
	"github.com/seed95/clean-web-service/application"
	"github.com/spf13/cobra"
	"log"
)

var (
	configFile = "./build/config/config.yaml"
	restType   string

	rootCmd = &cobra.Command{
		Use:   "app [-r echo | -s gin]",
		Short: "User is a very simple web service",
		Long: `A very simple web service with clean architecture and used popular library in Go such as gin, echo, gorm.

Complete documentation is available at https://github.com/seed95/clean-web-service`,
		Run: func(cmd *cobra.Command, args []string) {
			opt := &application.Options{
				ConfigFile: configFile,
				RestType:   restType,
			}

			if err := application.Run(opt); err != nil {
				log.Fatalln(err)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", configFile, "config file")
	rootCmd.Flags().StringVarP(&restType, "rest-mode", "r", "echo", "select rest framework between \"gin\" or \"echo\"")
}
