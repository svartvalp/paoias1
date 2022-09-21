package command

import (
	"errors"
)

type Type int

const (
	NONE Type = iota
	STORE
	POP
	IMEM
	IMEML
	MTS
	INC
	SWP
	DUP
	STR
	RTS
	SUM
	JEZ
	JGZ
	JLS
	JP
)

var strMap = map[string]Type{
	"STORE": STORE,
	"POP":   POP,
	"IMEM":  IMEM,
	"IMEML": IMEML,
	"MTS":   MTS,
	"INC":   INC,
	"SWP":   SWP,
	"DUP":   DUP,
	"STR":   STR,
	"RTS":   RTS,
	"SUM":   SUM,
	"JEZ":   JEZ,
	"JGZ":   JGZ,
	"JLS":   JLS,
	"JP":    JP,
}

func TypeFromString(in string) (Type, error) {
	if val, ok := strMap[in]; ok {
		return val, nil
	}
	return NONE, errors.New("not found")
}

type Command struct {
	Type Type
	Lit  int16
}
