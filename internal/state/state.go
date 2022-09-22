package state

import (
	"github.com/svartvalp/paoias1/internal/stack"
)

type State struct {
	CommandCounter uint32
	Registers      []uint16
	Stack          *stack.Stack
	Commands       []uint32
	Mem            []uint16
	Flags          []bool
}

type Flag int

const (
	GZ Flag = iota
	EZ
	LZ
)

func (s *State) SetCommand(index int, val uint32) {
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
		Registers:      make([]uint16, 16),
		Stack:          stack.New(),
		Commands:       make([]uint32, 64),
		Mem:            make([]uint16, 64),
		Flags:          make([]bool, 3),
	}
}
