package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/svartvalp/paoias1/internal/command"
	"github.com/svartvalp/paoias1/internal/state"
)

type Parser struct {
	parsed []ParsedCommand
	labels map[string]int32
	st     *state.State
}

type ParsedCommand struct {
	Cmd         command.Type
	Lit         int16
	Label       string
	IsLabeled   bool
	IsJumpLabel bool
}

func (p *Parser) Scan(in string) error {
	split := strings.Split(strings.TrimSpace(in), " ")
	split = filterEmpty(split)
	if len(split) == 0 {
		return nil
	}
	if split[0][len(split[0])-1] == ':' {
		return p.parseLabel(split)
	}
	cmd, err := command.TypeFromString(split[0])
	if err != nil {
		return err
	}
	if len(split) > 1 {
		sec := split[1]
		if isDigit(sec) {
			return p.parseWithLiteral(cmd, split[1])
		} else {
			return p.parseWithLitLabel(cmd, split[1])
		}
	}
	pars := ParsedCommand{
		Cmd: cmd,
	}
	p.parsed = append(p.parsed, pars)
	return nil
}

func filterEmpty(split []string) []string {
	res := make([]string, 0, len(split))
	for _, s := range split {
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}
		res = append(res, s)
	}
	return res
}

func isDigit(s string) bool {
	for i, c := range s {
		if i == 0 && c == '-' {
			continue
		}
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func (p *Parser) ProcessState() (*state.State, error) {
	for i, c := range p.parsed {
		if c.IsJumpLabel {
			lit, ok := p.labels[c.Label]
			if !ok {
				return nil, fmt.Errorf("label %v not found", c.Label)
			}
			c.Lit = int16(lit)
		}
		mCmd := int32(c.Cmd)<<16 + int32(c.Lit)
		p.st.SetCommand(i, mCmd)
	}

	return p.st, nil
}

func (p *Parser) parseLabel(splitted []string) error {
	label := splitted[0][:len(splitted[0])-1]
	cmd, err := command.TypeFromString(splitted[1])
	if err != nil {
		return err
	}
	pars := ParsedCommand{
		Cmd:         cmd,
		Lit:         0,
		Label:       label,
		IsLabeled:   true,
		IsJumpLabel: false,
	}
	if len(splitted) > 2 {
		lit, err := strconv.Atoi(splitted[2])
		if err != nil {
			return err
		}
		pars.Lit = int16(lit)
	}
	p.parsed = append(p.parsed, pars)
	p.labels[label] = int32(len(p.parsed) - 1)
	return nil
}

func (p *Parser) parseWithLiteral(cmd command.Type, s string) error {
	lit, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	pars := ParsedCommand{
		Cmd:         cmd,
		Lit:         int16(lit),
		Label:       "",
		IsLabeled:   false,
		IsJumpLabel: false,
	}
	p.parsed = append(p.parsed, pars)
	return nil
}

func (p *Parser) parseWithLitLabel(cmd command.Type, s string) error {
	pars := ParsedCommand{
		Cmd:         cmd,
		Lit:         0,
		Label:       s,
		IsLabeled:   false,
		IsJumpLabel: true,
	}
	p.parsed = append(p.parsed, pars)
	return nil
}

func New() *Parser {
	return &Parser{
		parsed: make([]ParsedCommand, 0),
		labels: make(map[string]int32),
		st:     state.New(),
	}
}
