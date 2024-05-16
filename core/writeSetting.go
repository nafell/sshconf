package core

import (
	"fmt"
	"strings"
	"slices"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

/// Intended for single execution & program exit, does not edit []HostEntry
func WriteSetting(configFileInfo *ConfigFileInfo, hostEntries []HostEntry, hostLabel string, settingKey string, newValue string) error {
	edited := false
	for i, entry := range hostEntries {
		if (entry.Label != hostLabel) {
			continue
		}
		
		index := findKeyIndex(configFileInfo.Blocks[i], settingKey)

		if index != -1 { // Edit existing line
			lineNumber := entry.ConfigFilePosition + index
			pattern, err := regexp.Compile("(^[\t ]+" + settingKey + "[\t ]+)(.+)")
			if (err != nil) {
				return err
			}

			configFileInfo.Lines[lineNumber] = pattern.ReplaceAllString(configFileInfo.Lines[lineNumber], "$1 " + newValue)
		} else { // Add line
			payload := "  " + settingKey + " " + newValue
			lineNumber := entry.ConfigFilePosition + len(configFileInfo.Blocks[i])
			configFileInfo.Lines = slices.Insert(configFileInfo.Lines, lineNumber, payload)
			// else insert to original lines array, +1 to following Positions
		}
		edited = true
	}

	if edited == false {
		return fmt.Errorf("Host '%s' not found.\n", hostLabel)
	}

	// write file
	content := strings.Join(configFileInfo.Lines, "\n")
	err := saveFile(content)
	if (err != nil) {
		return err
	}

	return nil
}

func saveFile(content string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filepath := homeDir + config_path
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return errors.Wrap(err, "Failed to access file" + filepath)
	}
	defer file.Close()

	_, errw := file.WriteString(content)
	if errw != nil {
		return errors.Wrap(errw, "Failed to write file" + filepath)
	}

	return nil
}

func findKeyIndex(block []string, settingKey string) int {
	index := -1
	for i, line := range block {
		if !strings.HasPrefix((strings.TrimSpace(line)), settingKey) {
			continue
		}

		index = i
		break
	}

	return index
}
