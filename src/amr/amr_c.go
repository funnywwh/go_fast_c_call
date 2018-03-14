package amr

import _ "unsafe"

//go:cgo_ldflag "-L./src/amr"
//go:cgo_ldflag "-lclib"

//go:cgo_import_static Encoder_Interface_init
//go:cgo_import_static Encoder_Interface_exit

//go:linkname Encoder_Interface_init Encoder_Interface_init
var Encoder_Interface_init byte

//go:linkname Encoder_Interface_exit Encoder_Interface_exit
var Encoder_Interface_exit byte
