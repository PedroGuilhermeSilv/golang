package main

import (
	"fmt"
)
func main(){
	res :=soma(50,30)
	fmt.Println(res)	
}

func soma(numeros ...int) int{
	var soma int
	for _, value := range numeros{
		soma += value 
	}
	return soma
}