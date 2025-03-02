package main

import (
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

	lexer := New(string(text))

	parser := yyNewParser()

	result := parser.Parse(lexer)
	log.Printf("parser result: %d", result)
}
