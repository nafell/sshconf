package core

import (
	"fmt"
	"os"
	"strings"
	//"strconv"

	"github.com/pkg/errors"
)

type HostEntry struct {
	Label string
	HostName string
	User string
	Port string
	ConfigFilePosition int
}

type ConfigFileInfo struct {
	Lines []string
	Blocks [][]string
	HostEntryPositions []int
}

func (hostEntry HostEntry) PrintPretty() {
	fmt.Printf("%s %s@%s:%s\n", hostEntry.Label, hostEntry.User, hostEntry.HostName, hostEntry.Port)
}

const config_path = "/.ssh/config"

func ReadConfigFile() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filepath := homeDir + config_path
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", errors.Wrap(err, "failed to load file" + filepath)
	}

	return string(content), nil
}

func SplitEntryBlocks(content string) (*ConfigFileInfo, error) {
	lines := strings.Split(content, "\n")
	
	hostEntryPositions := []int{}
	for i, line := range lines {
		if (strings.HasPrefix(line, "Host")) {
			hostEntryPositions = append(hostEntryPositions,i)
		}
	}

	entryLength := len(hostEntryPositions)
	if entryLength < 1 {
		return nil, errors.New("No host entry in config file.")
	}
	blocks := make([][]string, 0, entryLength)
	for i := 0; i < entryLength-1; i++ {
		blocks = append(blocks, lines[hostEntryPositions[i] : hostEntryPositions[i+1]])
	}
	blocks = append(blocks, lines[hostEntryPositions[entryLength-1]:])

	return &ConfigFileInfo{
		Lines: lines,
		Blocks: blocks,
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

		keys := map[string]*string {
			"HostName": &hostName, 
			"User": &user, 
			"Port": &port,
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

		entry := HostEntry {
			Label: label,
			HostName: hostName,
			User: user,
			Port: port,
			ConfigFilePosition: configFileInfo.HostEntryPositions[i],
		}
		
		results = append(results, entry)
	}

	return results
}
