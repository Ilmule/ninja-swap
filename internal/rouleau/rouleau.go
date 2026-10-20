package rouleau

import "fmt"

type Rouleaux struct {
	Ciel  []int // A
	Terre []int // B
	Count int
}

func New(a []int) *Rouleaux {
	ciel := make([]int, len(a))
	copy(ciel, a)
	return &Rouleaux{Ciel: ciel, Terre: []int{}}
}

func swap(s []int) {
	if len(s) < 2 {
		return
	}
	s[0], s[1] = s[1], s[0]
}

func rot(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	haut := s[0]
	copy(s, s[1:])
	s[n-1] = haut
}

func rrot(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	bas := s[n-1]
	copy(s[1:], s[:n-1])
	s[0] = bas
}

func push(src, dst []int) ([]int, []int) {
	if len(src) == 0 {
		return src, dst
	}
	haut := src[0]
	return src[1:], append([]int{haut}, dst...)
}

func (R *Rouleaux) Sa()  { swap(R.Ciel); R.Count++ }
func (R *Rouleaux) Sb()  { swap(R.Terre); R.Count++ }
func (R *Rouleaux) Ss()  { swap(R.Ciel); swap(R.Terre); R.Count++ }
func (R *Rouleaux) Ra()  { rot(R.Ciel); R.Count++ }
func (R *Rouleaux) Rb()  { rot(R.Terre); R.Count++ }
func (R *Rouleaux) Rr()  { rot(R.Ciel); rot(R.Terre); R.Count++ }
func (R *Rouleaux) Rra() { rrot(R.Ciel); R.Count++ }
func (R *Rouleaux) Rrb() { rrot(R.Terre); R.Count++ }
func (R *Rouleaux) Rrr() { rrot(R.Ciel); rrot(R.Terre); R.Count++ }

func (R *Rouleaux) Pa() {
	R.Terre, R.Ciel = push(R.Terre, R.Ciel)
	R.Count++
}

func (R *Rouleaux) Pb() {
	R.Ciel, R.Terre = push(R.Ciel, R.Terre)
	R.Count++
}

func (R *Rouleaux) Appliquer(nom string) error {
	switch nom {
	case "sa":
		R.Sa()
	case "sb":
		R.Sb()
	case "ss":
		R.Ss()
	case "pa":
		R.Pa()
	case "pb":
		R.Pb()
	case "ra":
		R.Ra()
	case "rb":
		R.Rb()
	case "rr":
		R.Rr()
	case "rra":
		R.Rra()
	case "rrb":
		R.Rrb()
	case "rrr":
		R.Rrr()
	default:
		return fmt.Errorf("jutsu inconnu : %q", nom)
	}
	return nil
}

func (R *Rouleaux) EstTrie() bool {
	if len(R.Terre) != 0 {
		return false
	}
	for i := 1; i < len(R.Ciel); i++ {
		if R.Ciel[i-1] > R.Ciel[i] {
			return false
		}
	}
	return true
}
