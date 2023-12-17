package main

import (
	"bufio"
	"fmt"
	"os"
)

func File_to_string_table(filename string) []string {
	//readfile
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The file was opened successfully")
	//convert to string
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

type Pattern struct {
	pattern_number int        //starts at 0
	grid           [][]string //table of "." or "#"
	vertical       []int      //index of the column before reflection line [0:len(pattern)-1[
	horizontal     []int      //index of the row above reflection line [0:len(pattern)-1[
}

func Builds_patterns_from_file(lines []string) []Pattern {
	var output []Pattern
	var pattern_number int = 0
	var grid [][]string
	var vertical []int
	var horizontal []int
	for line_i, line := range lines {
		if line != "" {
			var grid_line []string
			runeSlice := []rune(line)
			for i := 0; i < len(runeSlice); i++ {
				char := string(runeSlice[i])
				grid_line = append(grid_line, char)
			}
			grid = append(grid, grid_line)

		}
		if line == "" || line_i == len(lines)-1 {
			patt := Pattern{pattern_number, grid, vertical, horizontal}
			output = append(output, patt)
			pattern_number += 1
			var new_grid [][]string
			grid = new_grid
		}
	}
	return output
}

func (patt *Pattern) Updates_hor_reflections() {
	for possible_refl, _ := range patt.grid {
		var is_refl_index bool = true
		var i int = possible_refl
		var i_mirrored = possible_refl + 1
		for i >= 0 && i < len(patt.grid) && i_mirrored < len(patt.grid) {
			var j int = 0
			for j < len(patt.grid[i]) && j < len(patt.grid[i_mirrored]) {
				if patt.grid[i][j] != patt.grid[i_mirrored][j] {
					is_refl_index = false
				}
				j += 1
			}
			i = i - 1
			i_mirrored = i_mirrored + 1
		}
		if possible_refl == len(patt.grid)-1 {
			is_refl_index = false
		}
		if len(patt.grid) == 0 {
			is_refl_index = false
		}
		if is_refl_index {
			patt.horizontal = append(patt.horizontal, possible_refl)
		}
	}
}

func (patt *Pattern) Updates_ver_reflections() {
	for possible_refl, _ := range patt.grid[0] {
		var is_refl_index bool = true
		for i := 0; i < len(patt.grid); i++ {
			var j int = possible_refl
			var j_mirrored = possible_refl + 1
			for j >= 0 && j < len(patt.grid[i]) && j_mirrored < len(patt.grid[i]) {
				if patt.grid[i][j] != patt.grid[i][j_mirrored] {
					is_refl_index = false
				}
				j = j - 1
				j_mirrored = j_mirrored + 1
			}
			if possible_refl == len(patt.grid[i])-1 {
				is_refl_index = false
			}
		}
		if len(patt.grid) == 0 {
			is_refl_index = false
		}
		if is_refl_index {
			patt.vertical = append(patt.vertical, possible_refl)
		}
	}
}

func Calculates_answer(patt_array []Pattern) int {
	var res int = 0
	for _, patt := range patt_array {
		patt.Updates_hor_reflections()
		patt.Updates_ver_reflections()
		for _, vert_mirror := range patt.vertical {
			var col_num int = vert_mirror + 1
			res += col_num
		}
		for _, hor_mirror := range patt.horizontal {
			var line_num int = hor_mirror + 1
			res += 100 * line_num
		}
	}
	return res
}

func main() {
	var lines []string = File_to_string_table("input.txt")
	var patterns []Pattern = Builds_patterns_from_file(lines)
	fmt.Println(Calculates_answer(patterns))
}
