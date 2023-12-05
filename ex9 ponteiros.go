package main

import (
	"fmt"
)

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	desconto := 0.10 * livro.Preco
	livro.Preco -= desconto
}

func main() {
	livro := Livro{
		Titulo: "Livro A",
		Autor:  "Autor X",
		Preco:  50.0,
	}

	fmt.Println("Preço antes do desconto:", livro.Preco)
	aplicarDesconto(&livro)
	fmt.Println("Preço após desconto de 10%:", livro.Preco)
}
