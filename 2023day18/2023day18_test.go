package main

import (
	"testing"
)

func TestParsing_instructions(t *testing.T) {
	lines := []string{"R 6 (#70c710)", "D 5 (#0dc571)", "L 2 (#5713f0)", "D 2 (#d2c081)", "R 2 (#59c680)", "D 2 (#411b91)", "L 5 (#8ceee2)", "U 2 (#caa173)", "L 1 (#1b58a2)", "U 2 (#caa171)", "R 2 (#7807d2)", "U 3 (#a77fa3)", "L 2 (#015232)", "U 2 (#7a21e3)"}
	var instructions []Instruction = Parsing_instructions(lines)
	if len(instructions) != len(lines) {
		t.Errorf("len(instructions) is %d instead of len(lines)", len(instructions))
	}
	if instructions[1].direction != "D" {
		t.Errorf("instructions[1].direction is %s instead of D", instructions[1].direction)
	}
	if instructions[1].tile_num != 5 {
		t.Errorf("instructions[1].tile_num is %d instead of 5", instructions[1].tile_num)
	}
	if instructions[1].color_code != "#0dc571" {
		t.Errorf("instructions[1].color_code is %s instead of #0dc571", instructions[1].color_code)
	}
}

func TestResizing_tile_array(t *testing.T) {
	var empty_array [][]Tile
	var empty_line []Tile
	empty_line = append(empty_line, Tile{0, 0, true, false, ""})
	empty_array = append(empty_array, empty_line)
	output, _, _, _, _ := Resizing_tile_array(empty_array, 1, 1, 1, 1)
	output, _, _, _, _ = Resizing_tile_array(output, 1, 1, 1, 2)
	if len(output) != 2 {
		t.Errorf("len(output) is %d instead of 2", len(output))
	}
	if len(output[0]) != 3 {
		t.Errorf("len(output[0]) is %d instead of 3", len(output[0]))
	}
	output, _, _, _, _ = Resizing_tile_array(empty_array, 0, -1, 0, -1)
	if len(output) != 1 {
		t.Errorf("len(output) is %d instead of 1", len(output))
	}
	if len(output[0]) != 2 {
		t.Errorf("len(output[0]) is %d instead of 2", len(output[0]))
	}
}

func TestBuilds_frontier(t *testing.T) {
	lines := []string{"R 6 (#70c710)", "D 5 (#0dc571)", "L 2 (#5713f0)", "D 2 (#d2c081)", "R 2 (#59c680)", "D 2 (#411b91)", "L 5 (#8ceee2)", "U 2 (#caa173)", "L 1 (#1b58a2)", "U 2 (#caa171)", "R 2 (#7807d2)", "U 3 (#a77fa3)", "L 2 (#015232)", "U 2 (#7a21e3)"}
	var instructions []Instruction = Parsing_instructions(lines)
	var output [][]Tile = Builds_frontier(instructions)
	if len(output) != 10 {
		t.Errorf("len(output) is %d instead of 10", len(output))
	}
	if len(output[0]) != 7 {
		t.Errorf("len(output[0]) is %d instead of 7", len(output[0]))
	}
	if !output[0][2].is_on_edge {
		t.Errorf("output[0][2] is not on frontier when it should be")
	}
	if output[3][1].is_on_edge {
		t.Errorf("output[3][1] is on frontier when it should not be")
	}
	if output[1][3].is_on_edge {
		t.Errorf("output[1][3] is on frontier when it should not be")
	}
}
