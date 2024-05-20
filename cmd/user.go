/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nafell/sshconf/core"
	"github.com/spf13/cobra"
)

// hostnameCmd represents the hostname command
var userCmd = &cobra.Command{
	Use:   "user <ENTRY_NAME> <NEW_USERNAME>",
	Short: "Edits the User setting of the specified entry",
	Long: `Edits the User setting of the specified entry

example: sshconf user fooServer barUser

This command operates on "UPSERT" basis,
which REPLACES the existing value with <NEW_USERNAME> when the setting exists,
or APPENDS "  User <NEW_USERNAME>" to the entry when it does not.`,
	Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		content, err := core.ReadConfigFile()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		configFileInfo, err := core.SplitEntryBlocks(content)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		hostEntries := core.MapStruct(configFileInfo)
		hostLabel := args[0]
		newValue := args[1]

		errW := core.WriteSetting(configFileInfo, hostEntries, hostLabel, "User", newValue)
		if errW != nil {
			fmt.Fprintln(os.Stderr, errW)
			return
		}

		fmt.Printf("Successfully edited %v's User to %s\n", hostLabel, newValue)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
