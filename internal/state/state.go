package state

import (
	"github.com/svartvalp/paoias1/internal/stack"
)

type State struct {
	CommandCounter int32
	Registers      []int16
	Stack          *stack.Stack
	Commands       []int32
	Mem            []int16
	Flags          []bool
}

type Flag int

const (
	GZ Flag = iota
	EZ
	LZ
)

func (s *State) SetCommand(index int, val int32) {
	s.Commands[index] = val
}

func (s *State) SetFlag(f Flag, val bool) {
	s.Flags[f] = val
}

func (s *State) RestoreFlags() {
	s.Flags = make([]bool, 3)
}

func New() *State {
	return &State{
		CommandCounter: 0,
		Registers:      make([]int16, 4),
		Stack:          stack.New(),
		Commands:       make([]int32, 64),
		Mem:            make([]int16, 64),
		Flags:          make([]bool, 3),
	}
}
