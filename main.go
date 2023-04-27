package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "print-json",
	Short: "Print JSON to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		jsonString, err := cmd.Flags().GetString("json")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println(jsonString)
	},
}

func main() {
	rootCmd.Flags().String("json", "", "JSON string to print")
	rootCmd.MarkFlagRequired("json")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
	}
}
