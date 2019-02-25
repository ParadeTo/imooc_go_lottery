package main

import "fmt"

func fib(N int) int {
	var m = make(map[int]int)
	var f func(int) int

	f = func(N int) int {
		if N == 0 || N == 1 {
			return N
		}
		if _, ok := m[N-1]; !ok {
			m[N-1] = f(N-1)
		}
		if _, ok := m[N-2]; !ok {
			m[N-2] = f(N-2)
		}
		return m[N-1] + m[N-2]
	}

	return f(N)
}

func changeM(m map[int]int) {
	m[1] = 1
}

func main() {
	var m map[int]int
	m = make(map[int]int)
	changeM(m)
	fmt.Println(m)
}
