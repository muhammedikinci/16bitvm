# 16bitvm

---

## Registers

| Name | Description         |
| ---- | ------------------- |
| ip   | Instruction pointer |
| acc  | Accumulator         |
| r1   | General-purpose     |
| r2   | General-purpose     |
| r3   | General-purpose     |
| r4   | General-purpose     |
| r5   | General-purpose     |
| r6   | General-purpose     |
| r7   | General-purpose     |
| r8   | General-purpose     |
| sp   | Stack pointer       |
| fp   | Frame pointer       |

---

## Opcodes

| Opcode | Mnemonic    | Description                     | Format                                              |
| ------ | ----------- | ------------------------------- | --------------------------------------------------- |
| 0x10   | MOV_LIT_REG | Move literal into register      | 0x10 [lit_lo] [lit_hi] [reg_id]                     |
| 0x11   | MOV_REG_REG | Move value between registers    | 0x11 [src_reg_id] [dest_reg_id]                     |
| 0x12   | MOV_REG_MEM | Move register value to memory   | 0x12 [reg_id] [addr_lo] [addr_hi]                   |
| 0x13   | MOV_MEM_REG | Load memory value into register | 0x13 [addr_lo] [addr_hi] [reg_id]                   |
| 0x14   | ADD_REG_REG | Add two registers → acc         | 0x14 [reg1_id] [reg2_id]                            |
| 0x15   | JMP_NOT_EQ  | Jump if acc != value            | 0x15 [val_lo] [val_hi] [addr_lo] [addr_hi]          |
| 0x16   | PSH_LIT     | Push literal to stack           | 0x16 [lit_lo] [lit_hi]                              |
| 0x17   | PSH_REG     | Push register to stack          | 0x17 [reg_id]                                       |
| 0x18   | POP         | Pop from stack to register      | 0x18 [reg_id]                                       |
| 0x1A   | CAL_LIT     | Call subroutine at address      | 0x1A [addr_lo] [addr_hi]                            |
| 0x5E   | CAL_REG     | Call subroutine in register     | 0x5E [reg_id]                                       |
| 0x5F   | RET         | Return from subroutine          | 0x5F                                                |
| 0xFF   | HLT         | Halt execution                  | 0xFF                                                |

All literal and address values are 16-bit, little-endian.

---

## Features

- **Memory Mapped I/O**: Screen device mapped to 0x3000-0x30FF
- **Stack Operations**: Push/pop operations with SP and FP registers
- **Subroutine Calls**: Call and return with state preservation
- **Extensible Architecture**: Memory mapper for device integration

---

## Example Memory Layout

```
Address | Byte | Meaning
--------|------|------------------------------
0x00    | 0x10 | MOV_LIT_REG
0x01    | 0x34 | → r1 = 0x1234
0x02    | 0x12 |
0x03    | 0x04 | (register R1 address)

0x04    | 0x16 | PSH_LIT
0x05    | 0xAB | → push 0xCDAB to stack
0x06    | 0xCD |

0x07    | 0x1A | CAL_LIT
0x08    | 0x20 | → call subroutine at 0x0020
0x09    | 0x00 |

0x0A    | 0xFF | HLT
```

---

## Run

```bash
go run main.go
```
