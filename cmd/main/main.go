package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"

	"github.com/svartvalp/paoias1/internal/parser"
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
	// fmt.Printf("%+v", state)
	for _, c := range state.Commands {
		str := strconv.FormatInt(int64(c), 2)
		fmt.Printf("%032s \n", str)
	}
	fmt.Println()
	byt := signed(2)
	for _, b := range byt {
		fmt.Print(b)
	}
}

func unsigned(x uint32) []byte {
	b := make([]byte, 32)
	for i := range b {
		if bits.LeadingZeros32(x) == 0 {
			b[i] = 1
		}
		x = bits.RotateLeft32(x, 1)
	}
	return b
}

func signed(x int32) []byte {
	return unsigned(uint32(x))
}
