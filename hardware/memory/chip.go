package memory

import (
	"fmt"
	"gopher2600/errors"
	"gopher2600/hardware/memory/addresses"
)

// ChipMemory defines the information for and operations allowed for those
// memory areas accessed by the VCS chips as well as the CPU
type ChipMemory struct {
	DebuggerBus
	ChipBus
	CPUBus
	PeriphBus

	origin uint16
	memtop uint16
	memory []uint8

	// additional mask to further reduce address space when read from the CPU
	cpuReadMask uint16

	// when the CPU writes to chip memory it is not writing to memory in the
	// way we might expect. instead we note the address that has been written
	// to, and a boolean true to indicate that a write has been performed by
	// the CPU
	lastWriteAddress uint16 // mapped from 16bit to chip address length
	writeData        uint8
	writeSignal      bool

	// lastReadRegister works slightly different that lastWriteAddress. it stores
	// the register *name* of the last memory location *read* by the CPU
	lastReadRegister string
}

// Peek is an implementation of memory.DebuggerBus
func (area ChipMemory) Peek(address uint16) (uint8, error) {
	sym := addresses.Read[address]
	if sym == "" {
		return 0, errors.New(errors.UnreadableAddress, address)
	}
	return area.memory[address-area.origin], nil
}

// Poke is an implementation of memory.DebuggerBus
func (area ChipMemory) Poke(address uint16, value uint8) error {
	return errors.New(errors.UnpokeableAddress, address)
}

// ChipRead is an implementation of memory.ChipBus
func (area *ChipMemory) ChipRead() (bool, ChipData) {
	if area.writeSignal {
		area.writeSignal = false
		return true, ChipData{Name: addresses.Write[area.lastWriteAddress], Value: area.writeData}
	}

	return false, ChipData{}
}

// ChipWrite is an implementation of memory.ChipBus
func (area *ChipMemory) ChipWrite(address uint16, data uint8) {
	area.memory[address] = data
}

// LastReadRegister is an implementation of memory.ChipBus
func (area *ChipMemory) LastReadRegister() string {
	r := area.lastReadRegister
	area.lastReadRegister = ""
	return r
}

// Read is an implementation of memory.CPUBus
func (area *ChipMemory) Read(address uint16) (uint8, error) {
	// note the name of the register that we are reading
	area.lastReadRegister = addresses.Read[address]

	sym := addresses.Read[address]
	if sym == "" {
		return 0, errors.New(errors.UnreadableAddress, address)
	}

	return area.memory[area.origin|address^area.origin], nil
}

// Write is an implementation of memory.CPUBus
func (area *ChipMemory) Write(address uint16, data uint8) error {
	// check that the last write to this memory area has been serviced
	if area.writeSignal {
		return errors.New(errors.MemoryError, fmt.Sprintf("unserviced write to chip memory (%s)", addresses.Write[area.lastWriteAddress]))
	}

	sym := addresses.Write[address]
	if sym == "" {
		return errors.New(errors.UnwritableAddress, address)
	}

	// note address of write
	area.lastWriteAddress = address
	area.writeSignal = true
	area.writeData = data

	return nil
}

// PeriphWrite implements memory.PeriphBus
func (area *ChipMemory) PeriphWrite(address uint16, data uint8, mask uint8) {
	d := area.memory[address] & (mask ^ 0xff)
	area.memory[address] = data | d
}
