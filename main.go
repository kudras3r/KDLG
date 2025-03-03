package main

import (
	"kdlg/lexer"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: kdlg path/to/file/filename.kd")
	}
	filepath := os.Args[1]

	text, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("error when reading file: %v", err)
	}

	lexer := lexer.NewLexer(string(text))
	_ = lexer
}
