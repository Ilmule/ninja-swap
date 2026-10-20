// =====================================================================
//  cmd/juge/main.go — LE PARCHEMIN DE L'EXAMINATEUR
//  Mission 0 (Académie)
// =====================================================================
//
//  Rôle : vérifier qu'une suite de jutsus trie bien les archives.
//
//  Utilisation :
//      ./ninja_swap 2 3 1 4 | ./juge 2 3 1 4
//
//  À FAIRE :
//
//   1. Lire et valider les arguments EXACTEMENT comme ninja_swap
//      (réutilisez internal/parse, ne recopiez pas le code !).
//      En cas d'erreur → "Error" sur os.Stderr, os.Exit(1).
//
//   2. Construire le Rouleau du Ciel (A) et le Rouleau de la Terre (B, vide).
//
//   3. Lire l'entrée standard LIGNE PAR LIGNE (bufio.Scanner).
//      Pour chaque ligne, appliquer le jutsu correspondant
//      (sa, sb, ss, pa, pb, ra, rb, rr, rra, rrb, rrr).
//      Jutsu inconnu → "Error" sur os.Stderr, os.Exit(1).
//
//   4. À la fin de l'entrée :
//      - A trié par ordre croissant ET B vide → afficher "OK"
//      - sinon                                → afficher "KO"
//
//  CAS À TESTER ABSOLUMENT :
//   - printf 'pb\nsa\npa\nsa\n' | ./juge 2 3 1 4      → OK
//   - printf 'pb\n' | ./juge 1 2 3                    → KO (B n'est pas vide)
//   - printf 'kunai\n' | ./juge 1 2                   → Error
//   - ./juge 1 2 3 < /dev/null                        → OK (déjà trié)
//
//  Conseil d'Iruka : le juge est votre meilleur allié pour toute la suite.
//  Un juge faux, et toutes vos mesures deviennent fausses. Soignez-le.
//
// =====================================================================

package main

func main() {
	// TODO : à vous de jouer, genin.
}
