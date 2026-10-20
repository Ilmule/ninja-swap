// =====================================================================
//  cmd/reconnaissance/main.go — LES VÉLOS DE NANTES
//  MISSION SECRÈTE « Reconnaissance urbaine » (facultative · Rang B · +6)
// =====================================================================
//
//  Inspirée de l'open data Naolib de Nantes Métropole (format GBFS).
//  Ne commencez cette mission QU'APRÈS avoir atteint le rang Chûnin.
//
//  DONNÉES : le Hokage vous fournit un instantané de station_status.json.
//  Pas besoin de réseau. Placez-le dans ce dossier.
//
//  À FAIRE :
//
//   1. Lire station_status.json avec encoding/json.
//      Pour chaque station, récupérer son identifiant (station_id) et
//      le nombre de vélos disponibles (num_bikes_available).
//      Attention aux champs manquants : une station incomplète est ignorée.
//
//   2. Les données réelles contiennent des ÉGALITÉS : plusieurs stations
//      peuvent avoir le même nombre de vélos. Ajouter un mode --doublons
//      à ninja_swap ET au juge :
//      en cas d'égalité, la normalisation départage par ordre d'arrivée.
//
//   3. Produire le classement des 10 stations les plus vides, à visiter
//      en priorité pour le rééquilibrage.
//
//  À expliquer dans le README : quelle hypothèse de départ est tombée,
//  et ce que vous avez dû changer dans votre algorithme.
//
// =====================================================================

package main

func main() {
	// TODO : mission secrète.
}
