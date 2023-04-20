package main

import (
	"fmt"
)

func main() {
	a := 6
	b := 8
	c := 10
	squaredA := a * a
	squaredB := b * b
	squaredC := c * c
	if a > b && a > c {
		if squaredA == squaredB+squaredC {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Не прямоугольный")
		}

	} else if b > a && b > c {
		if squaredB == squaredA+squaredC {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Не прямоугольный")
		}
	} else {
		if squaredC == squaredB+squaredA {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Не прямоугольный")
		}
	}

}
