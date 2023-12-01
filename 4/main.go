package main

import "fmt"

type ID int

var (
	f ID
	a string
)

func main() {
	fmt.Printf("o tipo de f é %T", f)
	f = 10
	fmt.Printf("o valor de f é %v", f)
}
