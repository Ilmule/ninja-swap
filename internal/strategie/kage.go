// =====================================================================
//  internal/strategie/kage.go — LE CONTRÔLE DU CHAKRA
//  Mission 4 (Kage · Rang A)
// =====================================================================
//
//  Objectif : moins de 700 jutsus pour 100 archives
//             et moins de 5 500 pour 500.
//  Le temps de calcul doit rester de quelques secondes au plus pour 500.
//
//  À FAIRE :
//
//   1. Pour chaque archive candidate, CALCULER SON COÛT : le nombre de
//      rotations dans A + le nombre de rotations dans B pour la placer
//      au bon endroit dans l'autre rouleau.
//
//   2. À chaque tour, choisir l'archive LA MOINS CHÈRE, puis exécuter.
//
//   3. Fusionner les rotations : si A et B doivent tourner dans le même
//      sens, utiliser rr ou rrr au lieu de deux jutsus séparés.
//
//   4. Vérifier le temps de calcul pour 500 archives.
//
//  Conseil de Jiraiya : UNE amélioration à la fois. Mesurez, gardez-la
//  si elle aide, notez-la dans le journal, puis passez à la suivante.
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Que coûte en calcul la recherche de l'archive la moins chère à chaque
//     tour ? Quelle complexité totale ?
//   - Quelle amélioration vous a fait gagner le plus ? Laquelle vous a déçus ?
//
// =====================================================================

package strategie

import (
	"fmt"
	"math/rand"
	"time"

	"ninja-swap/internal/rouleau"
)

type mouvementKage struct {
	rotationsA int
	rotationsB int
	cout       int
}

// Kage attend des nombres distincts, deja valides par le package parse.
func Kage(nombres []int) []string {
	dejaTrie := true
	for i := 1; i < len(nombres); i++ {
		if nombres[i-1] > nombres[i] {
			dejaTrie = false
			break
		}
	}
	if dejaTrie {
		return []string{}
	}

	rangs := normaliserKage(nombres)
	var meilleuresOperations []string

	// Comparer quatre positions de depart, rotations initiales comprises.
	for depart := 0; depart < 4 && depart < len(rangs); depart++ {
		operations := kageDepuis(rangs, depart)
		if meilleuresOperations == nil || len(operations) < len(meilleuresOperations) {
			meilleuresOperations = operations
		}
	}

	return meilleuresOperations
}

func normaliserKage(nombres []int) []int {
	rangs := make([]int, len(nombres))
	for i, nombre := range nombres {
		for _, autre := range nombres {
			if autre < nombre {
				rangs[i]++
			}
		}
	}
	return rangs
}

func kageDepuis(rangs []int, depart int) []string {
	r := rouleau.New(rangs)
	operations := []string{}

	for i := 0; i < depart; i++ {
		r.Ra()
		operations = append(operations, "ra")
	}

	// Deux elements forment toujours un ordre circulaire dans B.
	for i := 0; i < 2 && len(r.Ciel) > 0; i++ {
		r.Pb()
		operations = append(operations, "pb")
	}

	for len(r.Ciel) > 0 {
		cibles := ciblesKage(r.Terre, len(rangs))
		meilleur := mouvementKage{cout: len(r.Ciel) + len(r.Terre) + 1}

		for position, valeur := range r.Ciel {
			candidat := choisirRotationsKage(
				position, len(r.Ciel),
				cibles[valeur], len(r.Terre),
			)
			if candidat.cout < meilleur.cout {
				meilleur = candidat
			}
			if meilleur.cout == 0 {
				break
			}
		}

		operations = tournerKage(
			r, meilleur.rotationsA, meilleur.rotationsB, operations,
		)
		r.Pb()
		operations = append(operations, "pb")
	}

	// Placer le maximum de B en haut avant de tout ramener dans A.
	positionMaximum := 0
	for i, valeur := range r.Terre {
		if valeur == len(rangs)-1 {
			positionMaximum = i
			break
		}
	}

	rotationsB := positionMaximum
	if positionMaximum > len(r.Terre)/2 {
		rotationsB -= len(r.Terre)
	}
	operations = tournerKage(r, 0, rotationsB, operations)

	for len(r.Terre) > 0 {
		r.Pa()
		operations = append(operations, "pa")
	}

	return operations
}

// B reste trie dans l'ordre decroissant, avec un decalage circulaire.
// La cible est le plus grand rang inferieur, ou le maximum de B.
func ciblesKage(b []int, nombreRangs int) []int {
	positions := make([]int, nombreRangs)
	for i := range positions {
		positions[i] = -1
	}

	maximum := -1
	positionMaximum := 0
	for position, rang := range b {
		positions[rang] = position
		if rang > maximum {
			maximum = rang
			positionMaximum = position
		}
	}

	cibles := make([]int, nombreRangs)
	positionCible := positionMaximum
	for rang := 0; rang < nombreRangs; rang++ {
		cibles[rang] = positionCible
		if positions[rang] != -1 {
			positionCible = positions[rang]
		}
	}
	return cibles
}

// Une rotation positive utilise ra/rb, une negative rra/rrb.
// Examiner les quatre combinaisons permet de profiter de rr et rrr.
func choisirRotationsKage(positionA, tailleA, positionB, tailleB int) mouvementKage {
	retourA := 0
	if positionA > 0 {
		retourA = positionA - tailleA
	}
	retourB := 0
	if positionB > 0 {
		retourB = positionB - tailleB
	}

	meilleur := mouvementKage{cout: tailleA + tailleB + 1}
	for _, a := range []int{positionA, retourA} {
		for _, b := range []int{positionB, retourB} {
			cout := coutRotationsKage(a, b)
			if cout < meilleur.cout {
				meilleur = mouvementKage{rotationsA: a, rotationsB: b, cout: cout}
			}
		}
	}
	return meilleur
}

func coutRotationsKage(a, b int) int {
	absA, absB := a, b
	if absA < 0 {
		absA = -absA
	}
	if absB < 0 {
		absB = -absB
	}

	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		if absA > absB {
			return absA
		}
		return absB
	}
	return absA + absB
}

func tournerKage(r *rouleau.Rouleaux, a, b int, operations []string) []string {
	for a > 0 && b > 0 {
		r.Rr()
		operations = append(operations, "rr")
		a--
		b--
	}
	for a < 0 && b < 0 {
		r.Rrr()
		operations = append(operations, "rrr")
		a++
		b++
	}
	for a > 0 {
		r.Ra()
		operations = append(operations, "ra")
		a--
	}
	for a < 0 {
		r.Rra()
		operations = append(operations, "rra")
		a++
	}
	for b > 0 {
		r.Rb()
		operations = append(operations, "rb")
		b--
	}
	for b < 0 {
		r.Rrb()
		operations = append(operations, "rrb")
		b++
	}
	return operations
}

// A appeler separement du programme qui affiche les jutsus pour le juge.
func MesurerKage() error {
	hasard := rand.New(rand.NewSource(42))

	for _, taille := range []int{100, 500} {
		total, minimum, maximum := 0, -1, 0
		var tempsMaximum time.Duration

		for essai := 0; essai < 20; essai++ {
			nombres := hasard.Perm(taille)
			debut := time.Now()
			operations := Kage(nombres)
			duree := time.Since(debut)
			if duree > tempsMaximum {
				tempsMaximum = duree
			}

			juge := rouleau.New(nombres)
			for _, operation := range operations {
				if err := juge.Appliquer(operation); err != nil {
					return err
				}
			}
			if !juge.EstTrie() {
				return fmt.Errorf("tri incorrect : taille %d, essai %d", taille, essai+1)
			}

			compte := len(operations)
			total += compte
			if minimum == -1 || compte < minimum {
				minimum = compte
			}
			if compte > maximum {
				maximum = compte
			}
		}

		fmt.Printf(
			"%d nombres : min=%d, moyenne=%.2f, max=%d, temps maximal=%s\n",
			taille, minimum, float64(total)/20, maximum,
			tempsMaximum.Round(time.Microsecond),
		)

		limite := 700
		if taille == 500 {
			limite = 5500
		}
		if maximum >= limite {
			return fmt.Errorf("objectif manque : %d operations pour %d nombres", maximum, taille)
		}
	}
	return nil
}


