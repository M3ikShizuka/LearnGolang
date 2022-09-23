package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var strNil string
	pStrNilStringHeader := (*reflect.StringHeader)(unsafe.Pointer(&strNil))

	fmt.Printf(
		"strNil: %+8v\n"+
			"pStrNilStringHeader: %+8v\n",
		strNil,
		pStrNilStringHeader,
	)

	strNil = "ABC"

	fmt.Printf(
		"strNil: %+8v\n"+
			"pStrNilStringHeader: %+8v\n"+
			"strNil[1]: %c\n",
		strNil,
		pStrNilStringHeader,
		strNil[1],
	)
}
