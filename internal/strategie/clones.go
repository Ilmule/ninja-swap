// =====================================================================
//  internal/strategie/clones.go — LA COURSE DES CLONES
//  MISSION SECRÈTE « Kage Bunshin » (facultative · Rang A · +6)
// =====================================================================
//
//  Inspirée de SNCF Connect & Tech (site de Nantes) : à grande échelle,
//  faire travailler plusieurs calculs en parallèle change tout.
//
//  Ne commencez cette mission QU'APRÈS avoir atteint le rang Chûnin.
//
//  À FAIRE :
//
//   1. Lancer plusieurs stratégies EN PARALLÈLE avec des goroutines :
//      escouades de taille 10 à 40, radix, coût minimal…
//      Chaque clone travaille sur SA PROPRE COPIE des archives.
//
//   2. Récupérer les résultats par un channel et garder la suite de
//      jutsus la plus courte.
//
//   3. Imposer un temps limite (par exemple 2 s) avec context.WithTimeout :
//      on garde la meilleure réponse disponible quand le temps est écoulé.
//
//   4. Mesurer le gain avec un benchmark (go test -bench) et repérer la
//      fonction la plus coûteuse avec pprof.
//
//   5. Tester la montée en charge jusqu'à 10 000 archives.
//
//  À expliquer dans le README : la différence entre « plus de calcul »
//  et « meilleur résultat ».
//
// =====================================================================

package strategie
