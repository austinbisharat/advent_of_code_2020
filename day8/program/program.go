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
	case STTerm:
		fmt.Printf("Program terminated. Counter: %d\n", r.accumulator)
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

func (p *Program) RemoveInfiniteLoop() error {
	p.statements = append(p.statements, statement{
		statementType: STTerm,
		argument:      0,
	})

	terminatingStatements := p.findTerminatingStatements()

	var ip int
	for !terminatingStatements[p.statements[ip].getNextIP(ip)] {
		ip = p.statements[ip].getNextIP(ip)
		statement := p.statements[ip]

		if statement.statementType == STJmp {
			statement.statementType = STNop
		} else if statement.statementType == STNop {
			statement.statementType = STJmp
		}

		if terminatingStatements[statement.getNextIP(ip)] {
			p.statements[ip] = statement
		}
	}

	return nil
}

func (p *Program) buildProgramGraph() (statementIdxToPredecessors [][]int) {
	statementIdxToPredecessors = make([][]int, len(p.statements))

	for ip, statement := range p.statements {
		nextIP := statement.getNextIP(ip)
		if nextIP >= len(p.statements) {
			continue
		}
		statementIdxToPredecessors[nextIP] = append(statementIdxToPredecessors[nextIP], ip)
	}
	return statementIdxToPredecessors
}

func (p *Program) findTerminatingStatements() map[int]bool {
	statementIdxToPredecessors := p.buildProgramGraph()

	terminatingStatementIdx := len(p.statements) - 1
	discovered := map[int]bool{terminatingStatementIdx: true}
	toVisit := []int{terminatingStatementIdx}

	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = toVisit[1:]
		for _, predecessor := range statementIdxToPredecessors[cur] {
			if discovered[predecessor] {
				continue
			}

			discovered[predecessor] = true
			toVisit = append(toVisit, predecessor)
		}
	}

	return discovered
}

func (p *Program) ProcessLine(_ int, line string) error {
	var s statement
	err := s.parse(line)
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
	STAcc  StatementType = "acc"
	STJmp  StatementType = "jmp"
	STNop  StatementType = "nop"
	STTerm StatementType = "term"
)

type statement struct {
	statementType StatementType
	argument      int
}

var statementRegex = regexp.MustCompile("^(acc|jmp|nop) ([+-][0-9]+)$")

func (s *statement) parse(statement string) error {
	submatches := statementRegex.FindStringSubmatch(statement)
	if len(submatches) != 3 {
		return fmt.Errorf("invalid statement: \"%s\"did not match regex %v", statement, statementRegex)
	}

	var err error
	s.statementType = StatementType(submatches[1])
	s.argument, err = strconv.Atoi(submatches[2])
	return err

}

func (s *statement) getNextIP(currentIP int) int {
	if s.statementType == STJmp {
		return currentIP + s.argument
	}
	return currentIP + 1
}
