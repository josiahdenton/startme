package orchestrator

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/josiahdenton/startme/internal/config"
)

func SaveStarter(name, ext string) {
	// file, err := os.CreateTemp("", fmt.Sprintf("%s.*.%s", s, ext))
	// if err != nil {
	// 	log.Fatalf("Error: failed open temp file: %v", err)
	// }
	//    openTempWithEditor("nvim", file.Name())
    // TODO - this can be moved to the main go file...
	if !config.DoesConfigDirectoryExist() {
		config.CreateConfigDirectory()
	}
	// now I want to just .... connect to sqlite db...
	// exec create table if not exists
	fmt.Println("SUCCESS!")
}

func openTempWithEditor(editor, path string) {

}


func Search() {
	// TODO open fzf
}
