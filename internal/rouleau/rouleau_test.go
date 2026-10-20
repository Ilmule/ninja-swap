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
