package starter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// TODO - switch to Log from charm.sh

// const dbPath string = "~/.startactivities.db"

func New(s, ext string) {
	file, err := os.CreateTemp("", fmt.Sprintf("%s.*.%s", s, ext))
	if err != nil {
		log.Fatalf("Error: failed open temp file: %v", err)
	}

	cmd := exec.Command("nvim", file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error: failed to open editor: %v", err)
	}
	cmd.Wait()

	if !isConfigFilePathReady() {
		createDbLocation()
	}

}

func isConfigFilePathReady() bool {
	return false
}

func createDbLocation() {

}

func Search() {
	// TODO open fzf
}
