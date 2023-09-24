package main

import (
	"log"

	"github.com/nikhilsbhat/linkerd-checker/cmd"
	"github.com/spf13/cobra/doc"
)

//go:generate go run github.com/nikhilsbhat/linkerd-checker/docs
func main() {
	commands := cmd.SetLinkerdCheckerCommands()
	if err := doc.GenMarkdownTree(commands, "doc"); err != nil {
		log.Fatal(err)
	}
}
