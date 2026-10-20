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


func (r *Rouleaux) Sa()  { swap(r.Ciel); r.Count++ }
func (r *Rouleaux) Sb()  { swap(r.Terre); r.Count++ }
func (r *Rouleaux) Ss()  { swap(r.Ciel); swap(r.Terre); r.Count++ }
func (r *Rouleaux) Ra()  { rot(r.Ciel); r.Count++ }
func (r *Rouleaux) Rb()  { rot(r.Terre); r.Count++ }
func (r *Rouleaux) Rr()  { rot(r.Ciel); rot(r.Terre); r.Count++ }
func (r *Rouleaux) Rra() { rrot(r.Ciel); r.Count++ }
func (r *Rouleaux) Rrb() { rrot(r.Terre); r.Count++ }
func (r *Rouleaux) Rrr() { rrot(r.Ciel); rrot(r.Terre); r.Count++ }

func (r *Rouleaux) Pa() {
	r.Terre, r.Ciel = push(r.Terre, r.Ciel)
	r.Count++
}

func (r *Rouleaux) Pb() {
	r.Ciel, r.Terre = push(r.Ciel, r.Terre)
	r.Count++
}

func (r *Rouleaux) Appliquer(nom string) error {
	switch nom {
	case "sa":
		r.Sa()
	case "sb":
		r.Sb()
	case "ss":
		r.Ss()
	case "pa":
		r.Pa()
	case "pb":
		r.Pb()
	case "ra":
		r.Ra()
	case "rb":
		r.Rb()
	case "rr":
		r.Rr()
	case "rra":
		r.Rra()
	case "rrb":
		r.Rrb()
	case "rrr":
		r.Rrr()
	default:
		return fmt.Errorf("jutsu inconnu : %q", nom)
	}
	return nil
}

func (r *Rouleaux) EstTrie() bool {
	if len(r.Terre) != 0 {
		return false
	}
	for i := 1; i < len(r.Ciel); i++ {
		if r.Ciel[i-1] > r.Ciel[i] {
			return false
		}
	}
	return true
}
