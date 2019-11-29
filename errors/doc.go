// Package errors is a helper package for the error type. It defines the
// AtariError type, a implementation of error the error interface, that allows
// code to wrap errors around other errors and to allow normalised formatted
// output of error messages.
//
// It also contains ID values for different classes of error and the English
// messages for those error classes. These could be translated to other
// languages, althought the i8n mechanism is not currently in place.
//
// The most useful feature is deduplication of wrapped errors. This means that
// code does not need to worry about the immediate context of the function
// which creates the error. For instance:
//
//	func main() { err := A() if err != nil { fmt.Println(err) } }
//
//	func A() error { err := B() if err != nil { return
//	errors.New(errors.DebuggerError, err) } return nil }
//
//	func B() error { err := C() if err != nil { return
//	errors.New(errors.DebuggerError, rr) } return nil }
//
//	func C() error { return errors.New(errors.PanicError, "C()", "not yet
//	implemented") }
//
// If we follow the code from main() we can see that first error created is a
// PanicError, wrapped in a DebuggerError, wrapped in a DebuggerError. The
// message for the returned error to main() will be:
//
//	error debugging vcs: panic: C(): not yet implemented
//
// and not
//
//	error debugging vcs: error debugging vcs: panic: C(): not yet implemented
//
// As can be seen,  the error message has been deduplicatd, or normalised.
//
// The PanicError, used in the above example, is a special error that should be
// used when something has happened such that the state of the emulation (or
// the tool) can no longer be guaranteed.
//
// Actual panics should only be used when the error is so terrible that there i
// nothing sensible to be done; useful for brute-enforcement of programming
// constraints and in init() functions.
package errors
