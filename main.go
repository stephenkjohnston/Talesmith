package main

import (
	"os"

	"github.com/stephenkjohnston/talesmith/core/parser"
)

func main() {
	bytes, _ := os.ReadFile("./tutorial.ta")
	parse := parser.NewParser(string(bytes))
	parse.Parse()
	// for _, token := range parse.Tokens {
	// 	token.Debug()
	// }
}
