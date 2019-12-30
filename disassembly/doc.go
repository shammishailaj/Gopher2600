// Package disassembly coordinates the disassembly Atari2600 (6507) cartridges.
//
// For quick disassemblies the FromCartridge() function can be used. Many
// debuggers will probably find it more useful however, to disassemble from the
// memory of an already instantiated VCS.
//
//	disasm, _ := disassembly.FromMemory(cartMem, symbols.NewTable())
//
// The FromMemory() function takes an instance of a symbols.Table, or nil. In
// the example above, we've used the result of NewTable(); which is fine but
// limits the potential of the disassembly package. For best results, the
// symbols.ReadSymbolsFile() function should be used (see symbols package for
// details). Note that the FromCartridge() function handles symbols files for
// you.
//
// The disassembly package performs two types of disassembly: what we call
// linear and flow disasseblies. Both are useful and eack eke out different
// information from cartridge memory. In a nutshell:
//
// Linear disassembly decodes every possible address in the cartridge. if the
// "execution" of the address succeeds the result is stored. Flow disassembly
// on the other hand decodes only those addresses that flow from the reset
// adddress as the program unfolds.
//
// In flow disassembly every branch and subroutine is considered. This is done
// by turning the CPU's "flow control" off and handling each and every the
// branch manually. Even with this method however, it is possible for a program
// to expect to be taken somewhere (when executed normally) not reachable. For
// example:
//
// - Addresses stuffed into the stack and RTS being called, without an explicit
// JSR.
//
// - Branching or jumping to non-cartridge memory. (ie. RAM) and executing code
// there.
//
// The flow disassembly collates any possible oddities it encounters and the
// Analysis() function can be used to summarise them.
//
// As already mentionied, linear disassembly looks at every possible memory
// location. The downside of this is that a lot of what is found will be
// nonsense (data segments never intended for execution, for instance). This
// make linear disassembly unsuitable for some applications. For example,
// presenting a disassembly of an entire cartridge.
//
// Where linear disassembly *is* useful is for referencing an address that you
// *know* contains a valid instruction - compare to flow disassembly where the
// address might not have been reached during the disassembly process.
//
// Note that linear cannot do anything about the posibility of executing code
// from area outside of cartridge space (ie. RAM).
//
// All that said, the flow/linear difference is invisible to the user of the
// disassembly package. Instead, the functions Get(), Write() and Grep() are
// used. These functions use the most appropriate disassembly for the use case.
//
//	Write() --> flow
//	Get()   --> linear
//	Grep()  --> flow
package disassembly