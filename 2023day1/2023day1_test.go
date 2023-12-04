package main

import (
	"testing"
)

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

func TestStartOfSpelledOutDigit(t *testing.T) {
	var input string = "one"
	IsStart, digit_found := StartOfSpelledOutDigit(input, 0)
	if IsStart != true || digit_found != "one" {
		t.Errorf("for string %s and index 0 ; got %t and %s", input, IsStart, digit_found)
	}
	input = "1one2"
	IsStart, digit_found = StartOfSpelledOutDigit(input, 1)
	if IsStart != true || digit_found != "one" {
		t.Errorf("for string %s and index 1 ; got %t and %s", input, IsStart, digit_found)
	}
}

func TestEndOfSpelledOutDigit(t *testing.T) {
	var input string = "one"
	IsEnd, digit_found := EndOfSpelledOutDigit(input, 2)
	if IsEnd != true || digit_found != "one" {
		t.Errorf("for string %s and index 2 ; got %t and %s", input, IsEnd, digit_found)
	}
	input = "1one2"
	IsEnd, digit_found = EndOfSpelledOutDigit(input, 3)
	if IsEnd != true || digit_found != "one" {
		t.Errorf("for string %s and index 3 ; got %t and %s", input, IsEnd, digit_found)
	}
}

func TestGets_digits2(t *testing.T) {
	input := []string{"1two", "3five4"}
	output := Gets_digits2(input)
	expected := []float64{12, 34}
	if !Equal_slices(expected, output) {
		t.Errorf("for string table %v ; got %v", input, output)
	}
}
