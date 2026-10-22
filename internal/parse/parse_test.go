// =====================================================================
//  internal/parse/parse_test.go — TESTER LES PIÈGES
//  Mission 0 (Académie)
// =====================================================================
//
//  Lancer les tests :   go test ./internal/parse/
//
//  À FAIRE : un test en table qui couvre AU MINIMUM ces cas.
//
//   Entrée                        Résultat attendu
//   ---------------------------   ----------------------------
//   "4" "kunai" "2"               erreur
//   "1" "5" "1"                   erreur (doublon)
//   "2147483648"                  erreur (trop grand)
//   "-2147483649"                 erreur (trop petit)
//   "-"  "3"                      erreur
//   "4-2"                         erreur
//   "" (argument vide)            erreur
//   "42"                          [42]
//   "-5" "0" "-12"                [-5 0 -12]
//   "2147483647" "-2147483648"    [2147483647 -2147483648]
//   (aucun argument)              liste vide, pas d'erreur
//
//  Ajoutez tous les pièges que vous imaginez : l'Akatsuki est sournoise,
//  les examinateurs aussi.
//
// =====================================================================

package parse

import "testing"

func TestArgs(t *testing.T) {
	bons := [][]string{
		{"-5", "0", "-12"},
		{"2147483647", "-2147483648"},
		{"42"},
		{},
	}
	for _, c := range bons {
		if _, err := Args(c); err != nil {
			t.Errorf("%v devrait être accepté : %v", c, err)
		}
	}

	mauvais := [][]string{
		{"4", "kunai", "2"},
		{"-"},
		{"4-2"},
		{"2147483648"},
		{"-2147483649"},
		{"1", "5", "1"},
		{""},
	}
	for _, c := range mauvais {
		if _, err := Args(c); err == nil {
			t.Errorf("%v devrait être refusé", c)
		}
	}

	if _, err := ArgsAvecDoublons([]string{"1", "5", "1"}); err != nil {
		t.Errorf("doublons devraient être permis : %v", err)
	}
}
