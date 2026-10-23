#!/usr/bin/env bash
# =====================================================================
#  bench.sh — LE SCRIPT DE MESURE ET DE TEST DE VOTRE ÉQUIPE
# =====================================================================
#
#  Utilisation :
#      ./bench.sh 100 20      → 20 tirages de 100 archives (mesure)
#      ./bench.sh 500 10      → 10 tirages de 500 archives (mesure)
#      ./bench.sh test        → teste les programmes ninja_swap et juge
#
#  Mesure : pour chaque tirage, vérifie avec le juge puis compte les
#  jutsus ; affiche à la fin le minimum, le maximum, la MOYENNE et le
#  temps de calcul total.
#
#  Test : vérifie le comportement des deux programmes (messages d'erreur,
#  cas limites, verdicts du juge, petites piles) — ce que les tests Go
#  de internal/ ne peuvent pas voir, car ils ne lancent pas les main.go.
#
#  Sous Windows : lancez ce script dans Git Bash ou WSL, pas dans PowerShell.
#
# =====================================================================

cd "$(dirname "$0")" || exit 1

compiler() {
    go build -o ninja_swap ./cmd/ninja_swap || exit 1
    go build -o juge ./cmd/juge || exit 1
}

# ---------------------------------------------------------------------
#  MODE MESURE
# ---------------------------------------------------------------------
mesurer() {
    local N=$1 COUNT=$2
    compiler

    local total=0 min=-1 max=0 debut fin
    debut=$(date +%s.%N)
    for ((i = 1; i <= COUNT; i++)); do
        local ARG
        ARG=$(shuf -i 1-10000 -n "$N" | tr '\n' ' ')

        local verdict
        verdict=$(./ninja_swap $ARG | ./juge $ARG)
        if [ "$verdict" != "OK" ]; then
            echo "ÉCHEC (verdict : $verdict) sur le tirage :"
            echo "$ARG"
            exit 1
        fi

        local n
        n=$(./ninja_swap $ARG | wc -l)
        total=$((total + n))
        if [ "$min" -eq -1 ] || [ "$n" -lt "$min" ]; then min=$n; fi
        if [ "$n" -gt "$max" ]; then max=$n; fi
    done
    fin=$(date +%s.%N)

    echo "$COUNT tirages de $N archives : min=$min  max=$max  moyenne=$(awk "BEGIN{printf \"%.2f\", $total/$COUNT}")"
    echo "temps total (avec vérification) : $(awk "BEGIN{printf \"%.2f\", $fin-$debut}") s"
}

# ---------------------------------------------------------------------
#  MODE TEST
# ---------------------------------------------------------------------
REUSSIS=0
ECHECS=0
ERR=$(mktemp)
trap 'rm -f "$ERR"' EXIT

ok() { REUSSIS=$((REUSSIS + 1)); }
ko() { ECHECS=$((ECHECS + 1)); echo "ÉCHEC : $1"; }

# attendre NOM ATTENDU OBTENU
attendre() {
    if [ "$2" == "$3" ]; then ok; else ko "$1 : attendu '$2', obtenu '$3'"; fi
}

# erreur_attendue NOM PROGRAMME [arguments...] [<<< entrée]
# Vérifie : stdout vide, stderr = "Error", code de sortie 1.
erreur_attendue() {
    local nom=$1 prog=$2
    shift 2
    local sortie code
    sortie=$("$prog" "$@" 2>"$ERR" </dev/null)
    code=$?
    attendre "$nom (stdout vide)" "" "$sortie"
    attendre "$nom (stderr)" "Error" "$(cat "$ERR")"
    attendre "$nom (code de sortie)" "1" "$code"
}

tester() {
    compiler
    echo "Tests de ninja_swap et du juge..."

    # ----- ninja_swap : arguments invalides -----
    erreur_attendue "ninja_swap non numérique"   ./ninja_swap 4 kunai 2
    erreur_attendue "ninja_swap signe seul"      ./ninja_swap -
    erreur_attendue "ninja_swap signe mal placé" ./ninja_swap 4-2
    erreur_attendue "ninja_swap > int32"         ./ninja_swap 2147483648
    erreur_attendue "ninja_swap < int32"         ./ninja_swap -2147483649
    erreur_attendue "ninja_swap doublon"         ./ninja_swap 1 5 1
    erreur_attendue "ninja_swap argument vide"   ./ninja_swap ""

    # ----- ninja_swap : rien à faire -----
    local sortie code
    sortie=$(./ninja_swap 2>"$ERR"); code=$?
    attendre "ninja_swap sans argument (stdout vide)" "" "$sortie"
    attendre "ninja_swap sans argument (stderr vide)" "" "$(cat "$ERR")"
    attendre "ninja_swap sans argument (code 0)" "0" "$code"

    attendre "ninja_swap une archive"      "" "$(./ninja_swap 42)"
    attendre "ninja_swap déjà trié"        "" "$(./ninja_swap 1 2 3 4 5)"
    attendre "ninja_swap bornes triées"    "" "$(./ninja_swap -2147483648 0 2147483647)"

    # ----- ninja_swap : valeurs limites acceptées et triées -----
    local ARG
    for ARG in "-5 0 -12" "2147483647 -2147483648" "2 1" "-1 -2 -3 -4 -5 -6 -7"; do
        attendre "ninja_swap trie « $ARG »" "OK" "$(./ninja_swap $ARG | ./juge $ARG)"
    done

    # ----- ninja_swap : petites piles (3, 4, 5 archives), maximums du sujet -----
    local p n
    for p in "1 3 2" "2 1 3" "2 3 1" "3 1 2" "3 2 1"; do
        attendre "ninja_swap trie « $p »" "OK" "$(./ninja_swap $p | ./juge $p)"
        n=$(./ninja_swap $p | wc -l)
        if [ "$n" -le 2 ]; then ok; else ko "ninja_swap « $p » : $n jutsus (maximum 2)"; fi
    done
    for ((i = 0; i < 30; i++)); do
        p=$(shuf -i 1-50 -n 5 | tr '\n' ' ')
        attendre "ninja_swap trie « $p »" "OK" "$(./ninja_swap $p | ./juge $p)"
        n=$(./ninja_swap $p | wc -l)
        if [ "$n" -le 12 ]; then ok; else ko "ninja_swap « $p » : $n jutsus (objectif 12 maximum)"; fi
    done

    # ----- ninja_swap : tirages aléatoires de toutes tailles -----
    local taille
    for taille in 2 4 6 7 10 25 100 300; do
        for ((i = 0; i < 5; i++)); do
            ARG=$(shuf -i 1-10000 -n "$taille" | tr '\n' ' ')
            attendre "ninja_swap trie $taille archives" "OK" "$(./ninja_swap $ARG | ./juge $ARG)"
        done
    done

    # ----- ninja_swap : sortie composée uniquement de jutsus valides -----
    ARG=$(shuf -i 1-10000 -n 50 | tr '\n' ' ')
    local invalides
    invalides=$(./ninja_swap $ARG | grep -vxE 'sa|sb|ss|pa|pb|ra|rb|rr|rra|rrb|rrr')
    attendre "ninja_swap n'écrit que des jutsus valides" "" "$invalides"

    # ----- juge : verdicts -----
    attendre "juge OK (sa)"                 "OK" "$(printf 'sa\n' | ./juge 2 1)"
    attendre "juge OK (déjà trié, rien lu)" "OK" "$(./juge 1 2 3 </dev/null)"
    attendre "juge OK (exemple du sujet)"   "OK" "$(printf 'pb\nsa\npa\nsa\n' | ./juge 2 3 1 4)"
    attendre "juge OK (une archive)"        "OK" "$(./juge 42 </dev/null)"
    attendre "juge KO (rien lu, non trié)"  "KO" "$(./juge 2 1 </dev/null)"
    attendre "juge KO (suite insuffisante)" "KO" "$(printf 'ra\n' | ./juge 3 2 1)"
    attendre "juge KO (B non vide)"         "KO" "$(printf 'pb\n' | ./juge 1 2)"

    # ----- juge : erreurs -----
    erreur_attendue "juge argument invalide" ./juge 4 kunai 2
    erreur_attendue "juge doublon"           ./juge 1 5 1
    sortie=$(printf 'sa\nxx\n' | ./juge 2 1 2>"$ERR"); code=$?
    attendre "juge jutsu inconnu (stdout vide)" "" "$sortie"
    attendre "juge jutsu inconnu (stderr)" "Error" "$(cat "$ERR")"
    attendre "juge jutsu inconnu (code 1)" "1" "$code"
    sortie=$(printf 'SA\n' | ./juge 2 1 2>"$ERR"); code=$?
    attendre "juge jutsu en majuscules refusé" "1" "$code"
    sortie=$(printf 'sa\n\n' | ./juge 2 1 2>"$ERR"); code=$?
    attendre "juge ligne vide refusée" "1" "$code"

    echo
    echo "$REUSSIS vérifications réussies, $ECHECS échec(s)."
    [ "$ECHECS" -eq 0 ]
}

# ---------------------------------------------------------------------
#  Point d'entrée
# ---------------------------------------------------------------------
if [ "$1" == "test" ]; then
    tester
    exit $?
fi

if [ $# -ne 2 ] || ! [[ "$1" =~ ^[0-9]+$ ]] || ! [[ "$2" =~ ^[0-9]+$ ]] || [ "$1" -lt 1 ] || [ "$2" -lt 1 ]; then
    echo "Utilisation : $0 N TIRAGES   (ex. $0 100 20)"
    echo "              $0 test"
    exit 1
fi

mesurer "$1" "$2"
