package main

import "fmt"

func carac(text string) (int, error) {
	n := len(text)
	if n == 0 {
		return 0, fmt.Errorf("Vazio")
	}
	return n, nil
}

func main() {
	text := ""
	fmt.Println(carac(text))
}
