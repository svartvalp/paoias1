package runner

import (
	"github.com/svartvalp/paoias1/internal/command"
	"github.com/svartvalp/paoias1/internal/state"
)

func (r *Runner) processSTORE(c command.Command) int16 {
	err := r.st.Stack.Push(c.Lit)
	if err != nil {
		panic(err)
	}
	return c.Lit
}

func (r *Runner) processPOP(c command.Command) int16 {
	v, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	return v
}

func (r *Runner) processIMEM(c command.Command) int16 {
	addr, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	val, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	r.st.Mem[addr] = val
	return val
}
func (r *Runner) processIMEML(c command.Command) int16 {
	addr, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	r.st.Mem[addr] = c.Lit
	return c.Lit
}
func (r *Runner) processMTS(c command.Command) int16 {
	addr, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	val := r.st.Mem[addr]
	err = r.st.Stack.Push(val)
	if err != nil {
		panic(err)
	}
	return val
}
func (r *Runner) processINC(c command.Command) int16 {
	val, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	val++
	err = r.st.Stack.Push(val)
	if err != nil {
		panic(err)
	}
	return val
}
func (r *Runner) processSWP(c command.Command) int16 {
	fir, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	sec, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	err = r.st.Stack.Push(fir)
	if err != nil {
		panic(err)
	}
	err = r.st.Stack.Push(sec)
	if err != nil {
		panic(err)
	}
	return sec
}

func (r *Runner) processDUP(c command.Command) int16 {
	val, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	err = r.st.Stack.Push(val)
	if err != nil {
		panic(err)
	}
	err = r.st.Stack.Push(val)
	if err != nil {
		panic(err)
	}
	return val
}
func (r *Runner) processSTR(c command.Command) int16 {
	addr, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	val, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	r.st.Registers[addr] = val
	return val
}
func (r *Runner) processRTS(c command.Command) int16 {
	addr, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	val := r.st.Registers[addr]
	err = r.st.Stack.Push(val)
	if err != nil {
		panic(err)
	}
	return val
}
func (r *Runner) processSUM(c command.Command) int16 {
	fir, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	sec, err := r.st.Stack.Pop()
	if err != nil {
		panic(err)
	}
	sum := fir + sec
	err = r.st.Stack.Push(sum)
	if err != nil {
		panic(err)
	}
	return sum
}
func (r *Runner) processJEZ(c command.Command) int16 {
	val := r.st.Flags[state.EZ]
	if val {
		r.st.CommandCounter = int32(c.Lit)
		r.jumped = true
	}
	return 0
}
func (r *Runner) processJGZ(c command.Command) int16 {
	val := r.st.Flags[state.GZ]
	if val {
		r.st.CommandCounter = int32(c.Lit)
		r.jumped = true
	}
	return 0
}
func (r *Runner) processJLS(c command.Command) int16 {
	val := r.st.Flags[state.LZ]
	if val {
		r.st.CommandCounter = int32(c.Lit)
		r.jumped = true
	}
	return 0
}

func (r *Runner) processJP(c command.Command) int16 {
	r.st.CommandCounter = int32(c.Lit)
	r.jumped = true
	return 0
}
