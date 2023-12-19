package main

import (
	"fmt"
)

type Cliente struct{
	Nome string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar(){
	c.Ativo =false
	fmt.Print("cliente desativado")
}

func main(){
	gui := Cliente{
		Nome: "Guilhermme",
		Idade: 22,
		Ativo: true,
	}
	gui.Desativar()
	fmt.Printf("Nome: %s, Idade %d, ativo: %t",gui.Nome,gui.Idade,gui.Ativo)
}

