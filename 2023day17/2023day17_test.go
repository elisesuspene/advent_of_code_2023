package main

import (
	"testing"
)

func TestBuilds_matrix(t *testing.T) {
	lines := []string{"12", "34"}
	mat := Builds_matrix(lines)
	if mat[0][0] != 1 {
		t.Errorf("mat[0][0] is %d instead of %d", mat[0][0], 1)
	}
	if mat[1][0] != 3 {
		t.Errorf("mat[1][0] is %d instead of %d", mat[1][0], 3)
	}
	if mat[0][1] != 2 {
		t.Errorf("mat[0][1] is %d instead of %d", mat[0][1], 2)
	}
}

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

func TestCalculates_path(t *testing.T) {
	lines := []string{"12", "34"}
	mat := Builds_matrix(lines)
	var last_step Step = Calculates_path(mat)
	if last_step.i != 1 {
		t.Errorf("last_step.i is %d instead of %d", last_step.i, 1)
	}
	if last_step.col_dir != 0 {
		t.Errorf("last_step.col_dir is %d instead of %d", last_step.col_dir, 0)
	}
	if last_step.n_dir != 1 {
		t.Errorf("last_step.n_dir is %d instead of %d", last_step.n_dir, 1)
	}
	if last_step.heat_loss != 6 {
		t.Errorf("last_step.heat_loss is %d instead of %d", last_step.heat_loss, 6)
	}
	var middle_step []int = last_step.path_to_step[len(last_step.path_to_step)-2]
	if middle_step[0] != 0 {
		t.Errorf("middle_step.i is %d instead of %d", middle_step[0], 0)
	}
	if middle_step[1] != 1 {
		t.Errorf("middle_step.j is %d instead of %d", middle_step[1], 1)
	}
	lines1 := []string{"2413432311323", "3215453535623", "3255245654254", "3446585845452", "4546657867536", "1438598798454", "4457876987766", "3637877979653", "4654967986887", "4564679986453", "1224686865563", "2546548887735", "4322674655533"}
	mat = Builds_matrix(lines1)
	last_step = Calculates_path(mat)
	expected := []int{1, 4}
	if !Equal_slices(last_step.path_to_step[4], expected) {
		t.Errorf("last_step.path_to_step[4] is %d instead of %d", last_step.path_to_step[4], expected)
	}
	if last_step.heat_loss != 102 {
		t.Errorf("last_step.heat_loss is %d instead of %d", last_step.heat_loss, 102)
	}
}
