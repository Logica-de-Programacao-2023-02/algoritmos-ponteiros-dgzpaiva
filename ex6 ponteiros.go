package main

import (
	"fmt"
)

type Livro struct {
	Titulo string
	Autor  string
}

func alteraTituloSeAutorAnonimo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro X",
		Autor:  "Anônimo",
	}
	fmt.Println("Livro antes:", livro)
	alteraTituloSeAutorAnonimo(&livro)
	fmt.Println("Livro depois:", livro)
}
