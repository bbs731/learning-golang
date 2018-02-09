//https://syslog.ravelin.com/go-and-memory-layout-6ef30c730d51
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MyData struct {
	aByte   byte
	aShort  int16
	anInt32 int32
	aSlice  []byte
}

func main() {
	typ := reflect.TypeOf(MyData{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	//fmt.Printf("Struct is %d bytes long\n", unsafe.Sizeof(unsafe.Pointer(&MyData{})))

	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}

	data := MyData{
		aByte:   0x1,
		aShort:  0x0203,
		anInt32: 0x04050607,
		aSlice: []byte{
			0x08, 0x09, 0x0a,
		},
	}

	dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	fmt.Printf("Bytes are %#v\n", dataBytes)

	dataslice := *(*reflect.SliceHeader)(unsafe.Pointer(&data.aSlice))
	fmt.Printf("Slice data is %#v\n",
		(*[3]byte)(unsafe.Pointer(dataslice.Data)))

}
