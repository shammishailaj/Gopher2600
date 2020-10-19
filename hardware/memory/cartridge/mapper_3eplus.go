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

package cartridge

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/jetsetilly/gopher2600/curated"
	"github.com/jetsetilly/gopher2600/hardware/memory/bus"
	"github.com/jetsetilly/gopher2600/hardware/memory/cartridge/mapper"
	"github.com/jetsetilly/gopher2600/hardware/memory/memorymap"
)

type m3ePlus struct {
	mappingID   string
	description string

	// 3e+ cartridge memory is segmented
	bankSize int
	banks    [][]uint8

	// 64 is the maximum number of banks possible under the 3e+ scheme
	ram [64][]uint8

	// cartridge memory is segmented in the 3e+ format. the 3e+ documentation
	// refers to these as slots. we prefer the segment terminology for
	// consistency.
	//
	// hotspots are provided by the Listen() function
	segment      [4]int
	segmentIsRAM [4]bool
}

// should work with any size cartridge that is a multiple of 1024
//  - tested with chess3E+20200519_3PQ6_SQ.bin
//	https://atariage.com/forums/topic/299157-chess/?do=findComment&comment=4541517
//
//	- specifciation:
//	https://atariage.com/forums/topic/307914-3e-and-macros-are-your-friend/?tab=comments#comment-4561287
func new3ePlus(data []byte) (mapper.CartMapper, error) {
	cart := &m3ePlus{
		mappingID:   "3E+",
		description: "", // no description
		bankSize:    1024,
	}

	// a ram bank is half the size of the available bank size. this is because
	// of how ram is accessed - half the addresses are for reading and the
	// other half are for writing
	const ramSize = 512

	if len(data)%cart.bankSize != 0 {
		return nil, curated.Errorf("%s: wrong number bytes in the cartridge file", cart.mappingID)
	}

	numBanks := len(data) / cart.bankSize
	cart.banks = make([][]uint8, numBanks)

	// partition binary
	for k := 0; k < numBanks; k++ {
		cart.banks[k] = make([]uint8, cart.bankSize)
		offset := k * cart.bankSize
		copy(cart.banks[k], data[offset:offset+cart.bankSize])
	}

	// allocate ram
	for k := 0; k < len(cart.ram); k++ {
		cart.ram[k] = make([]uint8, ramSize)
	}

	return cart, nil
}

func (cart m3ePlus) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%s segments: ", cart.mappingID))
	for i := range cart.segment {
		s.WriteString(fmt.Sprintf("%d", cart.segment[i]))
		if cart.segmentIsRAM[i] {
			s.WriteString("R ")
		} else {
			s.WriteString(" ")
		}
	}
	return s.String()
}

// ID implements the mapper.CartMapper interface.
func (cart m3ePlus) ID() string {
	return cart.mappingID
}

// Reset implements the mapper.CartMapper interface.
func (cart *m3ePlus) Reset(randomise bool) {
	for b := range cart.ram {
		for i := range cart.ram[b] {
			if randomise {
				cart.ram[b][i] = uint8(rand.Intn(0xff))
			} else {
				cart.ram[b][i] = 0
			}
		}
	}

	// The last 1K ROM ($FC00-$FFFF) segment in the 6502 address space (ie: $1C00-$1FFF)
	// is initialised to point to the FIRST 1K of the ROM image, so the reset vectors
	// must be placed at the end of the first 1K in the ROM image.
	for i := range cart.segment {
		cart.segment[i] = 0
		cart.segmentIsRAM[i] = false
	}
}

// Read implements the mapper.CartMapper interface.
func (cart *m3ePlus) Read(addr uint16, passive bool) (uint8, error) {
	var segment int

	if addr >= 0x0000 && addr <= 0x03ff {
		segment = 0
	} else if addr >= 0x0400 && addr <= 0x07ff {
		segment = 1
	} else if addr >= 0x0800 && addr <= 0x0bff {
		segment = 2
	} else if addr >= 0x0c00 && addr <= 0x0fff {
		segment = 3
	}

	var data uint8

	if cart.segmentIsRAM[segment] {
		data = cart.ram[cart.segment[segment]][addr&0x01ff]
	} else {
		bank := cart.segment[segment]
		if bank < len(cart.banks) {
			data = cart.banks[bank][addr&0x03ff]
		}
	}

	return data, nil
}

// Write implements the mapper.CartMapper interface.
func (cart *m3ePlus) Write(addr uint16, data uint8, passive bool, poke bool) error {
	if passive {
		return nil
	}

	var segment int

	if addr >= 0x0000 && addr <= 0x03ff {
		segment = 0
	} else if addr >= 0x0400 && addr <= 0x07ff {
		segment = 1
	} else if addr >= 0x0800 && addr <= 0x0bff {
		segment = 2
	} else if addr >= 0x0c00 && addr <= 0x0fff {
		segment = 3
	}

	if cart.segmentIsRAM[segment] {
		cart.ram[cart.segment[segment]][addr&0x01ff] = data
		return nil
	} else if poke {
		cart.banks[cart.segment[segment]][addr&0x03ff] = data
		return nil
	}

	return curated.Errorf(bus.AddressError, addr)
}

// NumBanks implements the mapper.CartMapper interface.
func (cart m3ePlus) NumBanks() int {
	return len(cart.banks)
}

// GetBank implements the mapper.CartMapper interface.
func (cart *m3ePlus) GetBank(addr uint16) mapper.BankInfo {
	var seg int
	if addr >= 0x0000 && addr <= 0x03ff {
		seg = 0
	} else if addr >= 0x0400 && addr <= 0x07ff {
		seg = 1
	} else if addr >= 0x0800 && addr <= 0x0bff {
		seg = 2
	} else { // remaining address is between 0x0c00 and 0x0fff
		seg = 3
	}

	if cart.segmentIsRAM[seg] {
		return mapper.BankInfo{Number: cart.segment[seg], IsRAM: true, Segment: seg}
	}
	return mapper.BankInfo{Number: cart.segment[seg], Segment: seg}
}

// Patch implements the mapper.CartMapper interface.
func (cart *m3ePlus) Patch(offset int, data uint8) error {
	if offset >= cart.bankSize*len(cart.banks) {
		return curated.Errorf("%s: patch offset too high (%v)", cart.ID(), offset)
	}

	bank := offset / cart.bankSize
	offset %= cart.bankSize
	cart.banks[bank][offset] = data
	return nil
}

// Listen implements the mapper.CartMapper interface.
func (cart *m3ePlus) Listen(addr uint16, data uint8) {
	// mapper 3e+ is a derivative of tigervision and so uses the same Listen()
	// mechanism. see the tigervision commentary for details

	// bankswitch on hotspot access
	if addr == 0x3f {
		segment := data >> 6
		bank := data & 0x3f
		cart.segment[segment] = int(bank)
		cart.segmentIsRAM[segment] = false
	} else if addr == 0x3e {
		segment := data >> 6
		bank := data & 0x3f
		cart.segment[segment] = int(bank)
		cart.segmentIsRAM[segment] = true
	}
}

// Step implements the mapper.CartMapper interface.
func (cart *m3ePlus) Step() {
}

// GetRAM implements the mapper.CartRAMBus interface.
func (cart m3ePlus) GetRAM() []mapper.CartRAM {
	r := make([]mapper.CartRAM, len(cart.ram))

	for i := range cart.ram {
		mapped := false
		origin := uint16(0x0000)

		for s := range cart.segment {
			mapped = cart.segment[s] == i && cart.segmentIsRAM[s]
			if mapped {
				switch s {
				case 0:
					origin = uint16(0x1000)
				case 1:
					origin = uint16(0x1400)
				case 2:
					origin = uint16(0x1800)
				case 3:
					origin = uint16(0x1c00)
				}
				break // for loop
			}
		}

		r[i] = mapper.CartRAM{
			Label:  fmt.Sprintf("%d", i),
			Origin: origin,
			Data:   make([]uint8, len(cart.ram[i])),
			Mapped: mapped,
		}
		copy(r[i].Data, cart.ram[i])
	}

	return r
}

// PutRAM implements the mapper.CartRAMBus interface.
func (cart *m3ePlus) PutRAM(bank int, idx int, data uint8) {
	cart.ram[bank][idx] = data
}

// IterateBank implements the mapper.CartMapper interface.
func (cart m3ePlus) CopyBanks() []mapper.BankContent {
	c := make([]mapper.BankContent, len(cart.banks))
	for b := 0; b < len(cart.banks); b++ {
		c[b] = mapper.BankContent{Number: b,
			Data: cart.banks[b],
			Origins: []uint16{
				memorymap.OriginCart,
				memorymap.OriginCart + uint16(cart.bankSize),
				memorymap.OriginCart + uint16(cart.bankSize)*2,
				memorymap.OriginCart + uint16(cart.bankSize)*3},
		}
	}
	return c
}
