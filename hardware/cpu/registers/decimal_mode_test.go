package registers_test

import (
	"gopher2600/hardware/cpu/registers"
	rtest "gopher2600/hardware/cpu/registers/test"
	"gopher2600/test"
	"testing"
)

func TestDecimalMode(t *testing.T) {
	var rcarry bool

	// initialisation
	r8 := registers.NewRegister(0, "test")

	// addition without carry
	rcarry = r8.AddDecimal(1, false)
	rtest.EquateRegisters(t, r8, 0x01)
	test.Equate(t, rcarry, false)

	// addition with carry
	rcarry = r8.AddDecimal(1, true)
	rtest.EquateRegisters(t, r8, 0x03)
	test.Equate(t, rcarry, false)

	// subtraction with carry (subtract value)
	r8.Load(9)
	rtest.EquateRegisters(t, r8, 0x09)
	rcarry = r8.SubtractDecimal(1, true)
	rtest.EquateRegisters(t, r8, 0x08)

	// subtraction without carry (subtract value and another 1)
	rcarry = r8.SubtractDecimal(1, false)
	rtest.EquateRegisters(t, r8, 0x06)

	// addition on tens boundary
	r8.Load(9)
	rtest.EquateRegisters(t, r8, 0x09)
	rcarry = r8.AddDecimal(1, false)
	rtest.EquateRegisters(t, r8, 0x10)

	// subtraction on tens boundary
	rcarry = r8.SubtractDecimal(1, true)
	rtest.EquateRegisters(t, r8, 0x09)

	// addition on hundreds boundary
	r8.Load(0x99)
	rtest.EquateRegisters(t, r8, 0x99)
	rcarry = r8.AddDecimal(1, false)
	rtest.EquateRegisters(t, r8, 0x00)
	test.Equate(t, rcarry, true)

	// subtraction on hundreds boundary
	rcarry = r8.SubtractDecimal(1, true)
	rtest.EquateRegisters(t, r8, 0x99)
}
