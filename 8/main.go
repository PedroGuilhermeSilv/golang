package main

import (
	"errors"
	"fmt"
)
func main(){
	res , err :=soma(50,30)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println(res)
	}
	
	
}

func soma(a,b int) (int, error){
	soma := a+b;
	if soma>=50 {
		return soma , errors.New("maior que 50")
	}
	return soma, nil
}