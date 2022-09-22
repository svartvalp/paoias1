package runner

import (
	"strconv"

	"github.com/svartvalp/paoias1/internal/command"
	"github.com/svartvalp/paoias1/internal/state"
)

type Runner struct {
	st      *state.State
	procMap map[command.Type]func(c command.Command) int16
	jumped  bool
}

func (r *Runner) Run() {
	for r.NextStep() {

	}
}

func (r *Runner) NextStep() bool {
	cmd, ok := r.PrepareCommand()
	if !ok {
		return false
	}
	r.jumped = false
	procFunc, ok := r.procMap[cmd.Type]
	if !ok {
		panic("not found type " + strconv.Itoa(int(cmd.Type)))
	}
	result := procFunc(*cmd)
	r.st.RestoreFlags()
	if result == 0 {
		r.st.SetFlag(state.EZ, true)
	}
	if result > 0 {
		r.st.SetFlag(state.GZ, true)
	}
	if result < 0 {
		r.st.SetFlag(state.LZ, true)
	}
	if !r.jumped {
		r.st.CommandCounter++
	}
	return true
}

func (r *Runner) PrepareCommand() (*command.Command, bool) {
	if len(r.st.Commands) <= int(r.st.CommandCounter) {
		return nil, false
	}
	cmdVal := r.st.Commands[r.st.CommandCounter]
	if cmdVal == 0 {
		return nil, false
	}
	return command.ParseCommand(cmdVal), true
}

func (r *Runner) loadProcMap() {
	r.procMap = map[command.Type]func(c command.Command) int16{
		command.STORE: r.processSTORE,
		command.POP:   r.processPOP,
		command.IMEM:  r.processIMEM,
		command.IMEML: r.processIMEML,
		command.MTS:   r.processMTS,
		command.INC:   r.processINC,
		command.SWP:   r.processSWP,
		command.DUP:   r.processDUP,
		command.STR:   r.processSTR,
		command.RTS:   r.processRTS,
		command.SUM:   r.processSUM,
		command.JEZ:   r.processJEZ,
		command.JGZ:   r.processJGZ,
		command.JLS:   r.processJLS,
		command.JP:    r.processJP,
	}
}

func New(st *state.State) *Runner {
	r := &Runner{
		st: st,
	}
	r.loadProcMap()
	return r
}
