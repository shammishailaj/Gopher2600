// generated code - do not change

package definitions

// GetInstructionDefinitions returns the opcode table for the MC6502
func GetInstructionDefinitions() ([]*InstructionDefinition, error) {
return []*InstructionDefinition{
&InstructionDefinition{ObjectCode:0x0, Mnemonic:"BRK", Bytes:1, Cycles:7, AddressingMode:0, PageSensitive:false, Effect:5},
&InstructionDefinition{ObjectCode:0x1, Mnemonic:"ORA", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:0},
nil,
nil,
&InstructionDefinition{ObjectCode:0x4, Mnemonic:"dop", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x5, Mnemonic:"ORA", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x6, Mnemonic:"ASL", Bytes:2, Cycles:5, AddressingMode:4, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x8, Mnemonic:"PHP", Bytes:1, Cycles:3, AddressingMode:0, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x9, Mnemonic:"ORA", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa, Mnemonic:"ASL", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xc, Mnemonic:"skw", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xd, Mnemonic:"ORA", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xe, Mnemonic:"ASL", Bytes:3, Cycles:6, AddressingMode:3, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x10, Mnemonic:"BPL", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0x11, Mnemonic:"ORA", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x15, Mnemonic:"ORA", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x16, Mnemonic:"ASL", Bytes:2, Cycles:6, AddressingMode:10, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x18, Mnemonic:"CLC", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x19, Mnemonic:"ORA", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x1d, Mnemonic:"ORA", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0x1e, Mnemonic:"ASL", Bytes:3, Cycles:7, AddressingMode:8, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x20, Mnemonic:"JSR", Bytes:3, Cycles:6, AddressingMode:3, PageSensitive:false, Effect:4},
&InstructionDefinition{ObjectCode:0x21, Mnemonic:"AND", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:0},
nil,
nil,
&InstructionDefinition{ObjectCode:0x24, Mnemonic:"BIT", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x25, Mnemonic:"AND", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x26, Mnemonic:"ROL", Bytes:2, Cycles:5, AddressingMode:4, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x28, Mnemonic:"PLP", Bytes:1, Cycles:4, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x29, Mnemonic:"AND", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x2a, Mnemonic:"ROL", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0x2c, Mnemonic:"BIT", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x2d, Mnemonic:"AND", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x2e, Mnemonic:"ROL", Bytes:3, Cycles:6, AddressingMode:3, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x30, Mnemonic:"BMI", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0x31, Mnemonic:"AND", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x35, Mnemonic:"AND", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x36, Mnemonic:"ROL", Bytes:2, Cycles:6, AddressingMode:10, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x38, Mnemonic:"SEC", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x39, Mnemonic:"AND", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x3d, Mnemonic:"AND", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0x3e, Mnemonic:"ROL", Bytes:3, Cycles:7, AddressingMode:8, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x40, Mnemonic:"RTI", Bytes:1, Cycles:6, AddressingMode:0, PageSensitive:false, Effect:5},
&InstructionDefinition{ObjectCode:0x41, Mnemonic:"EOR", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x45, Mnemonic:"EOR", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x46, Mnemonic:"LSR", Bytes:2, Cycles:5, AddressingMode:4, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x48, Mnemonic:"PHA", Bytes:1, Cycles:3, AddressingMode:0, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x49, Mnemonic:"EOR", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x4a, Mnemonic:"LSR", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x4b, Mnemonic:"asr", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x4c, Mnemonic:"JMP", Bytes:3, Cycles:3, AddressingMode:3, PageSensitive:false, Effect:3},
&InstructionDefinition{ObjectCode:0x4d, Mnemonic:"EOR", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x4e, Mnemonic:"LSR", Bytes:3, Cycles:6, AddressingMode:3, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x50, Mnemonic:"BVC", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0x51, Mnemonic:"EOR", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x55, Mnemonic:"EOR", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x56, Mnemonic:"LSR", Bytes:2, Cycles:6, AddressingMode:10, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x58, Mnemonic:"CLI", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x59, Mnemonic:"EOR", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x5d, Mnemonic:"EOR", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0x5e, Mnemonic:"LSR", Bytes:3, Cycles:7, AddressingMode:8, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x60, Mnemonic:"RTS", Bytes:1, Cycles:6, AddressingMode:0, PageSensitive:false, Effect:4},
&InstructionDefinition{ObjectCode:0x61, Mnemonic:"ADC", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x65, Mnemonic:"ADC", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x66, Mnemonic:"ROR", Bytes:2, Cycles:5, AddressingMode:4, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x68, Mnemonic:"PLA", Bytes:1, Cycles:4, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x69, Mnemonic:"ADC", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x6a, Mnemonic:"ROR", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0x6c, Mnemonic:"JMP", Bytes:3, Cycles:5, AddressingMode:5, PageSensitive:false, Effect:3},
&InstructionDefinition{ObjectCode:0x6d, Mnemonic:"ADC", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x6e, Mnemonic:"ROR", Bytes:3, Cycles:6, AddressingMode:3, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x70, Mnemonic:"BVS", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0x71, Mnemonic:"ADC", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x75, Mnemonic:"ADC", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x76, Mnemonic:"ROR", Bytes:2, Cycles:6, AddressingMode:10, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x78, Mnemonic:"SEI", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x79, Mnemonic:"ADC", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0x7d, Mnemonic:"ADC", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0x7e, Mnemonic:"ROR", Bytes:3, Cycles:7, AddressingMode:8, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0x80, Mnemonic:"dop", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x81, Mnemonic:"STA", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:1},
nil,
nil,
&InstructionDefinition{ObjectCode:0x84, Mnemonic:"STY", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x85, Mnemonic:"STA", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x86, Mnemonic:"STX", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:1},
nil,
&InstructionDefinition{ObjectCode:0x88, Mnemonic:"DEY", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0x8a, Mnemonic:"TXA", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0x8c, Mnemonic:"STY", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x8d, Mnemonic:"STA", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x8e, Mnemonic:"STX", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:1},
nil,
&InstructionDefinition{ObjectCode:0x90, Mnemonic:"BCC", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0x91, Mnemonic:"STA", Bytes:2, Cycles:6, AddressingMode:7, PageSensitive:false, Effect:1},
nil,
nil,
&InstructionDefinition{ObjectCode:0x94, Mnemonic:"STY", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x95, Mnemonic:"STA", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x96, Mnemonic:"STX", Bytes:2, Cycles:4, AddressingMode:11, PageSensitive:false, Effect:1},
nil,
&InstructionDefinition{ObjectCode:0x98, Mnemonic:"TYA", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0x99, Mnemonic:"STA", Bytes:3, Cycles:5, AddressingMode:9, PageSensitive:false, Effect:1},
&InstructionDefinition{ObjectCode:0x9a, Mnemonic:"TXS", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
nil,
&InstructionDefinition{ObjectCode:0x9d, Mnemonic:"STA", Bytes:3, Cycles:5, AddressingMode:8, PageSensitive:false, Effect:1},
nil,
nil,
&InstructionDefinition{ObjectCode:0xa0, Mnemonic:"LDY", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa1, Mnemonic:"LDA", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa2, Mnemonic:"LDX", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xa4, Mnemonic:"LDY", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa5, Mnemonic:"LDA", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa6, Mnemonic:"LDX", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa7, Mnemonic:"lax", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa8, Mnemonic:"TAY", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xa9, Mnemonic:"LDA", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xaa, Mnemonic:"TAX", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xac, Mnemonic:"LDY", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xad, Mnemonic:"LDA", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xae, Mnemonic:"LDX", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xb0, Mnemonic:"BCS", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0xb1, Mnemonic:"LDA", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xb3, Mnemonic:"lax", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0xb4, Mnemonic:"LDY", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xb5, Mnemonic:"LDA", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xb6, Mnemonic:"LDX", Bytes:2, Cycles:4, AddressingMode:11, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xb7, Mnemonic:"lax", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xb8, Mnemonic:"CLV", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xb9, Mnemonic:"LDA", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0xba, Mnemonic:"TSX", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xbc, Mnemonic:"LDY", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0xbd, Mnemonic:"LDA", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0xbe, Mnemonic:"LDX", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xc0, Mnemonic:"CPY", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xc1, Mnemonic:"CMP", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:0},
nil,
nil,
&InstructionDefinition{ObjectCode:0xc4, Mnemonic:"CPY", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xc5, Mnemonic:"CMP", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xc6, Mnemonic:"DEC", Bytes:2, Cycles:5, AddressingMode:4, PageSensitive:false, Effect:2},
&InstructionDefinition{ObjectCode:0xc7, Mnemonic:"dcp", Bytes:2, Cycles:5, AddressingMode:4, PageSensitive:false, Effect:2},
&InstructionDefinition{ObjectCode:0xc8, Mnemonic:"INY", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xc9, Mnemonic:"CMP", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xca, Mnemonic:"DEX", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xcc, Mnemonic:"CPY", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xcd, Mnemonic:"CMP", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xce, Mnemonic:"DEC", Bytes:3, Cycles:6, AddressingMode:3, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0xd0, Mnemonic:"BNE", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0xd1, Mnemonic:"CMP", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0xd5, Mnemonic:"CMP", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xd6, Mnemonic:"DEC", Bytes:2, Cycles:6, AddressingMode:10, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0xd8, Mnemonic:"CLD", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xd9, Mnemonic:"CMP", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0xdd, Mnemonic:"CMP", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0xde, Mnemonic:"DEC", Bytes:3, Cycles:7, AddressingMode:8, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0xe0, Mnemonic:"CPX", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xe1, Mnemonic:"SBC", Bytes:2, Cycles:6, AddressingMode:6, PageSensitive:false, Effect:0},
nil,
nil,
&InstructionDefinition{ObjectCode:0xe4, Mnemonic:"CPX", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xe5, Mnemonic:"SBC", Bytes:2, Cycles:3, AddressingMode:4, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xe6, Mnemonic:"INC", Bytes:2, Cycles:5, AddressingMode:4, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0xe8, Mnemonic:"INX", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xe9, Mnemonic:"SBC", Bytes:2, Cycles:2, AddressingMode:1, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xea, Mnemonic:"NOP", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
nil,
&InstructionDefinition{ObjectCode:0xec, Mnemonic:"CPX", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xed, Mnemonic:"SBC", Bytes:3, Cycles:4, AddressingMode:3, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xee, Mnemonic:"INC", Bytes:3, Cycles:6, AddressingMode:3, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0xf0, Mnemonic:"BEQ", Bytes:2, Cycles:2, AddressingMode:2, PageSensitive:true, Effect:3},
&InstructionDefinition{ObjectCode:0xf1, Mnemonic:"SBC", Bytes:2, Cycles:5, AddressingMode:7, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0xf5, Mnemonic:"SBC", Bytes:2, Cycles:4, AddressingMode:10, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xf6, Mnemonic:"INC", Bytes:2, Cycles:6, AddressingMode:10, PageSensitive:false, Effect:2},
nil,
&InstructionDefinition{ObjectCode:0xf8, Mnemonic:"SED", Bytes:1, Cycles:2, AddressingMode:0, PageSensitive:false, Effect:0},
&InstructionDefinition{ObjectCode:0xf9, Mnemonic:"SBC", Bytes:3, Cycles:4, AddressingMode:9, PageSensitive:true, Effect:0},
nil,
nil,
nil,
&InstructionDefinition{ObjectCode:0xfd, Mnemonic:"SBC", Bytes:3, Cycles:4, AddressingMode:8, PageSensitive:true, Effect:0},
&InstructionDefinition{ObjectCode:0xfe, Mnemonic:"INC", Bytes:3, Cycles:7, AddressingMode:8, PageSensitive:false, Effect:2},
nil,}, nil
}