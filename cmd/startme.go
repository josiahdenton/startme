package main

import (
	"flag"
	"fmt"
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
		createTemplate(name, extension)
	} else {
		searchTemplates()
	}
}

// TOOD - this logic may make more sense as part of the starter domain struct..
// then we just pass in the interfaces...
func createTemplate(name, extension string) {
	editor := nvim.New()
	fp := config.MustOpenTempFile(name, extension)
	if err := editor.Open(fp.Name()); err != nil {
		log.Fatalf("Error: failed to open new startme %v\n", err)
	}
	// read content from fp
	// remove the temp file
	content := readFromTemplate(fp)
	os.Remove(fp.Name())

	fmt.Println(content)

    
    // open connection to DB
    // send 
}

func readFromTemplate(fp *os.File) string {
	// TODO - change to buffered read
	// looking into using strings.Builder with string()?
	// for {
	//     n, err := fp.Read()
	//     if errors.Is(err, io.EOF) && n == 0 {
	//         break
	//     } else if err != nil {
	//
	//     }
	// }
	content, err := os.ReadFile(fp.Name())
	if err != nil {
		log.Fatalf("Error: failed to read the file: %v", err)
	}
	return string(content)
}

func searchTemplates() {

}
