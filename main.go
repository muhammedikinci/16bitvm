package main

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
	memoryMapper := NewMemoryMapper()
	screenDevice := ScreenDevice{}

	memoryMapper.Map(Region{
		Device: memory,
		Start:  0x0000,
		End:    0xFFFF,
		Remap:  false,
	})

	memoryMapper.Map(Region{
		Device: screenDevice,
		Start:  0x3000,
		End:    0x30FF,
		Remap:  true,
	})

	cpu := NewCPU(memoryMapper)

	writeCharToScreen := func(char, command, pos uint8, i int) int {
		memory[i] = MOV_LIT_REG
		memory[i+1] = command
		memory[i+2] = char
		memory[i+3] = R1

		memory[i+4] = MOV_REG_MEM
		memory[i+5] = R1
		memory[i+6] = 0x30
		memory[i+7] = pos

		return i + 8
	}

	insIndex := 0

	insIndex = writeCharToScreen(uint8(' '), 0xFF, 0, insIndex)

	for i := 0; i <= 0xFF; i++ {
		insIndex = writeCharToScreen(uint8('*'), 0x00, uint8(i), insIndex)
	}

	memory[insIndex] = HLT

	cpu.Run()
}
