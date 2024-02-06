package main

import (
	"entitx/internal"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: entitx COMMAND")
	}

	command := os.Args[1]

	if command == "generate" {
		if len(os.Args) < 4 {
			panic("usage: entitx generate ENTITY PATH_TO_PACKAGE")
		}

		internal.MustGenerate(os.Args[2], os.Args[3])
	}
}
