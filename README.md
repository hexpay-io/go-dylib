# Dloader

## [[example.go](https://github.com/hexpay-io/go-dylib/blob/main/example/example.go)]
```golang
package main

import (
	"fmt"
	"time"

	"example/callers"

	"github.com/hexpay-io/go-dylib/dloader"
)

func main() {
	lib, err := dloader.NewDylib("../dylib_example/target/debug/libexample.dylib")
	if err != nil {
		panic(err)
	}

	err = lib.Lookup("example")
	if err != nil {
		panic(err)
	}
	err = lib.Lookup("free_c_char")
	if err != nil {
		panic(err)
	}

	caller := callers.NewCaller(lib)

	for i := 0; i < 20; i++ {
		// go_str is a copy of the original string, so we need to free the original
		c_ptr, go_str := caller.Example()
		caller.FreeCChar(c_ptr)

		fmt.Println(go_str)

		err := lib.Update()
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
	}

	lib.Close()
}

```

## [[callers_example.go](https://github.com/hexpay-io/go-dylib/blob/main/example/callers/callers.go)]
```golang
/*
Dylib callers
TODO: Automatically generate this code based on the Rust AST tree
*/
package callers

/*
#include <stdint.h>

typedef char * (*example_closure)();
char * example_caller(example_closure f) {
	return f();
}

typedef void (*free_c_char_closure)(char *ptr);
void free_c_char_caller(free_c_char_closure f, char *ptr) {
	f(ptr);
}
*/
import "C"
import "unsafe"

func (d Caller) Example() (unsafe.Pointer, string) {
	c_ptr := C.example_caller(C.example_closure(d.lib.GetFuncPtr("example")))
	return unsafe.Pointer(c_ptr), C.GoString(c_ptr)
}

func (d Caller) FreeCChar(ptr unsafe.Pointer) {
	C.free_c_char_caller(C.free_c_char_closure(d.lib.GetFuncPtr("free_c_char")), (*C.char)(ptr))
}
```

## [[lib.rs](https://github.com/hexpay-io/go-dylib/blob/main/dylib_example/src/lib.rs)]
```rust
use std::ffi::{c_char, CString};

#[no_mangle]
pub extern "C" fn example() -> *mut c_char {
    CString::new("Hello from Rust!")
        .unwrap()
        .into_raw()
}

#[no_mangle]
pub unsafe extern "C" fn free_c_char(ptr: *mut c_char) {
    if ptr.is_null() {
        return;
    }
    let _ = CString::from_raw(ptr);
}
```