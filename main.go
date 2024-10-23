package main

import (
	"log"
	"os"

	"github.com/stephenkjohnston/talesmith/core/parser"
)

func main() {
	bytes, _ := os.ReadFile("./tutorial.ta")
	parse := parser.NewParser(string(bytes))
	if err := parse.Parse(); err != nil {
		log.Fatal(err)
	}

	for _, scene := range parse.Scenes {
		scene.ToString()
	}
}
