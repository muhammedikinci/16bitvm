package main

import "fmt"

type Memory []uint8

func NewMemory(size int) Memory {
	return make([]uint8, size)
}

func (m Memory) Get8(address uint16) uint8 {
	return m[address]
}

func (m Memory) Get16(address uint16) uint16 {
	return uint16(m[address+1]) | uint16(m[address])<<8
}

func (m Memory) Set8(address uint16, value uint8) {
	m[address] = value
}

func (m Memory) Set16(address uint16, value uint16) {
	m[address+1] = uint8(value & 0xFF)
	m[address] = uint8((value >> 8) & 0xFF)
}

func (m Memory) PrintAt(address uint16, num int) {
	for i := 0; i < num; i++ {
		currentAddr := address + uint16(i)
		val := m.Get8(currentAddr)
		fmt.Printf("0x%04X:0x%02X ", currentAddr, val)
	}

	fmt.Print("\n")
}
