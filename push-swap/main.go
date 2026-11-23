package main

import (
	"fmt"
	"os"
	"push_swap/solver"
	"push_swap/stack"
	"strconv"
	"strings"
)

func main() {
	UserArgs := os.Args[1:]

	// Vérifier s'il y a au moins un argument avant d'y accéder (évite un panic)
	if len(UserArgs) == 0 {
		return
	}
	var args []string
	doublon := make(map[int]bool)

	if len(UserArgs) == 1 { // S'il n'y a qu'un seul argument (ex: "3 2 1"), on le découpe
		args = strings.Fields(UserArgs[0])
	} else { // S'il y a plusieurs arguments (ex: 3, 2, 1), on les récupère directement
		args = UserArgs
	}

	for _, arg := range args {
		nombre, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}

		if doublon[nombre] {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
		doublon[nombre] = true

		stack.A = append(stack.A, nombre)

	}

	instructions := solver.TrierPile()

	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
