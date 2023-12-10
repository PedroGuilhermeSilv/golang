package main

import (
	"fmt"
)
func main(){
	result := func() int{
		return soma(10,90) *2
	}()
	fmt.Println(result)	
}

func soma(numeros ...int) int{
	var soma int
	for _, value := range numeros{
		soma += value 
	}
	return soma
}