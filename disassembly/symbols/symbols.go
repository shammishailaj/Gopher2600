package symbols

import (
	"fmt"
	"gopher2600/errors"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

// Table is the symbols table for the loaded programme
type Table struct {
	Locations    map[uint16]string
	ReadSymbols  map[uint16]string
	WriteSymbols map[uint16]string

	MaxLocationWidth int
	MaxSymbolWidth   int

	Valid bool
}

// StandardSymbols initialises a symbols table using the standard VCS symbols
func StandardSymbols() (*Table, error) {
	table := new(Table)
	table.ReadSymbols = vcsReadSymbols
	table.WriteSymbols = vcsWriteSymbols
	table.MaxSymbolWidth = vcsMaxSymbolWidth
	table.Valid = true
	return table, nil
}

// ReadSymbolsFile initialises a symbols table from the symbols file for the
// specified cartridge
func ReadSymbolsFile(cartridgeFilename string) (*Table, error) {
	table := new(Table)
	table.Locations = make(map[uint16]string)
	table.ReadSymbols = make(map[uint16]string)
	table.WriteSymbols = make(map[uint16]string)

	// try to open symbols file
	symFilename := cartridgeFilename
	ext := path.Ext(symFilename)
	if ext == ".BIN" {
		symFilename = fmt.Sprintf("%s.SYM", symFilename[:len(symFilename)-len(ext)])
	} else {
		symFilename = fmt.Sprintf("%s.sym", symFilename[:len(symFilename)-len(ext)])
	}

	sf, err := os.Open(symFilename)
	if err != nil {
		// if this is the empty cartridge then this error is expected. return
		// the empty symbol table
		if cartridgeFilename == "" {
			return table, nil
		}
		return table, &errors.GopherError{errors.NoSymbolsFile, errors.Values{cartridgeFilename}}
	}
	defer func() {
		_ = sf.Close()
	}()

	// get file info
	sfi, err := sf.Stat()
	if err != nil {
		return table, &errors.GopherError{errors.SymbolsFileError, errors.Values{err}}
	}

	// read symbols file and split into lines
	sym := make([]byte, sfi.Size())
	n, err := sf.Read(sym)
	if err != nil {
		return table, &errors.GopherError{errors.SymbolsFileError, errors.Values{err}}
	}
	if n != len(sym) {
		return table, &errors.GopherError{errors.SymbolsFileError, errors.Values{"file truncated"}}
	}
	lines := strings.Split(string(sym), "\n")

	// create new symbols table
	// loop over lines
	for _, ln := range lines {
		// ignore uninteresting lines
		p := strings.Fields(ln)
		if len(p) < 2 || p[0] == "---" {
			continue // for loop
		}

		// get address
		address, err := strconv.ParseUint(p[1], 16, 16)
		if err != nil {
			continue // for loop
		}

		// get symbol
		symbol := p[0]

		if unicode.IsDigit(rune(symbol[0])) {
			// if symbol begins with a number and a period then it is a location symbol
			i := strings.Index(symbol, ".")
			if i == -1 {
				continue // for loop
			}
			table.Locations[uint16(address)] = symbol[i:]
		} else {
			// put symbol in table
			table.ReadSymbols[uint16(address)] = symbol
			table.WriteSymbols[uint16(address)] = symbol
		}
	}

	// prioritise symbols with reference symbols for the selected system
	for k, v := range vcsReadSymbols {
		table.ReadSymbols[k] = v
	}
	for k, v := range vcsWriteSymbols {
		table.WriteSymbols[k] = v
	}

	// get max width of symbol in each list -- it may seem that we could keep
	// track of these width values as we go along but we can't really because
	// the overwriting of previous symbols, during the loops over
	// vcsRead/WriteSymbols above, causes havoc
	for _, s := range table.Locations {
		if len(s) > table.MaxLocationWidth {
			table.MaxLocationWidth = len(s)
		}
	}
	for _, s := range table.ReadSymbols {
		if len(s) > table.MaxSymbolWidth {
			table.MaxSymbolWidth = len(s)
		}
	}
	for _, s := range table.WriteSymbols {
		if len(s) > table.MaxSymbolWidth {
			table.MaxSymbolWidth = len(s)
		}
	}

	// indicate that symbol table should be used
	table.Valid = true

	return table, nil
}
