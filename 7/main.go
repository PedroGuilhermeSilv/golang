package main

import "fmt"
func main(){

	salarios := map[string]int{"pedro":1000, "guilherme":12000,"lucas":2000}

	fmt.Println(salarios["pedro"])

	for _, salario := range salarios{
		fmt.Printf("salario é %d \n",salario)
	}
}
