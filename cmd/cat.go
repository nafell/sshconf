/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/nafell/sshconf/core"
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat [Host]",
	Short: "Prints the contents of ~/.ssh/config  Optional: [Host]",
	Long: `Prints the contents of ~/.ssh/config to the standard output.

"sshconf cat"        prints the entire config.
"sshconf cat [Host]" prints the settings under entries exactly named "[Host]".`,
	RunE: func(cmd *cobra.Command, args []string) error {
		content, err := core.ReadConfigFile()
		if err != nil {
			return err
		}

		if len(args) < 1 {
			fmt.Println(content)
			return nil
		}

		configFileInfo, err := core.SplitEntryBlocks(content)
		if err != nil {
			return err
		}

		for _, block := range configFileInfo.Blocks {
			if strings.TrimSpace(strings.Replace(block[0], "Host", "", 1)) != args[0] {
				continue
			}

			fmt.Println(strings.Join(block, "\n"))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
