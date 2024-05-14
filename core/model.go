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
}

func (hostEntry HostEntry) PrintPretty() {
	fmt.Printf("%s %s@%s:%s\n", hostEntry.Label, hostEntry.User, hostEntry.HostName, hostEntry.Port)
}

const config_path = "/.ssh/config"

func ReadFile() ([]HostEntry, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filepath := homeDir + config_path
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load file" + filepath)
	}

	content_str := string(content)
	//fmt.Println(content_str)
	trimmed := strings.TrimSpace(content_str)

	blocks_raw := strings.Split(trimmed, "\nHost ")
	
	if (len(blocks_raw) == 0) {
		//error
	}

	blocks_raw[0] = strings.Replace(blocks_raw[0], "Host ", "", 1)

	results := []HostEntry{}
	for _, block := range blocks_raw {
		lines := strings.Split(block, "\n")
		label := lines[0]
		hostName := ""
		user := ""
		port := "22"
		for _, line := range lines {
			trimmed_line := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed_line, "HostName") {
				hostName = strings.TrimSpace(strings.Replace(trimmed_line, "HostName", "", 1))
			} else if strings.HasPrefix(trimmed_line, "User") {
				user = strings.TrimSpace(strings.Replace(trimmed_line, "User", "", 1))
			} else if strings.HasPrefix(trimmed_line, "Port") {
				portstr := strings.TrimSpace(strings.Replace(trimmed_line, "Port", "", 1))
				//value, err := strconv.Atoi(portstr)
				//if err != nil {
				//	fmt.Println("Could not parse port setting for " + label)
				//}
				port = portstr
			}
		}

		entry := HostEntry {
			Label: label,
			HostName: hostName,
			User: user,
			Port: port,
		}
		
		results = append(results, entry)
	}

	return results, nil
}
