package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num1 := 843
	str := ""
	digitsCount := int(math.Log10(float64(num1))) + 1
	for i := 0; i < digitsCount; i++ {
		c := num1 % 10
		num1 = num1 / 10
		str = str + strconv.Itoa(c)
	}
	num2, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка преобразования строки в число:", err)
		return
	}
	fmt.Println(num2)

}
