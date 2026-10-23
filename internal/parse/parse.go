// =====================================================================
//  internal/parse/parse.go — DÉJOUER LES PIÈGES DE L'AKATSUKI
//  Mission 0 (Académie)
// =====================================================================
//
//  Rôle : transformer les arguments texte en une liste d'entiers valide,
//  ou signaler une erreur. Utilisé par ninja_swap ET par le juge.
//
//  À FAIRE : une fonction qui reçoit les arguments (une slice de string)
//  et renvoie soit la liste d'entiers, soit une erreur.
//
//  Elle doit refuser (renvoyer une erreur) :
//
//   - un argument non numérique ............ 4 kunai 2
//   - un signe seul ou mal placé ........... -   /   4-2
//   - une valeur hors des limites d'un int32  2147483648   /   -2147483649
//   - un doublon ........................... 1 5 1
//
//  Elle doit accepter :
//
//   - les nombres négatifs ................. -5 0 -12
//   - les bornes exactes ................... 2147483647  /  -2147483648
//   - une seule archive .................... 42
//
//  Outils utiles : strconv.Atoi, math.MaxInt32, math.MinInt32,
//  et une map pour repérer les doublons.
//
//  Ne JAMAIS afficher "Error" ici : c'est le rôle des main.go.
//  Ici, on renvoie seulement l'erreur.
//
//  MISSION SECRÈTE « Reconnaissance urbaine » (facultatif) :
//  prévoir un mode où les doublons sont autorisés (option --doublons).
//
// =====================================================================

package parse

import (
	"errors"
	"fmt"
	"strconv"
)

func Args(args []string) ([]int, error) {
	return parse(args, false)
}

func ArgsAvecDoublons(args []string) ([]int, error) {
	return parse(args, true)
}

func parse(args []string, doublons bool) ([]int, error) {
	nums := make([]int, 0, len(args))
	vus := make(map[int]bool, len(args))

	for _, s := range args {
		n64, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("argument invalide %q : %w", s, err)
		}
		n := int(n64)

		if !doublons {
			if vus[n] {
				return nil, errors.New("doublon : " + s)
			}
			vus[n] = true
		}
		nums = append(nums, n)
	}
	return nums, nil
}