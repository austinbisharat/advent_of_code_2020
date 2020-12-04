package main

import (
	"bufio"
	"log"
	"os"
)

type LineProcessor interface {
	ProcessLine(int, string) error
}

func processLinesInFile(path string, processor LineProcessor) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		processor.ProcessLine(i, line)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
