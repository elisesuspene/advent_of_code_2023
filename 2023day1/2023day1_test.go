package main

import "testing"

func TestGets_digits(t *testing.T) {
	string_table := []string{"H1ell2o", "hi445", "by5e", "4go5odbye"}
	res_g_d := Gets_digits(string_table)
	expected := []float64{12, 45, 55, 45}
	var num_obtained float64
	var num_expected float64
	for i, word := range string_table {
		num_obtained = res_g_d[i]
		num_expected = expected[i]
		if num_expected != num_obtained {
			t.Errorf("for %s ; got %f", word, num_obtained)
		}
	}
}
