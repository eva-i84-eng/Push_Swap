package main

import (
	"fmt"
	"push_swap/stack" // Importe le dossier stack de ton module
)

func main() {
	// Exemple d'utilisation :
	stack.A = []int{3, 2, 1}
	stack.Sa()
	fmt.Println(stack.A) // Affiche [2 3 1]
}