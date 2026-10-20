#!/usr/bin/env bash
# =====================================================================
#  bench.sh — LE SCRIPT DE MESURE DE VOTRE ÉQUIPE
#  À utiliser dès la mission 2, puis à chaque nouvelle version
# =====================================================================
#
#  Rôle : mesurer votre ninja_swap sur des tirages aléatoires et afficher
#  les moyennes à recopier dans le journal d'optimisation du README.
#
#  Utilisation visée :
#      ./bench.sh 100 20      → 20 tirages de 100 archives
#      ./bench.sh 500 10      → 10 tirages de 500 archives
#
#  À FAIRE :
#
#   1. Compiler ninja_swap et le juge avant de mesurer
#      (go build -o ninja_swap ./cmd/ninja_swap, idem pour le juge).
#
#   2. Pour chaque tirage :
#      - générer N nombres distincts au hasard :
#            ARG=$(shuf -i 1-10000 -n "$N" | tr '\n' ' ')
#      - vérifier avec le juge :  ./ninja_swap $ARG | ./juge $ARG
#        Si le résultat n'est pas OK → afficher le tirage fautif et s'arrêter.
#      - compter les jutsus :      ./ninja_swap $ARG | wc -l
#
#   3. Afficher à la fin : le minimum, le maximum et la MOYENNE.
#
#  BONUS : mesurer aussi le temps de calcul (commande time).
#
#  Sous Windows : lancez ce script dans Git Bash ou WSL, pas dans PowerShell.
#
# =====================================================================

# TODO : à vous de jouer, genin.
