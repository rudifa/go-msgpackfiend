/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fromjsonCmd represents the fromjson command
var fromjsonCmd = &cobra.Command{
	Use:   "fromjson",
	Short: "Convert a json file to msgpack",
	Long: `Convert a json file to msgpack...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fromjson called")
	},
}

func init() {
	rootCmd.AddCommand(fromjsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fromjsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fromjsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
