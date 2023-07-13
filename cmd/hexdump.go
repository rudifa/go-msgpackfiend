/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hexdumpCmd represents the hexdump command
var hexdumpCmd = &cobra.Command{
	Use:   "hexdump",
	Short: "Hexdump a file",
	Long: `Hexdump a file...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hexdump called")
	},
}

func init() {
	rootCmd.AddCommand(hexdumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hexdumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hexdumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
