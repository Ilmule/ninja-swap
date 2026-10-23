# Rouleau de l'équipe n° jsp

> **Consigne :** ce README est le « Rouleau de l'équipe ». Il est noté (documentation et analyse de complexité).
> Remplacez chaque bloc de consignes en italique par votre propre contenu, et supprimez les consignes à la fin.
>
> **Rendu :** vendredi 23 octobre à 11 h précises, avec `git tag rendu` puis `git push --tags`.
> Poussez votre travail à la fin de CHAQUE séance, par les deux membres de l'équipe.
> Donnez accès à ce dépôt au Hokage et à Rémy dès la première séance.

**Genin :** _Emeric Van Autryve_ · _Hugo Laude_

---

## 1. Compiler, lancer, tester

_Donnez les commandes exactes pour :_

## 1. Compiler, lancer, tester

```bash
go build -o ninja_swap ./cmd/ninja_swap    # compiler
go build -o juge ./cmd/juge

ARG="2 3 1 4"
./ninja_swap $ARG | ./juge $ARG            # exemple + vérification : affiche OK

go test ./...                              # tests unitaires
```

`bench.sh` : _à compléter._

## 2. Notre stratégie finale

_Expliquez votre stratégie finale pour qu'un élève de B1 qui n'a pas fait le projet la comprenne._
_Un petit exemple déroulé à la main (5 ou 6 archives) vaut mieux qu'un long paragraphe._

## 3. Journal d'optimisation

_Une ligne par version. Ne supprimez jamais une ligne : le chemin compte autant que l'arrivée._

| Version | Stratégie | Big-O (jutsus) | Moy. 100 | Moy. 500 | Ce que nous avons appris |
|---|---|---|---|---|---|
| v1 | Rock Lee (brute-force) | O(n!) | — | — | Le brute force est un bon début pour voir ce que doit faire le programme mais ce n'est pas vraiment efficace |
| v2 | Patrouille | | | | |
| v3 | | | | | |

## 4. Questions de réflexion

### Mission 1 — Rock Lee
- _Que devient le temps de calcul du brute-force entre 5 et 7 archives ? Pourquoi ?_
Il se démultiplie car on à d'un seul coup bien plus de comparaisons et d'échanges à faire.
- _Quelle est la complexité (Big-O) de votre brute-force ?_
O(n!)

### Mission 2 — La patrouille
- _Pourquoi cette méthode est-elle en O(n²) ?_
car plus on as d'archive plus on as de chakra (fonction) c'est donc la caractéristique d'un O(n²).
- _Le programme est-il lent à calculer, ou seulement « cher » en jutsus ?_
il est juste cher en justus car chaque justu est très rapide a calculer, il y en as juste des milliers.

### Mission 3 — Shikamaru
- _Pourquoi la normalisation simplifie-t-elle autant le problème ?_
La normalisation simplifie le problème car elle rend les valeurs bien plus facile a utiliser.
- _Comment avez-vous choisi la taille de vos escouades ? Montrez vos mesures._
Le tri radix qui rend la chose plus simple car nous n'avons pas de taille d'escouade a définir.

### Mission 4 — Le contrôle du chakra
- _Que coûte en calcul la recherche de l'archive la moins chère à chaque tour ? Quelle complexité totale ?_
On regarde chaque archive, donc la recherche coûte O(n) par tour et O(n²) au total. Avec les déplacements dans les slices, le programme peut aller jusqu’à O(n³) dans le pire des cas.
- _Quelle amélioration vous a fait gagner le plus ? Laquelle vous a déçus ?_
Tester plusieurs départs nous a fait gagne plus : environ 5 171 à 5 117 jutsus en simulation. Par contre, envoyer trois archives dans B au départ, au lieu de deux n’a presque rien changé.

## 5. Le travail d'équipe
_Qui a fait quoi ? Comment avez-vous alterné pilote et copilote à chaque séance ?_
Nous avons travaillé en binome en faiseant en sorte que quand l'un travaille, l'autre l'aide mais nous ne sommes jamais 2 à travailler en même temps

## 6. Missions secrètes (facultatif)

_Pour chaque mission secrète tentée : ce qui fonctionne, vos mesures, ce que vous avez appris._
