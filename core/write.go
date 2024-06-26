package core

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	//"github.com/pkg/errors"
)

// / Intended for single execution & program exit, does not edit []HostEntry
func WriteSetting(configFileInfo *ConfigFileInfo, hostEntries []HostEntry, hostLabel string, settingKey string, newValue string) error {
	edited := false
	for i, entry := range hostEntries {
		if entry.Label != hostLabel {
			continue
		}

		index := findKeyIndex(configFileInfo.Blocks[i], settingKey)

		if index != -1 { // Edit existing line
			lineNumber := entry.ConfigFilePosition + index
			pattern, err := regexp.Compile("(^[\t ]+" + settingKey + "[\t ]+)(.+)")
			if err != nil {
				return err
			}

			configFileInfo.Lines[lineNumber] = pattern.ReplaceAllString(configFileInfo.Lines[lineNumber], "$1"+newValue)
		} else { // Add line
			payload := "  " + settingKey + " " + newValue
			lineNumber := entry.ConfigFilePosition + len(configFileInfo.Blocks[i])
			configFileInfo.Lines = slices.Insert(configFileInfo.Lines, lineNumber, payload)
			// else insert to original lines array, +1 to following Positions
		}
		edited = true
	}

	if !edited {
		return fmt.Errorf("host '%s' not found", hostLabel)
	}

	// write file
	content := strings.Join(configFileInfo.Lines, "\n")
	err := saveFile(content)
	if err != nil {
		return err
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
