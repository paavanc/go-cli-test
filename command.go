package main

import (

	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-cli-test/manager"
	"go-cli-test/utils"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-test",
	Short: "A sample project to house golang scripts",
}

var testCmd = &cobra.Command{
	Use:   "hello",
	Short: "hello world test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.HomeDir())
		fmt.Println("Hello World!")
	},
}

var testHttpRequest = &cobra.Command{
	Use:   "rest",
	Short: "makes a sample api call",
	Run: func(cmd *cobra.Command, args []string) {
		manager.RestTest();
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
func init() {
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(testHttpRequest)
}
