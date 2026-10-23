// =====================================================================
//  internal/rouleau/rouleau_test.go — L'ENTRAÎNEMENT DES JUTSUS
//  Mission 0 (Académie)
// =====================================================================
//
//  Lancer les tests :   go test ./internal/rouleau/
//
//  À FAIRE : au moins un test par jutsu, sur trois situations.
//
//   1. Cas normal
//      ex. A = 3 1 2   →  sa  →  A = 1 3 2
//      ex. A = 3 1 2   →  ra  →  A = 1 2 3
//      ex. A = 3 1 2   →  rra →  A = 2 3 1
//
//   2. Rouleau vide : le jutsu ne fait rien et ne panique pas.
//      ex. pa avec B vide, sa avec A vide.
//
//   3. Rouleau à une seule archive : sa, ra, rra ne changent rien.
//
//  À tester aussi :
//   - les jutsus « en duo » agissent bien sur les DEUX rouleaux ;
//   - le compteur de jutsus augmente de 1 (même pour ss, rr, rrr) ;
//   - l'application par nom renvoie une erreur pour un nom inconnu ;
//   - la vérification « A trié et B vide » sur quelques cas.
//
//  Astuce : les « tests en table » (une slice de cas + une boucle) évitent
//  d'écrire 30 fois la même fonction de test.
//
//  Ces tests comptent dans la note de robustesse.
//
// =====================================================================

package rouleau

import "testing"

func egal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestJutsus(t *testing.T) {
	test := []struct {
		jutsu string
		a     []int
		b     []int
		attA  []int
		attB  []int
	}{
		{"sa", []int{3, 1, 2}, []int{9, 8}, []int{1, 3, 2}, []int{9, 8}},
		{"sa", []int{}, []int{9, 8}, []int{}, []int{9, 8}},
		{"sa", []int{7}, []int{9, 8}, []int{7}, []int{9, 8}},

		{"sb", []int{9, 8}, []int{3, 1, 2}, []int{9, 8}, []int{1, 3, 2}},
		{"sb", []int{9, 8}, []int{}, []int{9, 8}, []int{}},
		{"sb", []int{9, 8}, []int{7}, []int{9, 8}, []int{7}},

		{"ss", []int{3, 1, 2}, []int{6, 4, 5}, []int{1, 3, 2}, []int{4, 6, 5}},
		{"ss", []int{7}, []int{6, 4, 5}, []int{7}, []int{4, 6, 5}},
		{"ss", []int{}, []int{6, 4, 5}, []int{}, []int{4, 6, 5}},
		{"ss", []int{}, []int{}, []int{}, []int{}},

		{"pa", []int{1, 2}, []int{3, 4}, []int{3, 1, 2}, []int{4}},
		{"pa", []int{1, 2}, []int{}, []int{1, 2}, []int{}},
		{"pa", []int{}, []int{3}, []int{3}, []int{}},

		{"pb", []int{3, 1, 2}, []int{}, []int{1, 2}, []int{3}},
		{"pb", []int{3, 1}, []int{4}, []int{1}, []int{3, 4}},
		{"pb", []int{}, []int{4}, []int{}, []int{4}},

		{"ra", []int{3, 1, 2}, []int{9, 8}, []int{1, 2, 3}, []int{9, 8}},
		{"ra", []int{}, []int{9, 8}, []int{}, []int{9, 8}},
		{"ra", []int{7}, []int{9, 8}, []int{7}, []int{9, 8}},

		{"rb", []int{9, 8}, []int{3, 1, 2}, []int{9, 8}, []int{1, 2, 3}},
		{"rb", []int{9, 8}, []int{}, []int{9, 8}, []int{}},
		{"rb", []int{9, 8}, []int{7}, []int{9, 8}, []int{7}},

		{"rr", []int{3, 1, 2}, []int{6, 4, 5}, []int{1, 2, 3}, []int{4, 5, 6}},
		{"rr", []int{7}, []int{6, 4, 5}, []int{7}, []int{4, 5, 6}},
		{"rr", []int{}, []int{}, []int{}, []int{}},

		{"rra", []int{3, 1, 2}, []int{9, 8}, []int{2, 3, 1}, []int{9, 8}},
		{"rra", []int{}, []int{9, 8}, []int{}, []int{9, 8}},
		{"rra", []int{7}, []int{9, 8}, []int{7}, []int{9, 8}},

		{"rrb", []int{9, 8}, []int{3, 1, 2}, []int{9, 8}, []int{2, 3, 1}},
		{"rrb", []int{9, 8}, []int{}, []int{9, 8}, []int{}},
		{"rrb", []int{9, 8}, []int{7}, []int{9, 8}, []int{7}},

		{"rrr", []int{3, 1, 2}, []int{6, 4, 5}, []int{2, 3, 1}, []int{5, 6, 4}},
		{"rrr", []int{7}, []int{6, 4, 5}, []int{7}, []int{5, 6, 4}},
		{"rrr", []int{}, []int{}, []int{}, []int{}},
	}

	for _, c := range test {
		r := New(c.a)
		r.Terre = append([]int{}, c.b...)

		if err := r.Appliquer(c.jutsu); err != nil {
			t.Errorf("%s : erreur inattendue : %v", c.jutsu, err)
			continue
		}
		if !egal(r.Ciel, c.attA) || !egal(r.Terre, c.attB) {
			t.Errorf("%s sur A=%v B=%v : obtenu A=%v B=%v, attendu A=%v B=%v",
				c.jutsu, c.a, c.b, r.Ciel, r.Terre, c.attA, c.attB)
		}
	}
}

// Le compteur augmente de 1 par jutsu, même pour les jutsus en duo.
func TestCompteur(t *testing.T) {
	for _, nom := range []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"} {
		r := New([]int{3, 1, 2})
		r.Terre = []int{6, 4, 5}
		if err := r.Appliquer(nom); err != nil {
			t.Fatalf("%s : %v", nom, err)
		}
		if r.Count != 1 {
			t.Errorf("%s : Count = %d, attendu 1", nom, r.Count)
		}
	}

	r := New([]int{3, 1, 2})
	r.Ra()
	r.Pb()
	r.Ss()
	if r.Count != 3 {
		t.Errorf("après 3 jutsus : Count = %d, attendu 3", r.Count)
	}
}

func TestJutsuInconnu(t *testing.T) {
	for _, nom := range []string{"", "xx", "SA", "sa ", "rrra", "pc"} {
		r := New([]int{3, 1, 2})
		if err := r.Appliquer(nom); err == nil {
			t.Errorf("%q devrait renvoyer une erreur", nom)
		}
		if !egal(r.Ciel, []int{3, 1, 2}) || len(r.Terre) != 0 || r.Count != 0 {
			t.Errorf("%q a modifié les rouleaux : A=%v B=%v Count=%d", nom, r.Ciel, r.Terre, r.Count)
		}
	}
}

func TestEstTrie(t *testing.T) {
	cas := []struct {
		a, b []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{}, true},
		{[]int{}, []int{}, true},
		{[]int{42}, []int{}, true},
		{[]int{-5, 0, 7}, []int{}, true},
		{[]int{2, 1, 3}, []int{}, false},
		{[]int{3, 2, 1}, []int{}, false},
		{[]int{1, 2, 3}, []int{4}, false}, 
		{[]int{}, []int{1}, false},
	}
	for _, c := range cas {
		r := New(c.a)
		r.Terre = append([]int{}, c.b...)
		if got := r.EstTrie(); got != c.want {
			t.Errorf("EstTrie(A=%v, B=%v) = %v, attendu %v", c.a, c.b, got, c.want)
		}
	}
}

func TestNewCopie(t *testing.T) {
	orig := []int{3, 1, 2}
	r := New(orig)
	r.Sa()
	r.Ra()
	if !egal(orig, []int{3, 1, 2}) {
		t.Errorf("la liste d'origine a été modifiée : %v", orig)
	}
	if r.Count != 2 || len(r.Terre) != 0 {
		t.Errorf("état initial incorrect : Count=%d B=%v", r.Count, r.Terre)
	}
}

func TestExempleSujet(t *testing.T) {
	r := New([]int{2, 3, 1, 4})
	for _, nom := range []string{"pb", "sa", "pa", "sa"} {
		if err := r.Appliquer(nom); err != nil {
			t.Fatal(err)
		}
	}
	if !r.EstTrie() || r.Count != 4 {
		t.Errorf("A=%v B=%v Count=%d", r.Ciel, r.Terre, r.Count)
	}
}
