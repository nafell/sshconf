/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nafell/sshconf/core"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all hosts written in ~/.ssh/config",
	Long: `Lists all hosts written in ~/.ssh/config

The output will be shown as following:
name user@example.com:22
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		contents, err := core.ReadConfigFile()
		if err != nil {
			return err
		}
		configFileInfo, err := core.SplitEntryBlocks(contents)
		if err != nil {
			return err
		}
		entries := core.MapStruct(configFileInfo)

		for _, entry := range entries {
			//fmt.Printf("%v", entry)
			entry.PrintPretty()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
