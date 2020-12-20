package parser

type Parser interface {
	IsValid(line string) bool
}
