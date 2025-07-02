package main

import "fmt"

const (
	IP  = 0
	ACC = 2
	R1  = 4
	R2  = 6
)

func main() {
	memory := NewMemory(256 * 256)
	cpu := NewCPU(memory)

	memory[0] = MOV_LIT_REG
	memory[1] = 0x12
	memory[2] = 0x34
	memory[3] = R1

	memory[4] = MOV_LIT_REG
	memory[5] = 0xAB
	memory[6] = 0xCD
	memory[7] = R2

	memory[8] = ADD_REG_REG
	memory[9] = R1
	memory[10] = R2

	memory[11] = MOV_REG_MEM
	memory[12] = ACC
	memory[13] = 0x01
	memory[14] = 0x00

	reg, _ := cpu.GetRegister("ip")
	memory.PrintAt(uint16(reg.value))
	memory.PrintAt(0x0100)
	fmt.Print("\n")

	cpu.Step()
	reg, _ = cpu.GetRegister("ip")
	memory.PrintAt(uint16(reg.value))
	memory.PrintAt(0x0100)
	fmt.Print("\n")

	cpu.Step()
	reg, _ = cpu.GetRegister("ip")
	memory.PrintAt(uint16(reg.value))
	memory.PrintAt(0x0100)
	fmt.Print("\n")

	cpu.Step()
	reg, _ = cpu.GetRegister("ip")
	memory.PrintAt(uint16(reg.value))
	memory.PrintAt(0x0100)
	fmt.Print("\n")

	cpu.Step()
	reg, _ = cpu.GetRegister("ip")
	memory.PrintAt(uint16(reg.value))
	memory.PrintAt(0x0100)
}
