package main

import (
	"flag"
	"fmt"

	"github.com/josiahdenton/startme/internal/orchestrator"
)

func main() {
	name := flag.String("n", "", "set the name of the template you wish to save")
	ext := flag.String("e", ".md", "set the extension of the template, default value is .md (markdown)")
    flag.Parse()

	if len(*name) > 0 {
        orchestrator.SaveStarter(*name, *ext)
	} else {
        orchestrator.Search()
	}
}
