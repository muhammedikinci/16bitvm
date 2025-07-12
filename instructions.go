package main

const (
	MOV_LIT_REG = 0x10
	MOV_REG_REG = 0x11
	MOV_REG_MEM = 0x12
	MOV_MEM_REG = 0x13
	ADD_REG_REG = 0x14
	JMP_NOT_EQ  = 0x15
	PSH_LIT     = 0x16
	PSH_REG     = 0x17
	POP         = 0x18
	CAL_LIT     = 0x1A
	CAL_REG     = 0x5E
	RET         = 0x5F
	HLT         = 0xFF
)
