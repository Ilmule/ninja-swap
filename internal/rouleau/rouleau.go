// =====================================================================
//  internal/rouleau/rouleau.go — LES DEUX ROULEAUX ET LES 11 JUTSUS
//  Mission 0 (Académie)
// =====================================================================
//
//  Rôle : représenter les deux piles et coder chaque jutsu une fois pour
//  toutes. ninja_swap, le juge et toutes vos stratégies utiliseront CE
//  package : ne réécrivez jamais un jutsu ailleurs.
//
//  À FAIRE :
//
//   1. Choisir une représentation des rouleaux.
//      Convention conseillée : une slice d'entiers, l'indice 0 = le HAUT.
//
//   2. Prévoir une structure qui regroupe le Rouleau du Ciel (A) et le
//      Rouleau de la Terre (B), et qui COMPTE les jutsus lancés.
//      (Utile pour vos stratégies et vos mesures.)
//
//   3. Coder les 11 jutsus, une fonction (ou méthode) chacun :
//
//      sa  / sb  / ss    Kawarimi      échange les 2 archives du haut
//      pa                Kuchiyose     haut de B → haut de A
//      pb                Fûin          haut de A → haut de B
//      ra  / rb  / rr    Rasengan      le haut passe tout en bas
//      rra / rrb / rrr   Rasengan inv. le bas remonte tout en haut
//
//      Les jutsus « en duo » (ss, rr, rrr) comptent pour UN seul jutsu.
//
//   4. Chaque jutsu doit être SANS DANGER :
//      sur un rouleau vide ou à 1 archive, il ne fait rien… et ne plante pas.
//
//   5. Prévoir un moyen d'appliquer un jutsu à partir de son nom ("sa", "pb"…).
//      Le juge en aura besoin. Nom inconnu → renvoyer une erreur.
//
//   6. Prévoir une fonction qui dit si A est trié et B vide.
//
//  Testez tout dans rouleau_test.go AVANT de passer aux stratégies.
//
// =====================================================================

package rouleau

type Vertex struct {
		Terre int
		Ciel int
}

var Terre []int 
var Ciel []int 


func sa() {
	if len(Ciel) < 2 {
		return
	}
	Ciel[0], Ciel[1] = Ciel[1], Ciel[0]
}

func sb() {
	if len(Terre) < 2 {
		return
	}
	Terre[0], Terre[1] = Terre[1], Terre[0]
}

func ss() {
	sa()
	sb()
}

func pa() {
	if len(Terre) == 0 {
		return
	}
	haut := Terre[0]
	Terre = Terre[1:]
	Ciel = append([]int{haut}, Ciel...)
}

func pb() {
	if len(Ciel) == 0 {
		return
	}
	haut := Ciel[0]
	Ciel = Ciel[1:]
	Terre = append([]int{haut}, Terre...)
}

func ra() {

}

func rb() {

}

func rr() {

}

func rra() {

}

func rrb() {

}

func rrr() {

}