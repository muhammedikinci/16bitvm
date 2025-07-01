package main

import (
	"errors"
	"fmt"
)

type CPU struct {
	memory    Memory
	registers map[string]Register
}

type Register struct {
	address uint8
	value   uint16
}

func NewCPU(memory Memory) *CPU {
	registers := newRegisterMap()

	return &CPU{
		memory:    memory,
		registers: registers,
	}
}

func (c *CPU) GetRegisterByAddress(address uint8) *Register {
	for _, reg := range c.registers {
		if address == reg.address {
			return &reg
		}
	}

	return nil
}

func (c *CPU) GetRegister(name string) (Register, error) {
	reg, ok := c.registers[name]
	if !ok {
		return Register{}, errors.New("register could not found")
	}

	return reg, nil
}

func (c *CPU) SetRegisterValue(name string, value uint16) error {
	reg, ok := c.registers[name]
	if !ok {
		return errors.New("register could not found")
	}

	reg.value = value
	c.registers[name] = reg

	return nil
}

func (c *CPU) Fetch() (uint8, error) {
	reg, err := c.GetRegister("ip")
	if err != nil {
		return 0, err
	}

	instruction := c.memory.Get8(reg.value)

	err = c.SetRegisterValue("ip", reg.value+1)
	if err != nil {
		return 0, err
	}

	return instruction, nil
}

func (c *CPU) Fetch16() (uint16, error) {
	reg, err := c.GetRegister("ip")
	if err != nil {
		return 0, err
	}

	instruction := c.memory.Get16(reg.value)

	err = c.SetRegisterValue("ip", reg.value+2)
	if err != nil {
		return 0, err
	}

	return instruction, nil
}

func (c *CPU) Execute(instruction uint8) {
	switch instruction {
	// move literal to r1
	case 0x10:
		literal, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = c.SetRegisterValue("r1", literal)
		if err != nil {
			fmt.Println(err.Error())
		}

	// move literal to r2
	case 0x11:
		literal, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = c.SetRegisterValue("r2", literal)
		if err != nil {
			fmt.Println(err.Error())
		}

	// add two register values
	case 0x12:
		regAddress1, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		regAddress2, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reg1 := c.GetRegisterByAddress(regAddress1)
		if reg1 == nil {
			fmt.Println("register could not found by given address:", regAddress1)
			return
		}

		reg2 := c.GetRegisterByAddress(regAddress2)
		if reg2 == nil {
			fmt.Println("register could not found by given address:", regAddress2)
			return
		}

		err = c.SetRegisterValue("acc", reg1.value+reg2.value)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		return
	}
}

func (c *CPU) Step() {
	instruction, err := c.Fetch()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.Execute(instruction)
}

func newRegisterMap() map[string]Register {
	registers := map[string]Register{
		"ip":  {},
		"acc": {},
		"r1":  {},
		"r2":  {},
		// "r3":  {},
		// "r4":  {},
		// "r5":  {},
		// "r6":  {},
		// "r7":  {},
		// "r8":  {},
	}

	order := []string{"ip", "acc", "r1", "r2"}

	for i, key := range order {
		reg := registers[key]
		reg.address = uint8(i * 2)
		registers[key] = reg
	}

	return registers
}
