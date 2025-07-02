package main

import (
	"errors"
	"fmt"
)

type CPU struct {
	memory    Memory
	registers map[string]Register
}

func NewCPU(memory Memory) *CPU {
	registers := NewRegisterMap()

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

func (c *CPU) SetRegisterValueByAddress(address uint8, value uint16) error {
	for name, reg := range c.registers {
		if address == reg.address {
			reg.value = value
			c.registers[name] = reg
			return nil
		}
	}

	return errors.New("register could not found")
}

func (c *CPU) Fetch() (uint8, error) {
	reg, err := c.GetRegister("ip")
	if err != nil {
		return 0, err
	}

	instruction := c.memory.Get8(reg.value)

	fmt.Printf("fetched 0x%02X\n", instruction)

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

	fmt.Printf("fetched 0x%04X\n", instruction)

	err = c.SetRegisterValue("ip", reg.value+2)
	if err != nil {
		return 0, err
	}

	return instruction, nil
}

func (c *CPU) Execute(instruction uint8) {
	switch instruction {
	case MOV_LIT_REG:
		literal, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		regAddress, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = c.SetRegisterValueByAddress(regAddress, literal)
		if err != nil {
			fmt.Println("register could not found by given address:", regAddress)
			return
		}

	case MOV_REG_REG:
		fromRegAddr, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		toRegAddr, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fromReg := c.GetRegisterByAddress(fromRegAddr)
		if fromReg == nil {
			fmt.Println("register could not found by given address:", fromRegAddr)
			return
		}

		err = c.SetRegisterValueByAddress(toRegAddr, fromReg.value)
		if err != nil {
			fmt.Println("register could not found by given address:", toRegAddr)
			return
		}

	case MOV_REG_MEM:
		regAddr, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		targetAddr, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		register := c.GetRegisterByAddress(regAddr)
		if register == nil {
			fmt.Println("register could not found by given address:", regAddr)
			return
		}

		// val := m.Get8(address + uint16(i))
		fmt.Printf("target 0x%04X\n", targetAddr)
		fmt.Printf("val 0x%02X\n", register.value)

		c.memory.Set16(targetAddr, register.value)

	case MOV_MEM_REG:
		memAddr, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		regAddr, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		register := c.GetRegisterByAddress(regAddr)
		if register == nil {
			fmt.Println("register could not found by given address:", regAddr)
			return
		}

		value := c.memory.Get16(memAddr)
		err = c.SetRegisterValueByAddress(regAddr, value)
		if err != nil {
			fmt.Println("register could not found by given address:", regAddr)
			return
		}

	case ADD_REG_REG:
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

	case JMP_NOT_EQ:
		checkValue, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		toAddress, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reg, err := c.GetRegister("acc")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if reg.value != checkValue {
			err = c.SetRegisterValue("ip", toAddress)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
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
