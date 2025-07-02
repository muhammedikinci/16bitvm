# 16bitvm

---

## Registers

| Name | Description         |
| ---- | ------------------- |
| ip   | Instruction pointer |
| acc  | Accumulator         |
| r1   | General-purpose     |
| r2   | General-purpose     |

---

## Opcodes

| Opcode | Mnemonic    | Description                     | Format                                              |
| ------ | ----------- | ------------------------------- | --------------------------------------------------- |
| 0x10   | MOV_LIT_REG | Move literal into register      | 0x10 [lit_lo] [lit_hi] [reg_id]                     |
| 0x11   | MOV_REG_REG | Move value between registers    | 0x11 [src_reg_id] [dest_reg_id]                     |
| 0x12   | MOV_REG_MEM | Move register value to memory   | 0x12 [reg_id] [addr_lo] [addr_hi]                   |
| 0x13   | MOV_MEM_REG | Load memory value into register | 0x13 [addr_lo] [addr_hi] [reg_id]                   |
| 0x14   | ADD_REG_REG | Add two registers → acc         | 0x14 [reg1_id] [reg2_id]                            |
| 0x15   | JMP_NOT_EQ  | Jump if reg != value            | 0x15 [reg_id] [val_lo] [val_hi] [addr_lo] [addr_hi] |

All literal and address values are 16-bit, little-endian.

---

## Example Memory Layout

```
Address | Byte | Meaning
--------|------|------------------------------
0x00    | 0x10 | LOAD_LITERAL r1
0x01    | 0x34 |
0x02    | 0x12 | → r1 = 0x1234 = 4660

0x03    | 0x11 | LOAD_LITERAL r2
0x04    | 0xCD |
0x05    | 0xAB | → r2 = 0xABCD = 43981

0x06    | 0x12 | ADD r1, r2
0x07    | 0x02 |
0x08    | 0x03 | → acc = r1 + r2 = 48641
```

---

## Run

```bash
go run main.go
```
