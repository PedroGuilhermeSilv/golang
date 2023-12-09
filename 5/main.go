package main

import "fmt"

type ID int

var (
	meuArray [3]int
)

func main() {
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3
	fmt.Println(meuArray[2])
	for i , v := range meuArray{
		fmt.Printf("O valor do indice é %d e o valor é %d ",i,v)
	}

}
 