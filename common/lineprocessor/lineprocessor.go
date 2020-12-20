package lineprocessor

import (
	"bufio"
	"log"
	"os"
)

type LineProcessor interface {
	ProcessLine(int, string) error
}

func ProcessLinesInFile(path string, processor LineProcessor) {
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

type LineMultiProcessor struct {
	Processors []LineProcessor
}

func (m *LineMultiProcessor) ProcessLine(i int, line string) error {
	for _, processor := range m.Processors {
		err := processor.ProcessLine(i, line)
		if err != nil {
			return err
		}
	}
	return nil
}
