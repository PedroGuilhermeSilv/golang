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

- Podemos imprimir o tipo e valor da variável usando a biblioteca ftm::
1. Tipo:
``` 
fmt.Printf("o tipo de f é %T", f)
``` 
2. Valor:
``` 
fmt.Printf("o valor de f é %v", f)
``` 