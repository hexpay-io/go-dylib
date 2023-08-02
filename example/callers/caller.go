package callers

import "github.com/Sssilencee/go-dylib/dloader"

// Encapsulates the fields and methods of the original Dylib struct
type Caller struct {
	lib dloader.Dylib
}

func NewCaller(lib dloader.Dylib) Caller {
	return Caller{lib: lib}
}
