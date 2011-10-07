// author: liuaifu
// laf163@gmail.com
// 2011-10-3

package main

import (
	"fmt"
	"buffer"
)

func main() {
	buf1:=buffer.New()
	buf1.WriteUint32(0x1122)
	fmt.Printf("%v\n", buf1.Buffer())
	fmt.Printf("%x\n", buf1.ReadUint32())
}

/*
$ 8g test.go
$ 8l test.8
$ ./8.out.exe
[34 17 0 0]
1122
*/
