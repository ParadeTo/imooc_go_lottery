package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const INT_SIZE int = int(unsafe.Sizeof(0))

//判断我们系统中的字节序类型
func systemEdian() {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	fmt.Println(bs)
	if bs[0] == 0 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}
}

func main() {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 1}
	fmt.Printf("% x\n", b) // 00 00 00 00 00 00 00 01
	fmt.Println(binary.LittleEndian.Uint64(b)) // 72057594037927936
	fmt.Println(binary.BigEndian.Uint64(b)) // 1
}
