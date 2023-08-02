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
