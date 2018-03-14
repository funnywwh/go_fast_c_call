export GOROOT=/home/winger/go1.9.2.src/go
export CGO_ENABLED=1
export GOARCH=arm64
export GOOS=android
export GOARM=7
export NDK_ROOT=$PWD/ndk21-arm64-linux

export CC=aarch64-linux-android-gcc
export CXX=aarch64-linux-android-g++
export AR=aarch64-linux-android-ar
export CGO_LDFLAGS=

export GO15VENDOREXPERIMENT=1

export GOPATH=$PWD
export PATH=$GOPATH/bin:$PWD:$NDK_ROOT/bin:$GOROOT/bin:$PATH


$CC -c clib/clib.c -o clib/clib.o
$AR crv ./src/amr/libclib.a clib/clib.o

go env
go build

