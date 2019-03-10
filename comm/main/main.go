package main

import (
	"fmt"
	"imooc_go_lottery/comm"
)

func main() {
	var nihongo = "A日本語"
	var r = []rune(nihongo)
	fmt.Println(len(r))
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	fmt.Println(comm.Ip4toInt("1.1.1.1"))
}
