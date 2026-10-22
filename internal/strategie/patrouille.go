// =====================================================================
//  internal/strategie/patrouille.go — LA PATROUILLE DU PLUS FAIBLE
//  Mission 2 (Chûnin · Rang C)
// =====================================================================
//
//  Objectif : trier N'IMPORTE QUELLE taille de rouleau, même si ce n'est
//  pas encore efficace. Ce sera votre point de comparaison pour la suite.
//
//  À FAIRE :
//
//   1. Tant que le Rouleau du Ciel (A) n'est pas vide :
//      a. trouver la plus petite archive de A ;
//      b. la faire remonter en haut par le chemin LE PLUS COURT :
//         - si elle est dans la moitié haute → ra
//         - si elle est dans la moitié basse → rra
//      c. la sceller dans B (pb).
//
//   2. Invoquer toutes les archives de B vers A (pa).
//
//   3. MESURER : moyenne de jutsus sur 20 tirages de 100 archives,
//      puis sur 20 tirages de 500 (utilisez bench.sh).
//      Ordre de grandeur attendu : ≈ 1 450 pour 100, ≈ 32 000 pour 500.
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Pourquoi cette méthode est-elle en O(n²) ? Expliquez avec vos mots.
//   - Le programme est-il lent à CALCULER, ou seulement « cher » en jutsus ?
//
// =====================================================================

package strategie

import (
	"math/rand"
	"testing"

	"ninja-swap/internal/rouleau"
)

func Patrouille(r *rouleau.Rouleaux) []string {
	suite := []string{}
	if r.EstTrie() {
		return suite
	}
	for len(r.Ciel) > 0 {
		remonterMin(r, &suite) // ra ou rra : le chemin le plus court
		jouer(r, &suite, "pb")
	}
	for len(r.Terre) > 0 {
		jouer(r, &suite, "pa")
	}
	return suite
}

func tirage(n int) []int {
	return rand.Perm(n)
}

func TestPatrouilleTrie(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 7, 50} {
		r := rouleau.New(tirage(n))
		Patrouille(r)
		if !r.EstTrie() {
			t.Fatalf("n=%d : A=%v B=%v", n, r.Ciel, r.Terre)
		}
	}
}

func TestPatrouilleMoyenne(t *testing.T) {
	for _, n := range []int{100, 500} {
		total := 0
		for i := 0; i < 20; i++ {
			total += len(Patrouille(rouleau.New(tirage(n))))
		}
		t.Logf("n=%d : moyenne %d jutsus", n, total/20)
	}
}
