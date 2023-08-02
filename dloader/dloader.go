package dloader

// #include "dloader.h"
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

type Dylib struct {
	path        string
	handler     unsafe.Pointer
	functions   map[string]unsafe.Pointer
	functionsMu *sync.Mutex
}

func NewDylib(path string) (Dylib, error) {
	var lib Dylib
	lib.path = path

	handler, err := lib.open()
	if err != nil {
		return Dylib{}, err
	}

	lib.handler = handler
	lib.functions = make(map[string]unsafe.Pointer)
	lib.functionsMu = new(sync.Mutex)

	return lib, nil
}

func (d Dylib) open() (unsafe.Pointer, error) {
	var cErr *C.char

	cPath := bufFromStr(d.path)
	handler := C.dloader_open((*C.char)(unsafe.Pointer(&cPath[0])), &cErr)
	if handler == nil {
		err := errors.New(C.GoString(cErr))
		return nil, err
	}

	return unsafe.Pointer(handler), nil
}

func (d Dylib) Lookup(symbol string) error {
	var cErr *C.char

	cSymbol := bufFromStr(symbol)
	f := C.dloader_lookup(d.handler, (*C.char)(unsafe.Pointer(&cSymbol[0])), &cErr)
	if f == nil {
		err := errors.New(C.GoString(cErr))
		return err
	}

	d.functionsMu.Lock()
	d.functions[symbol] = f
	d.functionsMu.Unlock()

	return nil
}

func (d Dylib) Update() error {
	// Create a temporary variable with the current ptr to close the library later
	h := d.handler

	handler, err := d.open()
	if err != nil {
		return err
	}
	d.handler = handler

	for symbol := range d.functions {
		err := d.Lookup(symbol)
		if err != nil {
			return err
		}
	}

	// Close the previous library only when the new one was loaded
	C.close(h)

	return nil
}

func (d Dylib) Close() {
	C.close(d.handler)
}

func (d Dylib) GetFuncPtr(symbol string) unsafe.Pointer {
	var ptr unsafe.Pointer
	d.functionsMu.Lock()

	if ptr = d.functions[symbol]; ptr == nil {
		d.functionsMu.Unlock()
		panic("You need to \"Lookup\" function before call it")
	}

	d.functionsMu.Unlock()

	return ptr
}

func bufFromStr(str string) []byte {
	// symbol + \0
	buf := make([]byte, len(str)+1)
	copy(buf, str)

	return buf
}
