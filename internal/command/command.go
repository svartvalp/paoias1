package command

func ParseCommand(cmdVal uint32) *Command {
	cmd := cmdVal >> 16
	lit := int16(cmdVal)
	return &Command{
		Type: Type(cmd),
		Lit:  lit,
	}
}
