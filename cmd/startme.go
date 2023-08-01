package main

import (
	"flag"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/josiahdenton/startme/internal/adapter/nvim"
	"github.com/josiahdenton/startme/internal/adapter/sqlite"
	"github.com/josiahdenton/startme/internal/config"
	"github.com/ktr0731/go-fuzzyfinder"
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
		// TODO - simplify the code here...
		createTemplate(name, extension)
	} else {
		searchTemplates()
	}
}

func createTemplate(name, extension string) {
	editor := nvim.New()
	fp := config.MustOpenTempFile(name, extension)
	if err := editor.Open(fp.Name()); err != nil {
		log.Fatalf("Error: failed to open new startme %v\n", err)
	}
	content := readFromTemplate(fp)
	os.Remove(fp.Name())
	saveTemplateToDb(name, content)
}

func saveTemplateToDb(name, content string) {
	db := sqlite.New(config.DbPath())
	db.Setup()
	defer db.Close()
	db.Save(name, content)
}

func readFromTemplate(fp *os.File) string {
	content, err := os.ReadFile(fp.Name())
	if err != nil {
		log.Fatalf("Error: failed to read the file: %v", err)
	}
	return string(content)
}

func searchTemplates() {
	db := sqlite.New(config.DbPath())
	db.Setup()
	defer db.Close()
	starters := db.All()

    selected, err := fuzzyfinder.Find(starters,
		func(i int) string {
			return starters[i].Title
    })
    if err != nil {
        log.Fatalf("Error: failed to select an item: %v", err)
    }
    clipboard.WriteAll(starters[selected].Content)
}
