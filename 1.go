package main

import "fmt"

func div(x, y int) int {
	if y == 0 {
		return 0
	} else {
		return x / y
	}
}

func main() {
	x := 50
	y := 10
	resul := div(x, y)
	fmt.Println(resul)
}
