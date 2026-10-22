// =====================================================================
//  internal/strategie/rocklee.go — L'ENTRAÎNEMENT DE ROCK LEE
//  Mission 1 (Genin · Rang D)
// =====================================================================
//
//  Rock Lee ne maîtrise aucun ninjutsu : il compense par l'effort pur.
//  Ce fichier contient deux choses.
//
//  PARTIE A — Les petites piles, résolues à la main
//
//   1. 3 archives : il n'existe que 6 ordres possibles.
//      Trouvez SUR PAPIER la meilleure suite de jutsus pour chacun,
//      puis codez-les. Objectif : 2 jutsus maximum.
//
//   2. 5 archives : objectif 12 jutsus maximum.
//      Idée de Gai-sensei : sceller les 2 plus petites dans B (pb),
//      trier les 3 restantes avec la partie 1, puis les invoquer (pa).
//      4 archives : même idée, en ne scellant que la plus petite.
//
//      Le rang Kage exige 8 jutsus au maximum pour 5 archives : c'est le
//      vrai minimum. Pour l'atteindre, utilisez le brute-force (partie B).
//
//  PARTIE B — Le brute-force
//
//   3. Essayer toutes les suites de jutsus de longueur 1, puis 2, puis 3…
//      jusqu'à en trouver une qui trie. C'est forcément la plus courte.
//      Indice : une file (parcours en largeur) évite de tout réessayer.
//
//   4. MESURER le temps de calcul de 3 à 8 archives
//      et noter les résultats dans le journal d'optimisation du README.
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Que devient le temps de calcul entre 5 et 8 archives ? Pourquoi ?
//   - Quelle est la complexité (Big-O) de votre brute-force ?
//
// =====================================================================

package strategie

import (
	"sort"

	"ninja-swap/internal/rouleau"
)

func jouer(r *rouleau.Rouleaux, suite *[]string, noms ...string) {
	for _, nom := range noms {
		if err := r.Appliquer(nom); err != nil {
			continue // nom connu en dur : ne devrait jamais arriver
		}
		*suite = append(*suite, nom)
	}
}

func trierDeux(r *rouleau.Rouleaux, suite *[]string) {
	if len(r.Ciel) == 2 && r.Ciel[0] > r.Ciel[1] {
		jouer(r, suite, "sa")
	}
}

func trierTrois(r *rouleau.Rouleaux, suite *[]string) {
	if len(r.Ciel) != 3 {
		return
	}
	a, b, c := r.Ciel[0], r.Ciel[1], r.Ciel[2]

	switch {
	case a < b && b < c:
		// déjà trié
	case a < c && c < b: // 1 3 2
		jouer(r, suite, "rra", "sa")
	case b < a && a < c: // 2 1 3
		jouer(r, suite, "sa")
	case c < a && a < b: // 2 3 1
		jouer(r, suite, "rra")
	case b < c && c < a: // 3 1 2
		jouer(r, suite, "ra")
	default: // 3 2 1
		jouer(r, suite, "sa", "rra")
	}
}

func remonterMin(r *rouleau.Rouleaux, suite *[]string) {
	n := len(r.Ciel)
	if n < 2 {
		return
	}
	idx := 0
	for i, v := range r.Ciel {
		if v < r.Ciel[idx] {
			idx = i
		}
	}
	if idx <= n/2 {
		for i := 0; i < idx; i++ {
			jouer(r, suite, "ra")
		}
	} else {
		for i := 0; i < n-idx; i++ {
			jouer(r, suite, "rra")
		}
	}
}

func sceller(r *rouleau.Rouleaux, suite *[]string, k int) {
	for i := 0; i < k; i++ {
		remonterMin(r, suite)
		jouer(r, suite, "pb")
	}
	trierTrois(r, suite)
	for i := 0; i < k; i++ {
		jouer(r, suite, "pa")
	}
}

func cinq(r *rouleau.Rouleaux, suite *[]string) { sceller(r, suite, 2) }

func quatre(r *rouleau.Rouleaux, suite *[]string) { sceller(r, suite, 1) }

const tailleMaxBrute = 8

var tousLesJutsus = []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

type etat struct {
	r      *rouleau.Rouleaux
	parent int
	jutsu  string
}

func cle(r *rouleau.Rouleaux) string {
	b := make([]byte, 0, len(r.Ciel)+len(r.Terre)+1)
	for _, v := range r.Ciel {
		b = append(b, byte(v))
	}
	b = append(b, 255)
	for _, v := range r.Terre {
		b = append(b, byte(v))
	}
	return string(b)
}

func copie(r *rouleau.Rouleaux) *rouleau.Rouleaux {
	c := rouleau.New(r.Ciel)
	c.Terre = append([]int{}, r.Terre...)
	return c
}

func BruteForce(r *rouleau.Rouleaux) []string {
	suite := []string{}
	n := len(r.Ciel)
	if r.EstTrie() || n > tailleMaxBrute || len(r.Terre) != 0 {
		return suite
	}

	tri := append([]int{}, r.Ciel...)
	sort.Ints(tri)
	rang := make(map[int]int, n)
	for i, v := range tri {
		rang[v] = i
	}
	norm := make([]int, n)
	for i, v := range r.Ciel {
		norm[i] = rang[v]
	}

	depart := rouleau.New(norm)
	file := []etat{{r: depart, parent: -1}}
	vus := map[string]bool{cle(depart): true}

	for i := 0; i < len(file); i++ {
		courant := file[i].r
		file[i].r = nil
		for _, nom := range tousLesJutsus {
			suivant := copie(courant)
			suivant.Appliquer(nom)
			k := cle(suivant)
			if vus[k] {
				continue
			}
			vus[k] = true
			if suivant.EstTrie() {
				chemin := []string{nom}
				for j := i; file[j].parent != -1; j = file[j].parent {
					chemin = append(chemin, file[j].jutsu)
				}
				for a, b := 0, len(chemin)-1; a < b; a, b = a+1, b-1 {
					chemin[a], chemin[b] = chemin[b], chemin[a]
				}
				for _, j := range chemin {
					jouer(r, &suite, j)
				}
				return suite
			}
			file[i].r = nil
			file = append(file, etat{r: suivant, parent: i, jutsu: nom})
		}
	}
	return suite 
}

func PetitesPiles(r *rouleau.Rouleaux) []string {
	suite := []string{}
	if r.EstTrie() {
		return suite
	}
	switch len(r.Ciel) {
	case 2:
		trierDeux(r, &suite)
	case 3:
		trierTrois(r, &suite)
	case 4, 5:
		return BruteForce(r)
	}
	return suite
}