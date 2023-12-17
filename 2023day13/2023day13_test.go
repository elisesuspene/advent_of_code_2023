package main

import (
	"testing"
)

func Equal_slices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, _ := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func TestUpdates_ver_reflections(t *testing.T) {
	grid := [][]string{{"#", ".", "#", "#", ".", ".", "#", "#", "."}, {".", ".", "#", ".", "#", "#", ".", "#", "."}, {"#", "#", ".", ".", ".", ".", ".", ".", "#"}, {"#", "#", ".", ".", ".", ".", ".", ".", "#"}, {".", ".", "#", ".", "#", "#", ".", "#", "."}, {".", ".", "#", "#", ".", ".", "#", "#", "."}, {"#", ".", "#", ".", "#", "#", ".", "#", "."}}
	var vertical []int
	var horizontal []int
	patt := Pattern{0, grid, vertical, horizontal}
	patt.Updates_hor_reflections()
	patt.Updates_ver_reflections()
	expected_ver := []int{4}
	var expected_hor []int
	if !Equal_slices(patt.vertical, expected_ver) {
		t.Errorf("patt.vertical is %v instead of %v", patt.vertical, expected_ver)
	}
	if !Equal_slices(patt.horizontal, expected_hor) {
		t.Errorf("patt.horizontal is %v instead of %v", patt.horizontal, expected_hor)
	}
	var empty_file []string
	empty_patt_array := Builds_patterns_from_file(empty_file)
	var expected_hor2 []int
	var expected_ver2 []int
	for _, patt := range empty_patt_array {
		patt.Updates_hor_reflections()
		patt.Updates_ver_reflections()
		if !Equal_slices(patt.horizontal, expected_hor2) {
			t.Errorf("patt.horizontal is %v instead of %v", patt.horizontal, expected_hor2)
		}
		if !Equal_slices(patt.vertical, expected_ver2) {
			t.Errorf("patt.horizontal is %v instead of %v", patt.vertical, expected_ver2)
		}
	}
}

func TestBuilds_patterns_from_file(t *testing.T) {
	file_lines := []string{"#.##..##.", "..#.##.#.", "##......#", "##......#", "..#.##.#.", "..##..##.", "#.#.##.#.", "", "#...##..#", "#....#..#", "..##..###", "#####.##.", "#####.##.", "..##..###", "#....#..#"}
	var patt_array []Pattern = Builds_patterns_from_file(file_lines)
	var empty []int
	for i, patt := range patt_array {
		if patt.pattern_number != i {
			t.Errorf("patt_number is %d instead of %d", patt.pattern_number, i)
		}
		if !Equal_slices(patt.horizontal, empty) {
			t.Errorf("patt.horizontal is %v instead of %v", patt.horizontal, empty)
		}
		if !Equal_slices(patt.vertical, empty) {
			t.Errorf("patt.vertical is %v instead of %v", patt.vertical, empty)
		}
	}
}

func TestCalculates_answer(t *testing.T) {
	file_lines := []string{"#.##..##.", "..#.##.#.", "##......#", "##......#", "..#.##.#.", "..##..##.", "#.#.##.#.", "", "#...##..#", "#....#..#", "..##..###", "#####.##.", "#####.##.", "..##..###", "#....#..#"}
	var patt_array []Pattern = Builds_patterns_from_file(file_lines)
	var result int = Calculates_answer(patt_array)
	var expected int = 405
	if result != expected {
		t.Errorf("Calculates_answer gave %d instead of %d", result, expected)
	}
}
