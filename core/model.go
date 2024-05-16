package core

import (
	"fmt"
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
