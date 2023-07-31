package main

import (
	"flag"
	"log"

	"github.com/josiahdenton/startme/internal/adapter/nvim"
	"github.com/josiahdenton/startme/internal/config"
)

func main() {
    config.SetupConfigDirectory()
	name := flag.String("n", "", "set the name of the template you wish to save")
	ext := flag.String("e", ".md", "set the extension of the template, default value is .md (markdown)")

    flag.Parse()

    run(*name, *ext)

}

func run(name, extension string) {
	if len(name) > 0 {
        newTemplate(name, extension)
	} else {
        searchTemplates()
	}
}

// TOOD - this logic may make more sense as part of the starter domain struct..
// then we just pass in the interfaces...
func newTemplate(name, extension string) {
    editor := nvim.New()
    err := editor.Open(config.TempFilePath(name, extension))
    if err != nil {
        log.Fatalf("Error: failed to open new startme\n")
    }
    err = editor.Wait()
    if err != nil {
        log.Fatalf("Error: failed to wait for new startme to save\n")
    }
}

func searchTemplates() {

}
