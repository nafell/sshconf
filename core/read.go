package core

import (
	"strings"

	"github.com/pkg/errors"
)

func SplitEntryBlocks(content string) (*ConfigFileInfo, error) {
	lines := strings.Split(content, "\n")

	hostEntryPositions := []int{}
	for i, line := range lines {
		if strings.HasPrefix(line, "Host") {
			hostEntryPositions = append(hostEntryPositions, i)
		}
	}

	entryLength := len(hostEntryPositions)
	if entryLength < 1 {
		return nil, errors.New("no host entry in config file")
	}
	blocks := make([][]string, 0, entryLength)
	for i := 0; i < entryLength-1; i++ {
		blocks = append(blocks, lines[hostEntryPositions[i]:hostEntryPositions[i+1]])
	}
	blocks = append(blocks, lines[hostEntryPositions[entryLength-1]:])

	return &ConfigFileInfo{
		Lines:              lines,
		Blocks:             blocks,
		HostEntryPositions: hostEntryPositions,
	}, nil
}

func MapStruct(configFileInfo *ConfigFileInfo) []HostEntry {
	results := []HostEntry{}
	for i, block := range configFileInfo.Blocks {
		label := strings.TrimSpace(strings.Replace(block[0], "Host", "", 1))

		hostName := ""
		user := ""
		port := "22"

		keys := map[string]*string{
			"HostName": &hostName,
			"User":     &user,
			"Port":     &port,
		}
		for _, line := range block {
			trimmed_line := strings.TrimSpace(line)
			for key, value := range keys {
				if strings.HasPrefix(trimmed_line, key) {
					*value = strings.TrimSpace(strings.Replace(trimmed_line, key, "", 1))
					break
				}
			}
		}

		entry := HostEntry{
			Label:              label,
			HostName:           hostName,
			User:               user,
			Port:               port,
			ConfigFilePosition: configFileInfo.HostEntryPositions[i],
		}

		results = append(results, entry)
	}

	return results
}
