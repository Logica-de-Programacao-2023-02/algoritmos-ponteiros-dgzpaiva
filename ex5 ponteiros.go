package main

import (
	"fmt"
	"math"
)

func atualizaMediaComPi(ponteiro *float64) {
	valorAtual := *ponteiro
	media := (valorAtual + math.Pi) / 2
	*ponteiro = media
}

func main() {
	numero := 10.0
	fmt.Println("Número inicial:", numero)
	atualizaMediaComPi(&numero)
	fmt.Println("Média entre o número e Pi:", numero)
}
