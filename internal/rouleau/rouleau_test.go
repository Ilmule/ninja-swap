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

package rouleau_test

import (
	"math/rand"
	"testing"

	"ninja-swap/internal/rouleau"
	"ninja-swap/internal/strategie"
)

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
		r := rouleau.New(c.a)
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
		r := rouleau.New([]int{3, 1, 2})
		r.Terre = []int{6, 4, 5}
		if err := r.Appliquer(nom); err != nil {
			t.Fatalf("%s : %v", nom, err)
		}
		if r.Count != 1 {
			t.Errorf("%s : Count = %d, attendu 1", nom, r.Count)
		}
	}

	r := rouleau.New([]int{3, 1, 2})
	r.Ra()
	r.Pb()
	r.Ss()
	if r.Count != 3 {
		t.Errorf("après 3 jutsus : Count = %d, attendu 3", r.Count)
	}
}

func TestJutsuInconnu(t *testing.T) {
	for _, nom := range []string{"", "xx", "SA", "sa ", "rrra", "pc"} {
		r := rouleau.New([]int{3, 1, 2})
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
		r := rouleau.New(c.a)
		r.Terre = append([]int{}, c.b...)
		if got := r.EstTrie(); got != c.want {
			t.Errorf("EstTrie(A=%v, B=%v) = %v, attendu %v", c.a, c.b, got, c.want)
		}
	}
}

func TestNewCopie(t *testing.T) {
	orig := []int{3, 1, 2}
	r := rouleau.New(orig)
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
	r := rouleau.New([]int{2, 3, 1, 4})
	for _, nom := range []string{"pb", "sa", "pa", "sa"} {
		if err := r.Appliquer(nom); err != nil {
			t.Fatal(err)
		}
	}
	if !r.EstTrie() || r.Count != 4 {
		t.Errorf("A=%v B=%v Count=%d", r.Ciel, r.Terre, r.Count)
	}
}

// =====================================================================
//  STRATÉGIES (Patrouille, PetitesPiles, BruteForce, Shikamaru, Kage)
// =====================================================================
//
//  Ces tests sont ici (package rouleau_test) car ce package externe peut
//  importer à la fois rouleau et strategie sans cycle d'import.
//  Seules les fonctions exportées de strategie sont testables d'ici.
//
// =====================================================================

// permutations renvoie toutes les permutations de 0..n-1.
func permutations(n int) [][]int {
	var res [][]int
	var rec func(cur []int, reste []int)
	rec = func(cur []int, reste []int) {
		if len(reste) == 0 {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := range reste {
			suivant := append([]int{}, reste[:i]...)
			suivant = append(suivant, reste[i+1:]...)
			rec(append(cur, reste[i]), suivant)
		}
	}
	base := make([]int, n)
	for i := range base {
		base[i] = i
	}
	rec([]int{}, base)
	return res
}

// verifierTri rejoue la suite sur une copie de nombres et vérifie « A trié, B vide ».
func verifierTri(t *testing.T, nom string, nombres []int, suite []string) {
	t.Helper()
	r := rouleau.New(nombres)
	for _, j := range suite {
		if err := r.Appliquer(j); err != nil {
			t.Fatalf("%s sur %v : jutsu invalide %q : %v", nom, nombres, j, err)
		}
	}
	if !r.EstTrie() {
		t.Errorf("%s sur %v : pas trié (A=%v B=%v) avec la suite %v", nom, nombres, r.Ciel, r.Terre, suite)
	}
}

func TestPatrouille(t *testing.T) {
	for n := 0; n <= 6; n++ {
		for _, p := range permutations(n) {
			r := rouleau.New(p)
			suite := strategie.Patrouille(r)
			verifierTri(t, "Patrouille", p, suite)
			if !r.EstTrie() {
				t.Errorf("Patrouille(%v) : le rouleau reçu n'est pas trié (A=%v B=%v)", p, r.Ciel, r.Terre)
			}
		}
	}
	if s := strategie.Patrouille(rouleau.New([]int{1, 2, 3, 4})); len(s) != 0 {
		t.Errorf("Patrouille sur une pile triée : suite = %v, attendu vide", s)
	}
	p := []int{5, -2, 1000, 0, -77, 8}
	verifierTri(t, "Patrouille", p, strategie.Patrouille(rouleau.New(p)))

	hasard := rand.New(rand.NewSource(7))
	for _, n := range []int{10, 50, 100} {
		p := hasard.Perm(n)
		verifierTri(t, "Patrouille", p, strategie.Patrouille(rouleau.New(p)))
	}
}

func TestBruteForce(t *testing.T) {
	maxi := map[int]int{2: 1, 3: 2, 4: 5, 5: 8}
	for n := 2; n <= 5; n++ {
		for _, p := range permutations(n) {
			r := rouleau.New(p)
			suite := strategie.BruteForce(r)
			verifierTri(t, "BruteForce", p, suite)
			if len(suite) > maxi[n] {
				t.Errorf("BruteForce(%v) : %d jutsus, maximum attendu %d", p, len(suite), maxi[n])
			}
			if !r.EstTrie() {
				t.Errorf("BruteForce(%v) : le rouleau reçu n'est pas trié", p)
			}
		}
	}

	p := []int{300, -10, 42}
	verifierTri(t, "BruteForce", p, strategie.BruteForce(rouleau.New(p)))

	b := rouleau.New([]int{2, 1})
	b.Terre = []int{3}
	rien := map[string]*rouleau.Rouleaux{
		"déjà trié":   rouleau.New([]int{1, 2, 3}),
		"vide":        rouleau.New([]int{}),
		"une archive": rouleau.New([]int{4}),
		"trop grand":  rouleau.New([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}),
		"B non vide":  b,
	}
	for nom, r := range rien {
		if s := strategie.BruteForce(r); len(s) != 0 {
			t.Errorf("BruteForce (%s) : suite = %v, attendu vide", nom, s)
		}
	}
}

func TestBruteForceSixEtSept(t *testing.T) {
	if testing.Short() {
		t.Skip("test plus long, sauté avec -short")
	}
	hasard := rand.New(rand.NewSource(1))
	for _, n := range []int{6, 7} {
		p := hasard.Perm(n)
		verifierTri(t, "BruteForce", p, strategie.BruteForce(rouleau.New(p)))
	}
}

func TestPetitesPiles(t *testing.T) {
	for n := 2; n <= 5; n++ {
		for _, p := range permutations(n) {
			suite := strategie.PetitesPiles(rouleau.New(p))
			verifierTri(t, "PetitesPiles", p, suite)
			if n == 3 && len(suite) > 2 {
				t.Errorf("PetitesPiles(%v) : %d jutsus, maximum 2", p, len(suite))
			}
		}
	}
	for _, a := range [][]int{{}, {5}, {1, 2}, {1, 2, 3, 4, 5}} {
		if s := strategie.PetitesPiles(rouleau.New(a)); len(s) != 0 {
			t.Errorf("PetitesPiles(%v) : suite = %v, attendu vide", a, s)
		}
	}
	// Plus de 5 archives : hors de sa portée, il ne touche à rien.
	r := rouleau.New([]int{6, 5, 4, 3, 2, 1})
	if s := strategie.PetitesPiles(r); len(s) != 0 || !egal(r.Ciel, []int{6, 5, 4, 3, 2, 1}) {
		t.Errorf("PetitesPiles(6 archives) : suite=%v A=%v", s, r.Ciel)
	}
}

func TestShikamaru(t *testing.T) {
	for _, a := range [][]int{{}, {9}, {1, 2, 3}, {-4, 0, 8}} {
		if s := strategie.Shikamaru(a); len(s) != 0 {
			t.Errorf("Shikamaru(%v) : suite = %v, attendu vide", a, s)
		}
	}
	for n := 2; n <= 6; n++ {
		for _, p := range permutations(n) {
			verifierTri(t, "Shikamaru", p, strategie.Shikamaru(p))
		}
	}
	p := []int{100, -50, 7, 0, 2147483647, -2147483648}
	verifierTri(t, "Shikamaru", p, strategie.Shikamaru(p))

	entree := []int{3, 1, 2}
	strategie.Shikamaru(entree)
	if !egal(entree, []int{3, 1, 2}) {
		t.Errorf("Shikamaru a modifié son entrée : %v", entree)
	}

	hasard := rand.New(rand.NewSource(3))
	for _, n := range []int{10, 100, 500} {
		p := hasard.Perm(n)
		s := strategie.Shikamaru(p)
		verifierTri(t, "Shikamaru", p, s)
		for _, nom := range s {
			if nom != "pb" && nom != "pa" && nom != "ra" {
				t.Fatalf("Shikamaru : jutsu inattendu %q", nom)
			}
		}
	}
}

func TestKage(t *testing.T) {
	for _, a := range [][]int{{}, {5}, {1, 2}, {-3, 0, 9}} {
		if s := strategie.Kage(a); len(s) != 0 {
			t.Errorf("Kage(%v) : suite = %v, attendu vide", a, s)
		}
	}
	for n := 2; n <= 7; n++ {
		for _, p := range permutations(n) {
			verifierTri(t, "Kage", p, strategie.Kage(p))
		}
	}
	p := []int{100, -50, 7, 0, 2147483647, -2147483648}
	verifierTri(t, "Kage", p, strategie.Kage(p))

	entree := []int{3, 1, 2}
	strategie.Kage(entree)
	if !egal(entree, []int{3, 1, 2}) {
		t.Errorf("Kage a modifié son entrée : %v", entree)
	}

	hasard := rand.New(rand.NewSource(11))
	for _, n := range []int{10, 100, 500} {
		p := hasard.Perm(n)
		verifierTri(t, "Kage", p, strategie.Kage(p))
	}
}

// Objectifs du sujet : < 700 jutsus pour 100 archives, < 5500 pour 500.
func TestKageObjectifs(t *testing.T) {
	hasard := rand.New(rand.NewSource(42))
	for taille, limite := range map[int]int{100: 700, 500: 5500} {
		for essai := 0; essai < 5; essai++ {
			p := hasard.Perm(taille)
			if l := len(strategie.Kage(p)); l >= limite {
				t.Errorf("Kage sur %d archives : %d jutsus, objectif < %d", taille, l, limite)
			}
		}
	}
}

func TestMesures(t *testing.T) {
	if testing.Short() {
		t.Skip("mesures longues, sautées avec -short")
	}
	if err := strategie.MesurerStrategies(); err != nil {
		t.Errorf("MesurerStrategies : %v", err)
	}
	if err := strategie.MesurerKage(); err != nil {
		t.Errorf("MesurerKage : %v", err)
	}
}
