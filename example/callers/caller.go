package callers

import "github.com/hexpay-io/go-dylib/dloader"

// Encapsulates the fields and methods of the original Dylib struct
type Caller struct {
	lib dloader.Dylib
}

func NewCaller(lib dloader.Dylib) Caller {
	return Caller{lib: lib}
}
