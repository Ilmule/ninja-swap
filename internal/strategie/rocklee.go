// =====================================================================
//  internal/strategie/rocklee.go — L'ENTRAÎNEMENT DE ROCK LEE
//  Mission 1 (Genin · Rang D)
// =====================================================================
//
//  Rock Lee ne maîtrise aucun ninjutsu : il compense par l'effort pur.
//  Ce fichier contient deux choses.
//
//  PARTIE A — Les petites piles, résolues à la main
//
//   1. 3 archives : il n'existe que 6 ordres possibles.
//      Trouvez SUR PAPIER la meilleure suite de jutsus pour chacun,
//      puis codez-les. Objectif : 2 jutsus maximum.
//
//   2. 5 archives : objectif 12 jutsus maximum.
//      Idée de Gai-sensei : sceller les 2 plus petites dans B (pb),
//      trier les 3 restantes avec la partie 1, puis les invoquer (pa).
//      4 archives : même idée, en ne scellant que la plus petite.
//
//      Le rang Kage exige 8 jutsus au maximum pour 5 archives : c'est le
//      vrai minimum. Pour l'atteindre, utilisez le brute-force (partie B).
//
//  PARTIE B — Le brute-force
//
//   3. Essayer toutes les suites de jutsus de longueur 1, puis 2, puis 3…
//      jusqu'à en trouver une qui trie. C'est forcément la plus courte.
//      Indice : une file (parcours en largeur) évite de tout réessayer.
//
//   4. MESURER le temps de calcul de 3 à 8 archives
//      et noter les résultats dans le journal d'optimisation du README.
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Que devient le temps de calcul entre 5 et 8 archives ? Pourquoi ?
//   - Quelle est la complexité (Big-O) de votre brute-force ?
//
// =====================================================================

package strategie
