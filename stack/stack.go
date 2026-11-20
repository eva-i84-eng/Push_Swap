package stack

// A et B : les deux tas de cartes, du sommet (case 0) jusqu'au fond.
var A, B []int

// Gestes : les gestes joués, dans l'ordre.
var Gestes []string

// PositionDuMin renvoie la position de la plus petite carte de A (0 = sommet).
func PositionDuMin() int {
	pos := 0
	for i := 1; i < len(A); i++ {
		if A[i] < A[pos] {
			pos = i
		}
	}
	return pos
}

// EstRangee dit si A est rangée : chaque carte plus petite que celle du dessous.
func EstRangee() bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

// Sa échange les deux cartes du sommet de A.
func Sa() {
	if len(A) >= 2 {
		A[0], A[1] = A[1], A[0]
	}
	noter("sa")
}

// Pb prend la carte du sommet de A et la pose au sommet de B.
func Pb() {
	if len(A) > 0 {
		haut := A[0]
		A = A[1:]
		B = append([]int{haut}, B...)
	}
	noter("pb")
}

// Pa prend la carte du sommet de B et la pose au sommet de A.
func Pa() {
	if len(B) > 0 {
		haut := B[0]
		B = B[1:]
		A = append([]int{haut}, A...)
	}
	noter("pa")
}

// Ra fait descendre la carte du sommet de A tout au fond.
func Ra() {
	if len(A) >= 2 {
		A = append(A[1:], A[0])
	}
	noter("ra")
}

// Rra fait remonter la carte du fond de A jusqu'au sommet.
func Rra() {
	if len(A) >= 2 {
		fond := A[len(A)-1]
		A = append([]int{fond}, A[:len(A)-1]...)
	}
	noter("rra")
}

// Sb Échange les deux premiers éléments du sommet de la pile B
func Sb() {
	if len(B) >= 2 {
		B[0], B[1] = B[1], B[0]
	}
	noter("sb")
}

// Ss échange les deux premiers éléments de A et les deux premiers éléments de B en même temps.
func Ss() {
	if len(A) >= 2 {
		A[0], A[1] = A[1], A[0]
	}
	if len(B) >= 2 {
		B[0], B[1] = B[1], B[0]
	}
	noter("ss")
}

// Rb fait descendre la carte du sommet de B tout au fond.
func Rb() {
	if len(B) >= 2 {
		B = append(B[1:], B[0])
	}
	noter("rb")
}

// Rr fait une rotation vers le haut de A et de B en même temps.
func Rr() {
	if len(A) >= 2 {
		A = append(A[1:], A[0])
	}
	if len(B) >= 2 {
		B = append(B[1:], B[0])
	}
	noter("rr")
}

// Rrb fait remonter la carte du fond de B jusqu'au sommet.
func Rrb() {
	if len(B) >= 2 {
		fond := B[len(B)-1]
		B = append([]int{fond}, B[:len(B)-1]...)
	}
	noter("rrb")
}

// Rrr fait une rotation inversée de A et de B en même temps.
func Rrr() {
	if len(A) >= 2 {
		fondA := A[len(A)-1]
		A = append([]int{fondA}, A[:len(A)-1]...)
	}
	if len(B) >= 2 {
		fondB := B[len(B)-1]
		B = append([]int{fondB}, B[:len(B)-1]...)
	}
	noter("rrr")
}

// noter ajoute un geste à la liste. (Restée en minuscule car utilisée uniquement à l'intérieur de ce package).
func noter(geste string) {
	if len(Gestes) >= 100000 {
		panic("plus de 100 000 gestes, une boucle ne s'arrête sans doute jamais")
	}
	Gestes = append(Gestes, geste)
}
