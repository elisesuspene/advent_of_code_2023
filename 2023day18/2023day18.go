package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func File_to_string_table(filename string) []string {
	//readfile
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The file was opened successfully")
	//convert to []string
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

type Instruction struct {
	i          int //starts at 0
	direction  string
	tile_num   int
	color_code string
}

func Parsing_instructions(lines []string) []Instruction {
	var instructions []Instruction
	for i, line := range lines {
		var split_line []string = strings.Split(line, " ")
		var direction string = split_line[0]
		tile_num, _ := strconv.Atoi(split_line[1])
		var color_code string = split_line[2][1 : len(split_line[2])-1]
		instruction := Instruction{i, direction, tile_num, color_code}
		instructions = append(instructions, instruction)
	}
	return instructions
}

type Tile struct {
	i              int //counting from left to right and from up to down
	j              int
	is_on_edge     bool
	is_on_interior bool
	color          string
}

func Resizing_tile_array(tile_array [][]Tile, frontier_start_i int, frontier_start_j int, i int, j int) ([][]Tile, int, int, int, int) {
	var new_frontier_start_i int = frontier_start_i
	var new_frontier_start_j int = frontier_start_j
	var new_i int = i
	var new_j int = j
	if i >= len(tile_array) {
		var new_tile_array [][]Tile = tile_array
		for k := len(tile_array); k <= i; k++ {
			var blank_line []Tile
			for j := 0; j < len(tile_array[0]); j++ {
				blank_line = append(blank_line, Tile{k, j, false, false, ""})
			}
			new_tile_array = append(new_tile_array, blank_line)
		}
		tile_array = new_tile_array
	}
	if j >= len(tile_array[0]) {
		for i := 0; i < len(tile_array); i++ {
			var new_line []Tile = tile_array[i]
			for k := len(tile_array[i]); k <= j; k++ {
				new_line = append(new_line, Tile{i, k, false, false, ""})
			}
			tile_array[i] = new_line
		}
	}
	if i < 0 {
		var new_tile_array [][]Tile
		//making space for lines on top
		for k := 0; k < -i; k++ {
			var blank_line []Tile
			for j := 0; j < len(tile_array[0]); j++ {
				blank_line = append(blank_line, Tile{k, j, false, false, ""})
			}
			new_tile_array = append(new_tile_array, blank_line)
		}
		for q := 0; q < len(tile_array); q++ {
			var new_line []Tile
			for _, old_tile := range tile_array[q] {
				old_tile.i = old_tile.i - i
				new_line = append(new_line, old_tile)
			}
			new_tile_array = append(new_tile_array, new_line)
		}
		tile_array = new_tile_array
		new_i = 0
		new_frontier_start_i = -i - 1
	}
	if j < 0 {
		var new_tile_array [][]Tile
		// making space for columns on the left
		for i := 0; i < len(tile_array); i++ {
			var new_line []Tile
			for k := 0; k < -j; k++ {
				new_line = append(new_line, Tile{i, k, false, false, ""})
			}
			for _, old_tile := range tile_array[0] {
				old_tile.j = old_tile.j - j
				new_line = append(new_line, old_tile)
			}
			new_tile_array = append(new_tile_array, new_line)
		}
		tile_array = new_tile_array
		new_j = 0
		new_frontier_start_j = -j - 1
	}
	return tile_array, new_frontier_start_i, new_frontier_start_j, new_i, new_j
}

func Builds_frontier(instructions []Instruction) [][]Tile {
	var tile_array [][]Tile
	var tile_line []Tile
	tile := Tile{0, 0, true, false, ""}
	tile_line = append(tile_line, tile)
	tile_array = append(tile_array, tile_line)
	var i int = 0
	var j int = 0
	var frontier_start_i int = i
	var frontier_start_j int = j
	for _, instruction := range instructions {
		frontier_start_i = i
		frontier_start_j = j
		if instruction.direction == "U" {
			frontier_start_i = i - 1
			i = i - instruction.tile_num
		}
		if instruction.direction == "D" {
			frontier_start_i = i + 1
			i = i + instruction.tile_num
		}
		if instruction.direction == "R" {
			frontier_start_j = j + 1
			j = j + instruction.tile_num
		}
		if instruction.direction == "L" {
			frontier_start_j = j - 1
			j = j - instruction.tile_num
		}
		if i >= len(tile_array) || j >= len(tile_array[0]) || i < 0 || j < 0 {
			tile_array, frontier_start_i, frontier_start_j, i, j = Resizing_tile_array(tile_array, frontier_start_i, frontier_start_j, i, j)
		}
		for k := min(frontier_start_i, i); k <= max(frontier_start_i, i); k++ {
			for q := min(frontier_start_j, j); q <= max(frontier_start_j, j); q++ {
				tile_array[k][q].is_on_edge = true
				tile_array[k][q].color = instruction.color_code
			}
		}
	}
	return tile_array
}

func Flood_fill(tile_array [][]Tile) ([][]Tile, []Tile) {
	var interior []Tile
	var i int
	var j int
	//finding a start tile
	var found_a_start bool = false
	for k := 0; k < len(tile_array); k++ {
		for q := 0; q < len(tile_array[0]); q++ {
			if !found_a_start {
				if tile_array[k][q].is_on_edge {
					i = k
					j = q
					found_a_start = true
				}
			}
		}
	}
	var neighbors []Tile
	var possible_neighbors []Tile
	if i+1 < len(tile_array) {
		possible_neighbors = append(possible_neighbors, tile_array[i+1][j])
	}
	if i-1 >= 0 {
		possible_neighbors = append(possible_neighbors, tile_array[i-1][j])
	}
	if j+1 < len(tile_array[0]) {
		possible_neighbors = append(possible_neighbors, tile_array[i][j+1])
	}
	if j-1 >= 0 {
		possible_neighbors = append(possible_neighbors, tile_array[i][j-1])
	}
	for _, poss_neighbor := range possible_neighbors {
		if !poss_neighbor.is_on_edge {
			neighbors = append(neighbors, poss_neighbor)
		}
	}
	for len(neighbors) != 0 {
		if !neighbors[0].is_on_edge {
			neighbors[0].is_on_interior = true
			interior = append(interior, neighbors[0])
		}
		i = neighbors[0].i
		j = neighbors[0].j
		var new_possible_neighbors []Tile
		if i+1 < len(tile_array) {
			new_possible_neighbors = append(new_possible_neighbors, tile_array[i+1][j])
		}
		if i-1 >= 0 {
			new_possible_neighbors = append(new_possible_neighbors, tile_array[i-1][j])
		}
		if j+1 < len(tile_array[0]) {
			new_possible_neighbors = append(new_possible_neighbors, tile_array[i][j+1])
		}
		if j-1 >= 0 {
			new_possible_neighbors = append(new_possible_neighbors, tile_array[i][j-1])
		}
		for _, poss_neighbor := range new_possible_neighbors {
			if !poss_neighbor.is_on_edge {
				neighbors = append(neighbors, poss_neighbor)
			}
		}
	}
	return tile_array, interior
}

func Calculates_answer(tile_array [][]Tile, interior []Tile) int {
	var output int
	output += len(interior)
	for i := 0; i < len(tile_array); i++ {
		for j := 0; j < len(tile_array[0]); j++ {
			if tile_array[i][j].is_on_edge {
				output += 1
			}
		}
	}
	return output
}

func main() {
	var lines []string = File_to_string_table("input.txt")
	var instructions []Instruction = Parsing_instructions(lines)
	var tile_array [][]Tile = Builds_frontier(instructions)
	_, interior := Flood_fill(tile_array)
	var output int = Calculates_answer(tile_array, interior)
	fmt.Println(output)
}
