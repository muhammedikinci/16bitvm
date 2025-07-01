package main

import "fmt"

func main() {
	memory := NewMemory(16)
	cpu := NewCPU(memory)

	memory[0] = 0x10
	memory[1] = 0x00
	memory[2] = 0x01

	memory[3] = 0x11
	memory[4] = 0x00
	memory[5] = 0x02

	memory[6] = 0x12
	memory[7] = 4
	memory[8] = 6

	fmt.Println(cpu)
	cpu.Step()
	fmt.Println(cpu)
	cpu.Step()
	fmt.Println(cpu)
	cpu.Step()
	fmt.Println(cpu)

	r1reg, _ := cpu.GetRegister("r1")
	fmt.Println("r1 value:", r1reg.value>>8)

	r2reg, _ := cpu.GetRegister("r2")
	fmt.Println("r2 value:", r2reg.value>>8)

	accreg, _ := cpu.GetRegister("acc")
	fmt.Println("acc value:", accreg.value>>8)
}
