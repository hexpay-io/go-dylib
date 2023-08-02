package main

import (
	"fmt"
	"time"

	"example/callers"

	"github.com/Sssilencee/go-dylib/dloader"
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
