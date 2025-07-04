package main

import (
	"errors"
	"fmt"
)

type CPU struct {
	memory         Memory
	registers      map[string]Register
	stackFrameSize uint16
}

func NewCPU(memory Memory) *CPU {
	registers := NewRegisterMap()

	cpu := &CPU{
		memory:    memory,
		registers: registers,
	}

	_ = cpu.SetRegisterValue("sp", uint16(len(memory)-1-1))
	_ = cpu.SetRegisterValue("fp", uint16(len(memory)-1-1))

	cpu.stackFrameSize = 0

	return cpu
}

func (c *CPU) PrintRegisters() {
	fmt.Printf("//registers//\n")
	for _, key := range RegisterOrder {
		reg := c.registers[key]
		fmt.Printf("%s 0x%04X \n", reg.name, reg.value)
	}
	fmt.Printf("\\\\registers\\\\\n")
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

	case PSH_LIT:
		value, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		c.Push(value)

	case PSH_REG:
		regAddr, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reg := c.GetRegisterByAddress(regAddr)

		c.Push(reg.value)

	case POP:
		regAddr, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = c.SetRegisterValueByAddress(regAddr, c.Pop())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case CAL_LIT:
		address, err := c.Fetch16()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		c.PushState()

		err = c.SetRegisterValue("ip", address)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case CAL_REG:
		regAddr, err := c.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		c.PushState()

		reg := c.GetRegisterByAddress(regAddr)

		err = c.SetRegisterValue("ip", uint16(reg.address))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case RET:
		c.PopState()

	}
}

func (c *CPU) Push(value uint16) {
	spAddr, err := c.GetRegister("sp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.memory.Set16(spAddr.value, value)

	err = c.SetRegisterValue("sp", spAddr.value-2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (c *CPU) Pop() uint16 {
	reg, err := c.GetRegister("sp")
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	err = c.SetRegisterValue("sp", reg.value+2)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return c.memory.Get16(reg.value + 2)
}

func (c *CPU) PushState() {
	r1, _ := c.GetRegister("r1")
	r2, _ := c.GetRegister("r2")
	r3, _ := c.GetRegister("r3")
	r4, _ := c.GetRegister("r4")
	r5, _ := c.GetRegister("r5")
	r6, _ := c.GetRegister("r6")
	r7, _ := c.GetRegister("r7")
	r8, _ := c.GetRegister("r8")
	ip, _ := c.GetRegister("ip")

	c.Push(r1.value)
	c.Push(r2.value)
	c.Push(r3.value)
	c.Push(r4.value)
	c.Push(r5.value)
	c.Push(r6.value)
	c.Push(r7.value)
	c.Push(r8.value)
	c.Push(ip.value)
	c.Push(c.stackFrameSize + 2)

	sp, _ := c.GetRegister("sp")

	_ = c.SetRegisterValue("fp", sp.value)
	c.stackFrameSize = 0
}

func (c *CPU) PopState() {
	fp, _ := c.GetRegister("fp")
	_ = c.SetRegisterValue("sp", fp.value)

	c.stackFrameSize = c.Pop()

	frameSize := c.stackFrameSize

	c.SetRegisterValue("ip", c.Pop())
	c.SetRegisterValue("r8", c.Pop())
	c.SetRegisterValue("r7", c.Pop())
	c.SetRegisterValue("r6", c.Pop())
	c.SetRegisterValue("r5", c.Pop())
	c.SetRegisterValue("r4", c.Pop())
	c.SetRegisterValue("r3", c.Pop())
	c.SetRegisterValue("r2", c.Pop())
	c.SetRegisterValue("r1", c.Pop())

	args := int(c.Pop())
	for i := 0; i < args; i++ {
		c.Pop()
	}

	c.SetRegisterValue("fp", uint16(fp.address)+frameSize)
}

func (c *CPU) Step() {
	instruction, err := c.Fetch()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.Execute(instruction)
}
