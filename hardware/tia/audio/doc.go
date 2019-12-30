// Package audio implements the audio generation of the TIA. The implementation
// is taken almost directly from Ron Fries' original implementation, found in
// TIASound.c (easily searchable). The bit patterns are taken from there and
// the channels are mixed in the same way.
//
// Unlike the Fries' implementation, the Mix() function is called every video
// cycle, returning a new sample every 114th video clock. The TIA_Process()
// function in Frie's implementation meanwhile is called to fill a buffer. The
// samepl buffer in this emulation must sit outside of the TIA emulation and
// somwhere inside the television implementation.
//
// TIASound.c is published under the GNU Library GPL v2.0
//
// Some modifications were made to Fries' alogorithm in accordance to similar
// modifications made to the TIASnd.cxx file of the Stella emulator v5.1.3.
// Stella is published under the GNU GPL v2.0
package audio