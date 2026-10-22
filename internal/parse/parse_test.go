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

// Cas limites supplémentaires.
func TestArgsLimites(t *testing.T) {
	mauvais := [][]string{
		{" 5"},
		{"5 "},
		{"1.5"},
		{"0x10"},
		{"--5"},
		{"5-"},
		{"1", "01"},  // 01 vaut 1 : doublon
		{"0", "-0"},  // -0 vaut 0 : doublon
		{"1", "2", "3", "2"},
	}
	for _, c := range mauvais {
		if _, err := Args(c); err == nil {
			t.Errorf("%q devrait être refusé", c)
		}
	}
}

// Les valeurs renvoyées sont dans l'ordre des arguments.
func TestArgsValeurs(t *testing.T) {
	got, err := Args([]string{"3", "-1", "2147483647", "-2147483648", "0"})
	if err != nil {
		t.Fatal(err)
	}
	want := []int{3, -1, 2147483647, -2147483648, 0}
	if len(got) != len(want) {
		t.Fatalf("obtenu %v, attendu %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("obtenu %v, attendu %v", got, want)
		}
	}
}

// Avec doublons permis, les valeurs invalides restent refusées.
func TestArgsAvecDoublonsInvalides(t *testing.T) {
	got, err := ArgsAvecDoublons([]string{"1", "5", "1"})
	if err != nil || len(got) != 3 {
		t.Errorf("1 5 1 : obtenu %v, %v", got, err)
	}
	for _, c := range [][]string{{"kunai"}, {"-"}, {"2147483648"}, {""}} {
		if _, err := ArgsAvecDoublons(c); err == nil {
			t.Errorf("%q devrait être refusé même avec doublons", c)
		}
	}
}
