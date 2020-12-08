package program

import (
	"advent/common/lineprocessor"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Runtime struct {
	program            *Program
	accumulator        int
	instructionPointer int
	visitedIPs         []bool

	hasSelfCorrected bool
}

func NewRuntime(program *Program) *Runtime {
	return &Runtime{
		program:            program,
		accumulator:        0,
		instructionPointer: 0,
		visitedIPs:         make([]bool, len(program.statements)),

		hasSelfCorrected: false,
	}
}

// returns true iff the program terminated because we reached the end
func (r *Runtime) RunUntilInfiniteLoop() bool {
	for r.instructionPointer < len(r.program.statements) && !r.visitedIPs[r.instructionPointer] {
		r.visitedIPs[r.instructionPointer] = true
		r.runStep()
	}
	return r.instructionPointer == len(r.program.statements)
}

func (r *Runtime) RunSelfCorrecting() int {
	runtimesToTry := []*Runtime{r}

	for len(runtimesToTry) > 0 {
		runtime := runtimesToTry[0]
		runtimesToTry = runtimesToTry[1:]

		if runtime.hasSelfCorrected && runtime.RunUntilInfiniteLoop() {
			return runtime.accumulator
		} else if !runtime.hasSelfCorrected {
			runtimesToTry = append(runtimesToTry, r.runStepSelfCorrecting()...)
		}
	}

	return -1
}

func (r *Runtime) runStep() {
	statement := r.program.statements[r.instructionPointer]
	switch statement.statementType {
	case STAcc:
		r.accumulator += statement.argument
		r.instructionPointer++
	case STJmp:
		r.instructionPointer += statement.argument
	case STNop:
		r.instructionPointer++
	default:
		log.Fatalf("Unknown statemnt type: %s", statement.statementType)
	}
}

func (r *Runtime) runStepSelfCorrecting() []*Runtime {
	statement := r.program.statements[r.instructionPointer]

	runtimes := []*Runtime{r}
	switch statement.statementType {
	case STAcc:
		r.accumulator += statement.argument
		r.instructionPointer++
	case STJmp:
		cpy := r.Copy()
		cpy.instructionPointer++
		cpy.hasSelfCorrected = true
		runtimes = append(runtimes, cpy)

		r.instructionPointer += statement.argument
	case STNop:
		cpy := r.Copy()
		cpy.instructionPointer += statement.argument
		cpy.hasSelfCorrected = true
		runtimes = append(runtimes, cpy)

		r.instructionPointer++
	default:
		log.Fatalf("Unknown statemnt type: %s", statement.statementType)
	}
	return runtimes
}

func (r *Runtime) GetAccumulator() int {
	return r.accumulator
}

func (r *Runtime) Copy() *Runtime {
	cpy := *r
	cpy.visitedIPs = make([]bool, len(cpy.visitedIPs))
	copy(cpy.visitedIPs, r.visitedIPs)
	return &cpy
}

type Program struct {
	statements []statement
}

func (p *Program) ProcessLine(_ int, line string) error {
	var s statement
	err := s.Parse(line)
	if err != nil {
		return err
	}

	p.statements = append(p.statements, s)
	return nil
}

func LoadProgram(filepath string) *Program {
	var program Program
	lineprocessor.ProcessLinesInFile(filepath, &program)
	return &program
}

type StatementType string

const (
	STAcc StatementType = "acc"
	STJmp StatementType = "jmp"
	STNop StatementType = "nop"
)

type statement struct {
	statementType StatementType
	argument      int
}

var statementRegex = regexp.MustCompile("^(acc|jmp|nop) ([+-][0-9]+)$")

func (s *statement) Parse(statement string) error {
	submatches := statementRegex.FindStringSubmatch(statement)
	if len(submatches) != 3 {
		return fmt.Errorf("invalid statement: \"%s\"did not match regex %v", statement, statementRegex)
	}

	var err error
	s.statementType = StatementType(submatches[1])
	s.argument, err = strconv.Atoi(submatches[2])
	return err

}
