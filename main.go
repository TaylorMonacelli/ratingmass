package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "json-print",
	Short: "Print a JSON object to stdout",
	Run:   printJson,
}

var jsonObject string

func init() {
	rootCmd.PersistentFlags().StringVar(&jsonObject, "json", "", "The JSON object to print")
	rootCmd.MarkPersistentFlagRequired("json")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printJson(cmd *cobra.Command, args []string) {
	var obj interface{}
	err := json.Unmarshal([]byte(jsonObject), &obj)
	if err != nil {
		fmt.Println("Failed to parse JSON object", err)
		return
	}

	out, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Failed to serialize JSON object", err)
		return
	}

	fmt.Println(string(out))
}
