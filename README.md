# Go
##  3-Fundação
- Para declaração de varáveis você pode usar a anotação := que o go vai fazer a inferência do tipo (obs: só pode ser usada uma vez na criação da variável). Exemplo (pasta 2):
```
var string a = "texto"
a := "texto"
```

- Podemos criar tipos de variáveis. Exemplo (pasta 3):
```
type ID int
var a ID
``` 

- Podemos imprimir o tipo e valor da variável usando a biblioteca ftm (pasta 4):
1. Tipo:
``` 
fmt.Printf("o tipo de f é %T", f)
``` 
2. Valor:
``` 
fmt.Printf("o valor de f é %v", f)
``` 

## Slice
- O slice é um array que permite aummentar sua capacitade de acordo com a demanda. Todo append() feito ultrapasse a capacidade inicial o Go criará um array com a mesma capacidade duplicando o tamanho original. 

## Função
- No golang podemos retornar dois resultados em uma mesma função por exemplo:
```
func main(){
	res , err :=soma(50,30)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println(res)
	}	
}
```

```
func soma(a,b int) (int, error){
	soma := a+b;
	if soma>=50 {
		return soma , errors.New("maior que 50")
	}
	return soma, nil
}
```

Essa lógica é utilizada como try catch em outras linguagens.

- As funções podem ser variáticas, ou seja, posso definir um valor inifito para a quantidade de parâmetros que ela pode receber por exemplo:

```
func main(){
	res :=soma(50,30)
	fmt.Println(res)	
}
```
```
func soma(numeros ...int) int{
	var soma int
	for _, value := range numeros{
		soma += value 
	}
	return soma
}
```

- Closures: funções anônimas ou funções que rodam dentro de outras funções são geralmente usada para executar comandos de rotinas, por exemplo:

```
func() {
	rodar webserver
}()
```

- o go não é ummma linguagem orientada a objetos então ela trabalha com algo similar chamado de "struct". A struct permite você criar algo semente a uma classe, por exemplo:

```
type Cliente struct{
	Nome string
	Idade int
	Ativo bool
}
```
