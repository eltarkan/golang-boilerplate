package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/config"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "welcome",
	Short: "Boilerplate",
	Long:  `Welcome To Boilerplate`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome To Boilerplate")
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `Test the Boilerplate`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tests")
	},
}

func init() {
	config.LoadConfig()
	rootCmd.AddCommand(testCmd)

}
func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
