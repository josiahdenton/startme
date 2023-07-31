package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"

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
        // createTemplate(name, extension)
	} else {
        searchTemplates()
	}
}

// TOOD - this logic may make more sense as part of the starter domain struct..
// then we just pass in the interfaces...
func createTemplate(name, extension string) {
    editor := nvim.New()
    fp := config.MustOpenTempFile(name, extension)
    err := editor.Open(fp.Name())
    if err != nil {
        log.Fatalf("Error: failed to open new startme\n")
    }
    err = editor.Wait()
    if err != nil {
        log.Fatalf("Error: failed to wait for new startme to save\n")
    }
    // read content from fp
}

func readFromTemplate(fp *os.File) string {
    // looking into using strings.Builder 
    for {
        n, err := fp.Read()
        if errors.Is(err, io.EOF) && n == 0 {
            break
        } else if err != nil {

        }
    }
}

func searchTemplates() {

}
