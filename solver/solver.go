package solver

import (
	"push_swap/stack"
)

func TrierPile() []string {
	// Si A est deja juste
	if len(stack.A) <= 1 || stack.EstRangee() {
		return nil
	}

	switch len(stack.A) {
	case 2:
		return trier2()
	case 3:
		return trier3()
	case 4, 5:
		return trierPetitePile()
	default:
		return trierGrandePile() // Pour 100 ou 500 nombres
	}
}

func trier2() []string {
	var givenArgs []string
	if stack.A[0] > stack.A[1] {
		stack.Sa()
		givenArgs = append(givenArgs, "sa")
	}
	return givenArgs
}

func trier3() []string {
	var givenArgs []string

	if len(stack.A) != 3 {
		return givenArgs
	}

	// Trouver index du + grand élément ds stack.A
	maxIdx := 0
	if stack.A[1] > stack.A[0] {
		maxIdx = 1
	}
	if stack.A[2] > stack.A[maxIdx] {
		maxIdx = 2
	}

	// Appliquer rotations en enregistrant les coups
	if maxIdx == 0 {
		stack.Ra()
		givenArgs = append(givenArgs, "ra")
	} else if maxIdx == 1 {
		stack.Rra()
		givenArgs = append(givenArgs, "rra")
	}

	// Si deux premiers ne sont pas dans le bon ordre : on swap
	if stack.A[0] > stack.A[1] {
		stack.Sa()
		givenArgs = append(givenArgs, "sa")
	}

	return givenArgs
}

func trierPetitePile() []string {
	var givenArgs []string

	if stack.EstRangee() {
		return givenArgs
	}
	// Tant qu'il reste + de 3 éléments ds A : on isole les + petits ds B
	for len(stack.A) > 3 {
		for stack.PositionDuMin() != 0 {
			if stack.PositionDuMin() <= len(stack.A)/2 {
				stack.Ra()
				givenArgs = append(givenArgs, "ra")
			} else {
				stack.Rra()
				givenArgs = append(givenArgs, "rra")
			}
		}
		stack.Pb()
		givenArgs = append(givenArgs, "pb")
	}

	// On trie les 3 éléments restants et récupère instructions générées
	givenArgs = append(givenArgs, trier3()...)

	// On remet les éléments de B dans A
	for len(stack.B) > 0 {
		stack.Pa()
		givenArgs = append(givenArgs, "pa")
	}

	return givenArgs
}

func trierGrandePile() []string {
	var givenArgs []string
	// Logique pour 100 / 500 nombres
	return givenArgs
}
