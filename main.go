package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	fmt.Println(unsafe.Sizeof(a))
}
