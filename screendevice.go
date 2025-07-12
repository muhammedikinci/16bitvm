package main

import (
	"fmt"
)

type ScreenDevice struct{}

func (s ScreenDevice) Get16(address uint16) uint16 { return 0 }
func (s ScreenDevice) Get8(address uint16) uint8   { return 0 }
func (s ScreenDevice) Set16(address, data uint16) {
	command := (0xFF00 & data) >> 8
	characterValue := data & 0x00FF
	if command == 0xFF {
		s.erase()
	}
	x := int((address % 16) + 1)
	y := int((address / 16) + 1)
	s.print(x*2, y)
	character := string(rune(characterValue))
	fmt.Println(character)
}

func (s ScreenDevice) Set8(address uint16, data uint8) {
}

func (s ScreenDevice) print(x, y int) {
	fmt.Printf("\x1b[%d;%dH", y, x)
}

func (s ScreenDevice) erase() {
	fmt.Print("\x1b[2J\n")
}
