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
