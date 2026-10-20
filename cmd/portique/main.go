// =====================================================================
//  cmd/portique/main.go — LE PORTIQUE DU PORT
//  MISSION SECRÈTE « Pays des Vagues » (facultative · Rang S · +8)
// =====================================================================
//
//  Inspirée du terminal à conteneurs de Montoir-de-Bretagne (TGO).
//  Ne commencez cette mission QU'APRÈS avoir atteint le rang Chûnin.
//
//  LE PROBLÈME :
//   - Plusieurs piles de conteneurs, chacune limitée en hauteur.
//   - Chaque conteneur porte le numéro de son ordre de départ (1, 2, 3…).
//   - Un seul geste possible : déplacer le conteneur du HAUT d'une pile
//     vers le HAUT d'une autre pile (sans dépasser la hauteur max).
//   - Il faut sortir le n°1, puis le n°2, etc., avec le moins possible
//     de déplacements inutiles (les « remaniements »).
//
//  FORMAT D'ENTRÉE CONSEILLÉ (un fichier passé en argument) :
//      4             ← hauteur maximale d'une pile
//      3 1 5         ← pile 1, du BAS vers le HAUT
//      2 6           ← pile 2
//      4             ← pile 3
//                    ← ligne vide = pile vide
//
//  SORTIE ATTENDUE (une action par ligne) :
//      1->3          ← déplacer le haut de la pile 1 vers la pile 3
//      out           ← le conteneur au sommet sort (c'est le prochain attendu)
//
//  À FAIRE :
//   1. Lire le fichier, vérifier qu'il est valide.
//   2. Règle simple : quand un conteneur gêne, le déplacer vers la pile
//      dont le sommet partira le plus tard.
//   3. Imaginer AU MOINS une autre règle.
//   4. Comparer les deux règles sur 50 terminaux générés aléatoirement
//      et présenter les moyennes de remaniements dans le README.
//
//  À savoir : ce problème est réputé trop difficile pour être résolu
//  parfaitement à grande échelle. On y répond par des heuristiques mesurées.
//
// =====================================================================

package main

func main() {
	// TODO : mission secrète.
}
