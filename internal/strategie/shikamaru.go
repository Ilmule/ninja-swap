// =====================================================================
//  internal/strategie/shikamaru.go — LA STRATÉGIE DE SHIKAMARU
//  Mission 3 (Jônin · Rang B)
// =====================================================================
//
//  Objectif : passer sous 1 100 jutsus pour 100 archives
//             et sous 8 500 pour 500.
//
//  À FAIRE :
//
//   ÉTAPE 1 — Normaliser
//   Remplacer chaque archive par son RANG : 0 pour la plus petite,
//   1 pour la suivante… jusqu'à n-1.
//   Exemple :  42 -7 100 3   →   2 0 3 1
//   (Le paquet sort est autorisé ici, et seulement ici.)
//
//   ÉTAPE 2 — Choisir AU MOINS une stratégie :
//
//   - Les escouades (chunks)
//     Sceller dans B d'abord tous les rangs 0 à 19, puis 20 à 39, etc.
//     Ensuite, invoquer vers A en prenant toujours la plus grande archive de B.
//     → Testez PLUSIEURS tailles d'escouade et gardez la meilleure.
//
//   - Le tri par bits (radix)
//     Pour chaque bit des rangs (du plus faible au plus fort) :
//     sceller dans B les archives dont le bit vaut 0, garder celles à 1,
//     puis tout invoquer. Ordre de grandeur : ≈ 1 080 pour 100.
//
//   ÉTAPE 3 — Mesurer et comparer avec la patrouille (mission 2).
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Pourquoi la normalisation simplifie-t-elle autant le problème ?
//   - Comment avez-vous choisi la taille de vos escouades ? Mesures à l'appui.
//
// =====================================================================

package strategie

import (
	"fmt"
	"math/rand"
	"sort"

	"ninja-swap/internal/rouleau"
)

func normaliser(nombres []int) []int {
	copie := make([]int, len(nombres))
	copy(copie, nombres)
	sort.Ints(copie)

	rangs := make([]int, len(nombres))

	for i, nombre := range nombres {
		for rang, valeur := range copie {
			if nombre == valeur {
				rangs[i] = rang
				break
			}
		}
	}

	return rangs
}

func Shikamaru(nombres []int) []string {
	r := rouleau.New(normaliser(nombres))
	operations := []string{}

	if r.EstTrie() {
		return operations
	}

	maximum := len(r.Ciel) - 1

	for bit := 0; (maximum >> bit) > 0; bit++ {
		nombreATraiter := len(r.Ciel)

		for i := 0; i < nombreATraiter; i++ {
			if (r.Ciel[0]>>bit)&1 == 0 {
				r.Pb()
				operations = append(operations, "pb")
			} else {
				r.Ra()
				operations = append(operations, "ra")
			}
		}

		for len(r.Terre) > 0 {
			r.Pa()
			operations = append(operations, "pa")
		}
	}

	return operations
}

func patrouilleReference(nombres []int) []string {
	r := rouleau.New(nombres)
	operations := []string{}

	if r.EstTrie() {
		return operations
	}

	for len(r.Ciel) > 0 {
		positionMinimum := 0

		for i := 1; i < len(r.Ciel); i++ {
			if r.Ciel[i] < r.Ciel[positionMinimum] {
				positionMinimum = i
			}
		}

		if positionMinimum <= len(r.Ciel)/2 {
			for i := 0; i < positionMinimum; i++ {
				r.Ra()
				operations = append(operations, "ra")
			}
		} else {
			rotations := len(r.Ciel) - positionMinimum

			for i := 0; i < rotations; i++ {
				r.Rra()
				operations = append(operations, "rra")
			}
		}

		r.Pb()
		operations = append(operations, "pb")
	}

	for len(r.Terre) > 0 {
		r.Pa()
		operations = append(operations, "pa")
	}

	return operations
}

func verifierOperations(nombres []int, operations []string) error {
	juge := rouleau.New(nombres)

	for _, operation := range operations {
		if err := juge.Appliquer(operation); err != nil {
			return err
		}
	}

	if !juge.EstTrie() {
		return fmt.Errorf("le resultat n'est pas trie")
	}

	return nil
}

func MesurerStrategies() error {
	hasard := rand.New(rand.NewSource(42))

	for _, taille := range []int{100, 500} {
		totalShikamaru := 0
		totalPatrouille := 0

		for essai := 0; essai < 20; essai++ {
			nombres := hasard.Perm(taille)

			resultatShikamaru := Shikamaru(nombres)
			resultatPatrouille := patrouilleReference(nombres)

			if err := verifierOperations(nombres, resultatShikamaru); err != nil {
				return fmt.Errorf(
					"Shikamaru, taille %d, essai %d : %w",
					taille, essai+1, err,
				)
			}

			if err := verifierOperations(nombres, resultatPatrouille); err != nil {
				return fmt.Errorf(
					"patrouille, taille %d, essai %d : %w",
					taille, essai+1, err,
				)
			}

			totalShikamaru += len(resultatShikamaru)
			totalPatrouille += len(resultatPatrouille)
		}

		fmt.Printf(
			"%d nombres : Shikamaru %.2f jutsus, patrouille %.2f jutsus\n",
			taille,
			float64(totalShikamaru)/20,
			float64(totalPatrouille)/20,
		)
	}

	return nil
}
