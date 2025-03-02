package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say Hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello %s\n", viper.GetString("name"))
	},
}
