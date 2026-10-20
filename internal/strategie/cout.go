// =====================================================================
//  internal/strategie/cout.go — LE CHAKRA PONDÉRÉ
//  MISSION SECRÈTE « Les Bâtisseurs » (facultative · Rang A · +6)
// =====================================================================
//
//  Inspirée des Chantiers de l'Atlantique : un levage de grue coûte bien
//  plus cher qu'un simple déplacement au sol.
//
//  Ne commencez cette mission QU'APRÈS avoir atteint le rang Chûnin.
//
//  À FAIRE :
//
//   1. Lire un fichier de coûts passé avec l'option --cout fichier.txt.
//      Format conseillé : une ligne par jutsu, « nom coût ».
//      Exemple :
//          pa 3
//          pb 3
//          ra 1
//          sa 2
//      Un jutsu absent du fichier coûte 1.
//
//   2. Adapter la stratégie de la mission 4 pour minimiser le COÛT TOTAL,
//      et plus seulement le nombre de jutsus.
//
//   3. Montrer, mesures à l'appui, que la meilleure stratégie CHANGE
//      quand les coûts changent (comparez au moins deux fichiers de coûts).
//
//  À expliquer dans le README : pourquoi un « algorithme optimal »
//  n'existe que par rapport à une mesure du coût.
//
// =====================================================================

package strategie
