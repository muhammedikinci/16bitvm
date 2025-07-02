package main

type Register struct {
	name    string
	address uint8
	value   uint16
}

func NewRegisterMap() map[string]Register {
	registers := map[string]Register{
		"ip": {
			name: "ip",
		},
		"acc": {
			name: "acc",
		},
		"r1": {
			name: "r1",
		},
		"r2": {
			name: "r2",
		},
		"r3": {
			name: "r3",
		},
		"r4": {
			name: "r4",
		},
		"r5": {
			name: "r5",
		},
		"r6": {
			name: "r6",
		},
		"r7": {
			name: "r7",
		},
		"r8": {
			name: "r8",
		},
	}

	order := []string{"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8"}

	for i, key := range order {
		reg := registers[key]
		reg.address = uint8(i * 2)
		registers[key] = reg
	}

	return registers
}
