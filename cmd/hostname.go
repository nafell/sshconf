/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/nafell/sshconf/core"
)

// hostnameCmd represents the hostname command
var hostnameCmd = &cobra.Command{
	Use:   "hostname <ENTRY_NAME> <NEW_HOST_NAME>",
	Short: "Edits the HostName setting of the specified entry",
	Long: `Edits the HostName setting of the specified entry

example: sshconf hostname fooServer 127.0.0.1

This command operates on "UPSERT" basis,
which REPLACES the existing value with <NEW_HOST_NAME> when the setting exists,
or APPENDS "  HostName <NEW_HOST_NAME>" to the entry when it does not.`,
	Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		content, err := core.ReadConfigFile()
		if (err != nil) {
			log.Fatal(err)
			return
		}

		configFileInfo, err := core.SplitEntryBlocks(content)
		if (err != nil) {
			log.Fatal(err)
			return
		}

		hostEntries := core.MapStruct(configFileInfo)
		hostLabel := args[0]
		newValue := args[1]

		errW := core.WriteSetting(configFileInfo, hostEntries, hostLabel, "HostName", newValue)

		if (errW != nil) {
			log.Fatal(errW)
			return
		}

		fmt.Printf("Successfully edited %v's HostName to %s\n", hostLabel, newValue)
	},
}

func init() {
	rootCmd.AddCommand(hostnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
