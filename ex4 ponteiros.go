package main

import "fmt"

func somaUltimosDigitos(ponteiro *int) {
    valor := *ponteiro
    ultimosDigitos := valor % 100 
    digito1 := ultimosDigitos / 10
    digito2 := ultimosDigitos % 10
    soma := digito1 + digito2
    *ponteiro = soma
}

func main() {
    numero := 1234
    fmt.Println("Número inicial:", numero)
    somaUltimosDigitos(&numero)
    fmt.Println("Soma dos dois últimos dígitos:", numero)
}
