package main

import (
	"fmt"
)

func main() {
	k := 13257
	if k > 86399 || k < 0 {
		fmt.Println("Число должно быть в диопазоне от 0 до 86399")
		return
	}
	m := k / 60
	h := m / 60
	m = m - (m - m%60)
	fmt.Println("It is", h, "hours", m, "minutes.")
}
