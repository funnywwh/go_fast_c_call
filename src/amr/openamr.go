package amr

import (
	"fmt"
	"unsafe"
)

//go:linkname Call runtime.asmcgocall
func Call(fn, arg unsafe.Pointer) int32

//go:cgo_unsafe_args
func Encoder_Init(mode, b int, fname uintptr) (r1 unsafe.Pointer) {
	fmt.Println("Encoder_Init")
	arg := struct {
		a     int32
		b     int32
		fname uintptr
		c     int64
		d     int
	}{int32(mode), int32(b), fname, 0, 0}
	a := Call(unsafe.Pointer(&Encoder_Interface_init), unsafe.Pointer(&arg))
	fmt.Printf("Encoder_Init a:%v\n", a)
	r1 = unsafe.Pointer(uintptr(a))
	return
}

//go:cgo_unsafe_args
func Encoder_Exit(obj unsafe.Pointer) {
	fmt.Printf("Encoder_Exit obj:%d\n", obj)

	Call(unsafe.Pointer(&Encoder_Interface_exit), obj)
}
