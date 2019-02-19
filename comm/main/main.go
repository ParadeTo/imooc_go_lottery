package main

import (
	"fmt"
)

func main() {
	var nihongo = "A日本語"
	var r = []rune(nihongo)
	fmt.Println(len(r))
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
