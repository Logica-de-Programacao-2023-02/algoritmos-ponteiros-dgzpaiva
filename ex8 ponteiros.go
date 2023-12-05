package main

import "fmt"

func modificaValor(ponteiro *int) {

	*ponteiro = 20
}

func main() {

	var minhaVariavel int
	ponteiro := &minhaVariavel

	fmt.Println("Valor inicial da variável:", minhaVariavel)

	modificaValor(ponteiro)

	fmt.Println("Valor após modificação:", minhaVariavel)

}
