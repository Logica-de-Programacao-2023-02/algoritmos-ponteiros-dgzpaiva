package main

import "fmt"

func somaNPrimeirosNumeros(ponteiro *int, n int) {
    soma := 0
    for i := 1; i <= n; i++ {
        soma += i
    }
    *ponteiro = soma
}

func main() {
    numero := 0
    n := 5
    somaNPrimeirosNumeros(&numero, n)
    fmt.Println("A soma dos", n, "primeiros números naturais é:", numero)
}
