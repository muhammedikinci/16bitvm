package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	IP  = 0
	ACC = 2
	R1  = 4
	R2  = 6
	R3  = 8
	R4  = 10
	R5  = 12
	R6  = 14
	R7  = 16
	R8  = 18
	SP  = 20
	FP  = 22
)

func main() {
	memory := NewMemory(256 * 256)
	cpu := NewCPU(memory)

	memory[0] = MOV_LIT_REG
	memory[1] = 0x51
	memory[2] = 0x51
	memory[3] = R1

	memory[4] = MOV_LIT_REG
	memory[5] = 0x42
	memory[6] = 0x42
	memory[7] = R2

	memory[8] = PSH_REG
	memory[9] = R1

	memory[10] = PSH_REG
	memory[11] = R2

	memory[12] = POP
	memory[13] = R1

	memory[14] = POP
	memory[15] = R2

	reg, _ := cpu.GetRegister("ip")
	memory.PrintAt(uint16(reg.value))
	memory.PrintAt(0xFFFF - 1 - 6)
	fmt.Print("\n")

	for {
		fmt.Println("press enter for next step")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		cpu.Step()

		reg, _ = cpu.GetRegister("ip")
		memory.PrintAt(uint16(reg.value))
		memory.PrintAt(0xFFFF - 1 - 6)
		fmt.Print("\n")
		cpu.PrintRegisters()
	}
}
