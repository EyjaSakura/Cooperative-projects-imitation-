package main

import "fmt"

func main() {
	var scale int
	var pattern byte

	fmt.Print("Enter your Scale: ")
	fmt.Scanf("%d\n", &scale)

	fmt.Print("Enter your Pattern: ")
	fmt.Scanf("%c\n", &pattern)

	fmt.Println("Your pic is: ")
	draw(scale, pattern)
}

func draw(scale int, pattern byte) {
	for i := 0; i < scale; i++ {
		drawLine(scale, i, pattern)
	}
	for i := scale - 2; i >= 0; i-- {
		drawLine(scale, i, pattern)
	}
}

func drawLine(scale, i int, pattern byte) {
	drawSpace(scale, i)
	drawPattern(scale, i, pattern)
	drawSpace(scale, i)
	fmt.Print("\n")
}

func drawSpace(scale, i int) {
	for j := 0; j < scale-i-1; j++ {
		fmt.Print(" ")
	}
}

func drawPattern(scale, i int, pattern byte) {
	for j := 0; j < 2*i+1; j++ {
		fmt.Printf("%c", pattern)
	}
}
