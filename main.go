package main

import (
	_ "github.com/zellyn/kooky/browser/all" // register cookie store finders!
	"konfig/cmd/konfig"
)

func main() {
	konfig.Execute()
}
