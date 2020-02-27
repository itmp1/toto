package main

/*
#cgo linux CFLAGS: -fplugin=./calc.so
#cgo freebsd CFLAGS: -fplugin=./calc.so
#cgo darwin CFLAGS: -fplugin=./calc_darwin.so
#cgo windows CFLAGS: -fplugin=./calc_win.so
*/
import "C"
import "fmt"

func main() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
}
