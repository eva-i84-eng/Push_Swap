package main

import (
	"bufio"
	"fmt"
	"os"
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
		nombre, err := strconv.Atoi(arg) // Ascii to integer conversion
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error") // standar error detection
			os.Exit(1)
		}

		if doublon[nombre] {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
		doublon[nombre] = true

		stack.A = append(stack.A, nombre)

	}

	// lit l'entrée standard ligne par ligne
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		instruction := scanner.Text()

		// associe chaque instruction textuelle à la fonction stack correspondante
		switch instruction {
		case "sa":
			stack.Sa()
		case "sb":
			stack.Sb()
		case "ss":
			stack.Ss()
		case "pa":
			stack.Pa()
		case "pb":
			stack.Pb()
		case "ra":
			stack.Ra()
		case "rb":
			stack.Rb()
		case "rr":
			stack.Rr()
		case "rra":
			stack.Rra()
		case "rrb":
			stack.Rrb()
		case "rrr":
			stack.Rrr()
		default:
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	// Gestion d'une erreur éventuelle du scanner
	// On récupère l'erreur dans une variable
	err := scanner.Err()
	
	// On vérifie si cette variable n'est pas vide
	if err != nil {
    	fmt.Fprintln(os.Stderr, "Error")
    	os.Exit(1)
	}

	if stack.EstRangee() && len(stack.B) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

}
