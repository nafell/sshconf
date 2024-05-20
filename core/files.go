package core

import (
	"os"

	"github.com/pkg/errors"
)

const config_path = "/.ssh/config"

func ReadConfigFile() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home directory")
	}

	filepath := homeDir + config_path
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", errors.Wrap(err, "failed to load file: "+filepath)
	}

	return string(content), nil
}

func saveFile(content string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filepath := homeDir + config_path
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return errors.Wrap(err, "failed to access file: "+filepath)
	}
	defer file.Close()

	_, errw := file.WriteString(content)
	if errw != nil {
		return errors.Wrap(errw, "failed to write file: "+filepath)
	}

	return nil
}
