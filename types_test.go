package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var x = make([]int, 3, 4)
	fmt.Println("Init slice")
	x[0] = 0
	x[1] = 1
	x[2] = 2
	PrintSliceHeader(&x)

	fmt.Println("Append")
	x = append(x, 3)
	PrintSliceHeader(&x)

	fmt.Println("Append")
	x = append(x, 4)
	PrintSliceHeader(&x)

	fmt.Println("Append")
	x = append(x, 5)
	PrintSliceHeader(&x)

	fmt.Println("Reslice")
	x2 := x[2:4]
	PrintSliceHeader(&x)
	PrintSliceHeader(&x2)
}

func PrintSliceHeader(x *[]int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(x))
	fmt.Printf("\tdata=%v len=%v cap=%v val=%v\n", hdr.Data, hdr.Len, hdr.Cap, *x)
}
