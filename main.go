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
)

func main() {
	memory := NewMemory(256 * 256)
	cpu := NewCPU(memory)

	memory[0] = MOV_MEM_REG
	memory[1] = 0x01
	memory[2] = 0x00
	memory[3] = R1

	memory[4] = MOV_LIT_REG
	memory[5] = 0x00
	memory[6] = 0x01
	memory[7] = R2

	memory[8] = ADD_REG_REG
	memory[9] = R1
	memory[10] = R2

	memory[11] = MOV_REG_MEM
	memory[12] = ACC
	memory[13] = 0x01
	memory[14] = 0x00

	memory[15] = JMP_NOT_EQ
	memory[16] = 0x00
	memory[17] = 0x03
	memory[18] = 0x00
	memory[19] = 0x00

	reg, _ := cpu.GetRegister("ip")
	memory.PrintAt(uint16(reg.value))
	memory.PrintAt(0x0100)
	fmt.Print("\n")

	for {
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		cpu.Step()

		reg, _ = cpu.GetRegister("ip")
		memory.PrintAt(uint16(reg.value))
		memory.PrintAt(0x0100)
		fmt.Print("\n")
		cpu.PrintRegisters()
	}
}
