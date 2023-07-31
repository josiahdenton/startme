package config

import (
	"log"
	"os"
	"path"
)

const configDir = ".startme"

func DoesConfigDirectoryExist() bool {
    home := userHome()
	path := path.Join(home, configDir)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
    }
	return info.IsDir()
}

func CreateConfigDirectory() {
    home := userHome()
	path := path.Join(home, configDir)
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatalf("Error: failed to create config directory: %v", err)
	}
}

func userHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: failed to get home dir: %v", err)
	}
    return home
}
