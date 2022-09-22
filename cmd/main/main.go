package main

import (
	"bufio"
	"os"

	"github.com/svartvalp/paoias1/internal/parser"
	"github.com/svartvalp/paoias1/internal/printer"
	"github.com/svartvalp/paoias1/internal/runner"
)

func main() {
	file, err := os.Open("./prog.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()
	scanner := bufio.NewScanner(file)
	pars := parser.New()
	for scanner.Scan() {
		str := scanner.Text()
		err = pars.Scan(str)
		if err != nil {
			panic(err)
		}
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	state, err := pars.ProcessState()
	if err != nil {
		panic(err)
	}
	run := runner.New(state)
	run.Run()
	printer.Print(state)
	wrFile, err := os.Create("./result.txt")
	if err != nil {
		panic(err)
	}
	printer.PrintToFile(state, wrFile)
}
