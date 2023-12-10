package main

import (
	"fmt"
)

type Cliente struct{
	Nome string
	Idade int
	Ativo bool
}

func main(){
	gui := Cliente{
		Nome: "Guilhermme",
		Idade: 22,
		Ativo: false,
	}
	fmt.Printf("Nome: %s, Idade %d, ativo: %t",gui.Nome,gui.Idade,gui.Ativo)
}

