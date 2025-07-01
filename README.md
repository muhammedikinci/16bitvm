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

| Opcode | Description          | Format                 |
| ------ | -------------------- | ---------------------- |
| 0x10   | Load literal into r1 | 0x10 [lit_lo] [lit_hi] |
| 0x11   | Load literal into r2 | 0x11 [lit_lo] [lit_hi] |
| 0x12   | Add r1 + r2 → acc    | 0x12 [r1_id] [r2_id]   |

All literal values are 16-bit and little-endian.

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
