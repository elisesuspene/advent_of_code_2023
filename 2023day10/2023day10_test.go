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
	grid := []string{".....", ".S-7.", ".|.|.", ".L-J.", "....."}
	var tiles [][]string = Build_tiles_from_file(grid)
	expected := [][]string{{".", ".", ".", ".", "."}, {".", "S", "-", "7", "."}, {".", "|", ".", "|", "."}, {".", "L", "-", "J", "."}, {".", ".", ".", ".", "."}}
	if !Equal_slices(tiles, expected) {
		t.Errorf("Build_tiles_from_file did not return the expected table of table of strings.\n Instead of %+v, found :%+v", expected, tiles)
	}
}
