package main

type Memory []uint8

func NewMemory(size int) Memory {
	return make([]uint8, size)
}

func (m Memory) Get8(address uint16) uint8 {
	return m[address]
}

func (m Memory) Get16(address uint16) uint16 {
	return uint16(m[address]) | uint16(m[address+1])<<8
}

func (m Memory) Set8(address uint16, value uint8) {
	m[address] = value
}

func (m Memory) Set16(address uint16, value uint16) {
	m[address] = uint8(value & 0xFF)
	m[address+1] = uint8((value >> 8) & 0xFF)
}
