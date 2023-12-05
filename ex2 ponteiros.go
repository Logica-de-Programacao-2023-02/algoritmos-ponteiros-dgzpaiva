package main

import "fmt"

func verificaParImpar(ponteiro *int) {
    if *ponteiro%2 == 0 {
        *ponteiro = 0 // Atualiza para 0 se for par
    } else {
        *ponteiro = 1 // Atualiza para 1 se for ímpar
    }
}

func main() {
    numero := 7
    fmt.Println("Número inicial:", numero)
    verificaParImpar(&numero)
    fmt.Println("Número após verificação:", numero)
}
