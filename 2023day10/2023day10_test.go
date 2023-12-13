package main

import (
	"testing"
)

func Equal_slices(slice1 [][]string, slice2 [][]string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, _ := range slice1 {
		for j, _ := range slice1[0] {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}
	}
	return true
}

func TestBuilds_tiles_from_file(t *testing.T) {
	var grid []string
	var st string = "....."
	grid = append(grid, st)
	st = ".S-7."
	grid = append(grid, st)
	st = ".|.|."
	grid = append(grid, st)
	st = ".L-J."
	grid = append(grid, st)
	st = "....."
	grid = append(grid, st)
	var tiles [][]string = Build_tiles_from_file(grid)
	var expected [][]string
	line := []string{".", ".", ".", ".", "."}
	expected = append(expected, line)
	line[1] = "S"
	line[2] = "-"
	line[3] = "7"
	expected = append(expected, line)
	line[1] = "|"
	line[2] = "."
	line[3] = "|"
	expected = append(expected, line)
	line[1] = "L"
	line[2] = "-"
	line[3] = "J"
	expected = append(expected, line)
	line[1] = "."
	line[2] = "."
	line[3] = "."
	expected = append(expected, line)
	if !Equal_slices(tiles, expected) {
		t.Errorf("Build_tiles_from_file did not return the expected table of table of strings")
	}
}
