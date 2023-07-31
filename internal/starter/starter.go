package starter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/josiahdenton/startme/internal/config"
)

// TODO - switch to Log from charm.sh

const (
	dbName    = "templates.db"
)

func New(s, ext string) {
	// file, err := os.CreateTemp("", fmt.Sprintf("%s.*.%s", s, ext))
	// if err != nil {
	// 	log.Fatalf("Error: failed open temp file: %v", err)
	// }
	//    openTempWithEditor("nvim", file.Name())
	if !config.DoesConfigDirectoryExist() {
		config.CreateConfigDirectory()
	}
	// now I want to just .... connect to sqlite db...
	// exec create table if not exists
	fmt.Println("SUCCESS!")
}

func openTempWithEditor(editor, path string) {
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: failed to open editor: %v", err)
	}
	cmd.Wait()
}


func Search() {
	// TODO open fzf
}
