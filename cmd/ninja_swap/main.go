// =====================================================================
//  cmd/ninja_swap/main.go — LE PROGRAMME PRINCIPAL
//  Missions 0 à 4
// =====================================================================
//
//  Rôle : recevoir les archives en arguments et afficher, une par ligne,
//  la suite de jutsus la plus courte possible pour les trier.
//
//  Ce fichier doit rester COURT. Il ne fait que relier les pièces :
//  toute la logique vit dans internal/.
//
//  À FAIRE, dans l'ordre :
//
//   1. Lire les arguments (os.Args[1:]) et les confier à internal/parse.
//      - Aucun argument            → ne rien afficher, sortir avec le code 0.
//      - Erreur de validation      → afficher "Error" sur os.Stderr,
//                                    puis os.Exit(1). Rien sur la sortie standard.
//
//   2. Construire les deux rouleaux avec internal/rouleau :
//      - Rouleau du Ciel (A) : les archives, le PREMIER argument en haut.
//      - Rouleau de la Terre (B) : vide.
//
//   3. Si A est déjà trié → ne rien afficher.
//
//   4. Choisir la stratégie selon le nombre d'archives (internal/strategie) :
//      - 2 ou 3 archives   → votre solution optimale (mission 1)
//      - 4 ou 5 archives   → votre solution « petites piles » (mission 1)
//      - 6 et plus         → votre meilleure stratégie (missions 2, 3 puis 4)
//
//   5. Afficher chaque jutsu sur sa propre ligne (fmt.Println).
//
//  RAPPELS :
//   - Le programme ne doit JAMAIS paniquer (pas de panic, pas d'index hors limite).
//   - Testez toujours avec le juge :  ./ninja_swap $ARG | ./juge $ARG
//
//  MISSIONS SECRÈTES qui modifient ce fichier (facultatif) :
//   - Les Bâtisseurs : option --cout fichier.txt (voir internal/strategie/cout.go)
//   - Reconnaissance urbaine : option --doublons (voir internal/parse/parse.go)
//
// =====================================================================

package main

import (
	"bufio"
	"fmt"
	"os"

	"ninja-swap/internal/parse"
	"ninja-swap/internal/rouleau"
	"ninja-swap/internal/strategie"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	nombres, err := parse.Args(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	r := rouleau.New(nombres)
	if r.EstTrie() {
		return
	}

	var suite []string
	if n := len(r.Ciel); n <= 5 {
		suite = strategie.PetitesPiles(r)
	} else {
		suite = strategie.Patrouille(r)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, jutsu := range suite {
		fmt.Fprintln(out, jutsu)
	}

	
}
