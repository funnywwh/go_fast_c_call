package main

import (
	"fmt"
	"unsafe"

	"amr"
)

func main() {
	name := []byte("hello world")
	obj := amr.Encoder_Init(27, 28, uintptr(unsafe.Pointer(&name[0])))
	fmt.Printf("obj:%d\n", obj)
	amr.Encoder_Exit(obj)
}
