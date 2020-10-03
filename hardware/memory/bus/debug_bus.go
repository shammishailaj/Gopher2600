// This file is part of Gopher2600.
//
// Gopher2600 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Gopher2600 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Gopher2600.  If not, see <https://www.gnu.org/licenses/>.

// Package bus defines the memory bus concept. For an explanation see the
// memory package documentation.
package bus

import (
	"fmt"
)

// DebugBus defines the meta-operations for all memory areas. Think of these
// functions as "debugging" functions, that is operations outside of the normal
// operation of the machine.
type DebugBus interface {
	Peek(address uint16) (uint8, error)
	Poke(address uint16, value uint8) error
}

// CartRegistersBus defines the operations required for a debugger to access the
// registers in a cartridge.
//
// The mapper is allowed to panic if it is not interfaced with correctly.
//
// You should know the precise cartridge mapper for the CartRegisters to be
// usable.
//
// So what's the point of the interface if you need to know the details of the
// underlying type? Couldn't we just use a type assertion?
//
// Yes, but doing it this way helps with the lazy evaluation system used by
// debugging GUIs. The point of the lazy system is to prevent race conditions
// and the way we do that is to make copies of system variables before using it
// in the GUI. Now, because we must know the internals of the cartridge format,
// could we not just make those copies manually? Again, yes. But that would
// mean another place where the cartridge's internal knowledge needs to be
// coded (we need to use that knowledge in the GUI code but it would be nice to
// avoid it in the lazy system).
//
// The GetRegisters() allows us to conceptualise the copying process and to
// keep the details inside the cartridge implementation as much as possible.
type CartRegistersBus interface {
	// GetRegisters returns a copy of the cartridge's registers
	GetRegisters() CartRegisters

	// Update a register in the cartridge with new data.
	//
	// Depending on the complexity of the cartridge, the register argument may
	// need to be a structured string to uniquely identify a register (eg. a
	// JSON string, although that's probably going over the top). The details
	// of what is valid should be specified in the documentation of the mappers
	// that use the CartDebugBus.
	//
	// The data string will be converted to whatever type is required for the
	// register. For simple types then this will be usual Go representation,
	// (eg. true of false for boolean types) but it may be a more complex
	// representation. Again, the details of what is valid should be specified
	// in the mapper documentation.
	PutRegister(register string, data string)
}

// CartRegisters conceptualises the cartridge specific registers that are
// inaccessible through normal addressing
type CartRegisters interface {
	fmt.Stringer
}

// CartRAMbus is implemented for catridge mappers that have an addressable RAM
// area. This differs from a Static area which is not addressable by the VCS.
//
// Note that for convenience, some mappers will implement this interface but
// have no RAM for the specific cartridge. In these case GetRAM() will return
// nil.
//
// The test for whether a specific cartridge has additional RAM should include
// a interface type asserstion as well as checking GetRAM() == nil
type CartRAMbus interface {
	GetRAM() []CartRAM
	PutRAM(bank int, idx int, data uint8)
}

// CartRAM represents a single segment of RAM in the cartridge. A cartridge may
// contain more than one segment of RAM. The Label field can help distinguish
// between the different segments.
//
// The Origin field specifies the address of the lowest byte in RAM. The Data
// field is a copy of the actual bytes in the cartidge RAM. Because Cartidge is
// addressable, it is also possible to update cartridge RAM through the normal
// memory buses; although in the context of a debugger it is probably more
// convience to use PutRAM() in the CartRAMbus interface
type CartRAM struct {
	Label  string
	Origin uint16
	Data   []uint8
	Mapped bool
}

// CartStaticBus defines the operations required for a debugger to access the
// static area of a cartridge.
type CartStaticBus interface {
	// GetStatic returns a copy of the cartridge's static areas
	GetStatic() []CartStatic
	PutStatic(tag string, addr uint16, data uint8) error
}

// CartStatic conceptualises a static data area that is inaccessible through.
// Of the cartridge types that have static areas some have more than one static
// area.
//
// Unlike CartRAM, there is no indication of the origin address of the
// StaticArea. This is because the areas are not addressable in the usual way
// and so the concept of an origin address is meaningless and possibly
// misleading.
type CartStatic struct {
	Label string
	Data  []uint8
}

// CartTapeState is the current state of the tape
type CartTapeState struct {
	Counter    int
	MaxCounter int
	Time       float64
	MaxTime    float64
	Data       []float32
}

// CartTapeBus defines additional debugging functions for cartridge types that use tapes
type CartTapeBus interface {
	// Move tape loading to the specified mark. returns true if rewind was
	// effective
	Rewind() bool

	// GetTapeState retrieves a copy of the current state of the tape. returns
	// true is state is valid
	GetTapeState() (bool, CartTapeState)
}

// CartHotspotAction defines the action of a hotspot address
type CartHotspotAction int

// List of valid CartHotspotActions
const (
	HotspotBankSwitch CartHotspotAction = iota
	HotspotRegister
	HotspotFunction
)

// HotspotInfo details the name and purpose of hotspot address
type CartHotspotInfo struct {
	Symbol string
	Action CartHotspotAction
}

// CartHotspots will be implemented for cartridge mappers that want to report
// details of any special addresses. We'll call these hotspots for all types of
// special addresses, not just bank switches.
//
// The index to the returned maps, must be addresses in the cartridge address
// range. For normality, this should be in the primary cartridge mirror (ie.
// 0x1000 to 0x1fff)
type CartHotspots interface {
	ReadHotspots() map[uint16]CartHotspotInfo
	WriteHotspots() map[uint16]CartHotspotInfo
}
