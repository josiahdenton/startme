package config

import (
	"fmt"
	"log"
	"os"
	"path"
)

const (
    configDir = ".startme"
    dbName = "startme.db"
)

func PreferredEditor() string {
	// TODO - add a config file to fetch this...
	return "nvim"
}

func DbPath() string {
    return path.Join(userHome(), configDir, dbName)
}

func SetupConfigDirectory() {
	if !doesConfigDirectoryExist() {
		createConfigDirectory()
	}
}

func MustOpenTempFile(file, extension string) *os.File {
    fp, err := os.CreateTemp("", fmt.Sprintf("%s.*.%s", file, extension))
    if err != nil {
        // conv to charm log
        log.Fatalf("Error: failed to open file: %v", err)
    }
	return fp
}

func doesConfigDirectoryExist() bool {
	home := userHome()
	path := path.Join(home, configDir)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func createConfigDirectory() {
	home := userHome()
	path := path.Join(home, configDir)
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
        // TODO - move to charm log and make this not panic
		log.Fatalf("Error: failed to create config directory: %v", err)
	}
}

func userHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
        // TODO - move to charm log and make this not panic
		log.Fatalf("Error: failed to get home dir: %v", err)
	}
	return home
}
