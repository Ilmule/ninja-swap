// =====================================================================
//  internal/strategie/kage.go — LE CONTRÔLE DU CHAKRA
//  Mission 4 (Kage · Rang A)
// =====================================================================
//
//  Objectif : moins de 700 jutsus pour 100 archives
//             et moins de 5 500 pour 500.
//  Le temps de calcul doit rester de quelques secondes au plus pour 500.
//
//  À FAIRE :
//
//   1. Pour chaque archive candidate, CALCULER SON COÛT : le nombre de
//      rotations dans A + le nombre de rotations dans B pour la placer
//      au bon endroit dans l'autre rouleau.
//
//   2. À chaque tour, choisir l'archive LA MOINS CHÈRE, puis exécuter.
//
//   3. Fusionner les rotations : si A et B doivent tourner dans le même
//      sens, utiliser rr ou rrr au lieu de deux jutsus séparés.
//
//   4. Vérifier le temps de calcul pour 500 archives.
//
//  Conseil de Jiraiya : UNE amélioration à la fois. Mesurez, gardez-la
//  si elle aide, notez-la dans le journal, puis passez à la suivante.
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Que coûte en calcul la recherche de l'archive la moins chère à chaque
//     tour ? Quelle complexité totale ?
//   - Quelle amélioration vous a fait gagner le plus ? Laquelle vous a déçus ?
//
// =====================================================================

package strategie
