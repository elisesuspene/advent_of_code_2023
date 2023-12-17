package main

import (
	"testing"
)

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

func TestNew_reflections(t *testing.T) {
	file_lines := []string{"#.##..##.", "..#.##.#.", "##......#", "##......#", "..#.##.#.", "..##..##.", "#.#.##.#.", "", "#...##..#", "#....#..#", "..##..###", "#####.##.", "#####.##.", "..##..###", "#....#..#"}
	var patt_array []Pattern = Builds_patterns_from_file(file_lines)
	patt_array[0].Updates_hor_reflections()
	patt_array[0].Updates_ver_reflections()
	patt_array[1].Updates_hor_reflections()
	patt_array[1].Updates_ver_reflections()
	horizontal0, vertical0 := patt_array[0].New_reflections()
	horizontal1, vertical1 := patt_array[1].New_reflections()
	exp_hor0 := []int{2}
	var exp_ver0 []int
	exp_hor1 := []int{0}
	var exp_ver1 []int
	if !Equal_slices(horizontal0, exp_hor0) {
		t.Errorf("New_reflections for horizontals of the first pattern gave %v instead of %v", horizontal0, exp_hor0)
	}
	if !Equal_slices(vertical0, exp_ver0) {
		t.Errorf("New_reflections for verticals of the first pattern gave %v instead of %v", vertical0, exp_ver0)
	}
	if !Equal_slices(horizontal1, exp_hor1) {
		t.Errorf("New_reflections for horizontals of the second pattern gave %v instead of %v", horizontal1, exp_hor1)
	}
	if !Equal_slices(vertical1, exp_ver1) {
		t.Errorf("New_reflections for verticals of the second pattern gave %v instead of %v", vertical1, exp_ver1)
	}
}
