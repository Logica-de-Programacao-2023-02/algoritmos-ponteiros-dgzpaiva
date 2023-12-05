package main

import (
    "fmt"
)

func inverteString(ponteiro *string) {
    texto := *ponteiro
    runas := []rune(texto)
    for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
        runas[i], runas[j] = runas[j], runas[i]
    }
    *ponteiro = string(runas)
}

func main() {
    texto := "Olá, mundo!"
    fmt.Println("Texto original:", texto)
    inverteString(&texto)
    fmt.Println("Texto invertido:", texto)
}
