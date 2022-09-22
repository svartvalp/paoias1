package printer

import (
	"fmt"
	"io"
	"math/bits"

	"github.com/svartvalp/paoias1/internal/command"
	"github.com/svartvalp/paoias1/internal/state"
)

func PrintToFile(st *state.State, file io.Writer) {
	for i, m := range st.Mem {
		_, _ = fmt.Fprintf(file, "%v: ", i)
		printBytes(unsigned16(m), func(args ...any) (int, error) {
			return fmt.Fprint(file, args...)
		})
		_, _ = fmt.Fprintf(file, " (%v)\n", m)
	}
}

func Print(st *state.State) {
	fmt.Printf("Счетчик команд: ")
	printBytes(unsigned32(st.CommandCounter), fmt.Print)
	fmt.Println()
	var cmd *command.Command
	if len(st.Commands) <= int(st.CommandCounter) {
		cmdVal := st.Commands[st.CommandCounter]
		if cmdVal != 0 {
			cmd = command.ParseCommand(cmdVal)
		}
	}
	fmt.Printf("Тип команды      | Литерал")
	fmt.Println()
	if cmd != nil {
		printBytes(signed16(int16(cmd.Type)), fmt.Print)
		printBytes(signed16(cmd.Lit), fmt.Print)
	} else {
		printBytes(signed16(0), fmt.Print)
		fmt.Printf(" | ")
		printBytes(signed16(0), fmt.Print)
	}
	fmt.Println()
	fmt.Println("Регистры")
	fmt.Println("Адрес    |  Значение (+0)   |  Значение (+1)   |  Значение (+2)   |  Значение (+3)   | ")
	regRows := len(st.Registers) / 4
	for i := 0; i < regRows; i++ {
		k := i * 4
		printBytes(unsigned8(uint8(k)), fmt.Print)
		fmt.Printf(" | ")
		for j := 0; j < 4; j++ {
			printBytes(unsigned16(st.Registers[k+j]), fmt.Print)
			fmt.Printf(" | ")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Память данных")
	fmt.Println("Адрес    |  Значение (+0)   |  Значение (+1)   |  Значение (+2)   |  Значение (+3)   | ")
	memRows := len(st.Mem) / 4
	for i := 0; i < memRows; i++ {
		k := i * 4
		printBytes(unsigned8(uint8(k)), fmt.Print)
		fmt.Printf(" | ")
		for j := 0; j < 4; j++ {
			printBytes(unsigned16(st.Mem[k+j]), fmt.Print)
			fmt.Printf(" | ")
		}
		fmt.Println()
	}
	fmt.Println("Память команд")
	fmt.Println("Адрес    |           Значение (+0)          |           Значение (+1)          |           Значение (+2)          |           Значение (+3)          | ")
	cmdRows := len(st.Commands) / 4
	for i := 0; i < cmdRows; i++ {
		k := i * 4
		printBytes(unsigned8(uint8(k)), fmt.Print)
		fmt.Printf(" | ")
		for j := 0; j < 4; j++ {
			printBytes(unsigned32(st.Commands[k+j]), fmt.Print)
			fmt.Printf(" | ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printBytes(bytes []byte, print func(args ...any) (int, error)) {
	for _, b := range bytes {
		_, _ = print(b)
	}
}

func unsigned32(x uint32) []byte {
	b := make([]byte, 32)
	for i := range b {
		if bits.LeadingZeros32(x) == 0 {
			b[i] = 1
		}
		x = bits.RotateLeft32(x, 1)
	}
	return b
}

func signed32(x int32) []byte {
	return unsigned32(uint32(x))
}

func unsigned16(x uint16) []byte {
	b := make([]byte, 16)
	for i := range b {
		if bits.LeadingZeros16(x) == 0 {
			b[i] = 1
		}
		x = bits.RotateLeft16(x, 1)
	}
	return b
}

func signed16(x int16) []byte {
	return unsigned16(uint16(x))
}

func unsigned8(x uint8) []byte {
	b := make([]byte, 8)
	for i := range b {
		if bits.LeadingZeros8(x) == 0 {
			b[i] = 1
		}
		x = bits.RotateLeft8(x, 1)
	}
	return b
}

func signed8(x int8) []byte {
	return unsigned8(uint8(x))
}
