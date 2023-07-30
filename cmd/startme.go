package main

import (
	"flag"

	"github.com/josiahdenton/startme/internal/starter"
)

func main() {
	name := flag.String("new", "", "set the name of the template you wish to save")
	ext := flag.String("ext", ".md", "set the extension of the template, default value is .md (markdown)")
	// if no flags are set, assume search

	if len(*name) > 0 {
		starter.New(*name, *ext)
	} else {
		starter.Search()
	}
}
