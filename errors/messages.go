package errors

var messages = map[Errno]string{
	FatalError: "fatality: %s: %s",

	// sentinals
	UserInterrupt:   "user interrupt",                        // sentinal
	ScriptEnd:       "end of script (%s)",                    // sentinal
	PowerOff:        "emulated machine has been powered off", // sentinal
	PeriphUnplugged: "controller unplugged from '%s'",        // sentinal
	OutOfSpec:       "out of spec: %s",                       // sentinal

	// program modes
	PlayError:     "error emulating vcs: %s",
	DebuggerError: "error debugging vcs: %s",
	FPSError:      "error during fps profiling: %s",
	DisasmError:   "error during disassembly: %s",

	// debugger
	ParserError:     "parser error: %s: %s (char %d)", // first placeholder is the command definition
	ValidationError: "%s for %s",
	InvalidTarget:   "invalid target: %s",
	CommandError:    "%s",
	TerminalError:   "%s",

	// script
	ScriptFileError:       "script error: %s",
	ScriptFileUnavailable: "script error: cannot open script file (%s)",
	ScriptRunError:        "script error: use of '%s' is not allowed in scripts [%s::%d]",
	ScriptScribeError:     "script scribe error: %s",

	// recorder
	RecordingError:    "controller recording error: %s",
	PlaybackError:     "controller playback error: %s",
	PlaybackHashError: "controller playback error: hash error: %s",

	// regression
	RegressionDBError: "regression database error: %s",
	RegressionFail:    "regression fail: %s",

	// symbols
	SymbolsFileError:       "symbols error: error processing symbols file: %s",
	SymbolsFileUnavailable: "symbols error: no symbols file for %s",
	SymbolUnknown:          "symbols error: unrecognised symbol (%s)",

	// vcs
	VCSError: "error creating vcs: %s",

	// cpu
	UnimplementedInstruction:       "cpu error: unimplemented instruction (%0#x) at (%#04x)",
	InvalidOpcode:                  "cpu error: invalid opcode (%#04x)",
	InvalidResult:                  "cpu error: %s: %s",
	ProgramCounterCycled:           "cpu error: program counter cycled back to 0x0000",
	InvalidOperationMidInstruction: "cpu error: invalid operation mid-instruction (%s)",

	// memory
	MemoryError:         "memory error: %s",
	UnreadableAddress:   "memory error: memory location is not readable (%#04x)",
	UnwritableAddress:   "memory error: memory location is not writable (%#04x)",
	UnpokeableAddress:   "memory error: cannot poke address (%v)",
	UnrecognisedAddress: "memory error: address unrecognised (%v)",

	// cartridges
	CartridgeFileError:       "cartridge error: %s",
	CartridgeFileUnavailable: "cartridge error: cannot open cartridge file (%s)",
	CartridgeError:           "cartridge error: %s",
	CartridgeMissing:         "cartridge error: no cartridge attached",

	// peripherals
	PeriphHardwareUnavailable: "peripheral error: controller hardware unavailable (%s)",
	UnknownPeriphEvent:        "periperal error: %s: unsupported event (%v)",

	// tv
	UnknownTVRequest: "tv error: unsupported tv request (%v)",
	BasicTelevision:  "tv error: BasicTelevision: %s",
	ImageTV:          "tv error: ImageTV: %s",
	DigestTV:         "tv error: DigestTV: %s",

	// gui
	UnknownGUIRequest: "gui error: unsupported gui request (%v)",
	SDL:               "gui error: SDL: %s",
}
