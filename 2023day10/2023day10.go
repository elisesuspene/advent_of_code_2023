package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func Build_tiles_from_file(lines []string) [][]string {
	var tiles [][]string
	for _, line := range lines {
		var tile_line []string
		runeSlice := []rune(line)
		for i := 0; i < len(runeSlice); i++ {
			var tile string = string(runeSlice[i])
			tile_line = append(tile_line, tile)
		}
		tiles = append(tiles, tile_line)
	}
	return tiles
}

func Next_tile(i int, j int, come_from string, pipe_shape string) (int, int, string, bool) {
	var next_i int
	var next_j int
	var next_come_from string = ""
	var succeed bool = false
	if come_from == "W" && pipe_shape == "-" {
		next_i = i + 1
		next_j = j
		next_come_from = "W"
		succeed = true
	}
	if come_from == "E" && pipe_shape == "-" {
		next_i = i - 1
		next_j = j
		next_come_from = "E"
		succeed = true
	}
	if come_from == "N" && pipe_shape == "|" {
		next_i = i
		next_j = j + 1
		next_come_from = "N"
		succeed = true
	}
	if come_from == "S" && pipe_shape == "|" {
		next_i = i
		next_j = j - 1
		next_come_from = "S"
		succeed = true
	}
	if come_from == "N" && pipe_shape == "L" {
		next_i = i + 1
		next_j = j + 1
		next_come_from = "W"
		succeed = true
	}
	if come_from == "E" && pipe_shape == "L" {
		next_i = i - 1
		next_j = j - 1
		next_come_from = "S"
		succeed = true
	}
	if come_from == "E" && pipe_shape == "F" {
		next_i = i - 1
		next_j = j + 1
		next_come_from = "N"
		succeed = true
	}
	if come_from == "S" && pipe_shape == "F" {
		next_i = i + 1
		next_j = j - 1
		next_come_from = "W"
		succeed = true
	}
	if come_from == "W" && pipe_shape == "7" {
		next_i = i + 1
		next_j = j + 1
		next_come_from = "N"
		succeed = true
	}
	if come_from == "S" && pipe_shape == "7" {
		next_i = i - 1
		next_j = j - 1
		next_come_from = "E"
		succeed = true
	}
	if come_from == "W" && pipe_shape == "J" {
		next_i = i + 1
		next_j = j - 1
		next_come_from = "S"
		succeed = true
	}
	if come_from == "N" && pipe_shape == "J" {
		next_i = i - 1
		next_j = j + 1
		next_come_from = "E"
		succeed = true
	}
	return next_i, next_j, next_come_from, succeed
}

func Extracts_main_loop(tiles [][]string) ([]int, []int, []string) {
	var i_indexes []int
	var j_indexes []int
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == "S" {
				i_indexes = append(i_indexes, i)
				j_indexes = append(j_indexes, j)
			}
		}
	}
	var distances []string
	distances = append(distances, strconv.Itoa(0))
	//trying to loop starting in WE direction
	var i int = i_indexes[len(i_indexes)-1]
	var j int = j_indexes[len(j_indexes)-1]
	shape_list := []string{"-", "|", "7", "J", "L", "F"}
	come_from_list := []string{"W", "E", "N", "S"}
	for _, start_come_from := range come_from_list {
		var come_from string = start_come_from
		for _, start_pipe_shape := range shape_list {
			var pipe_shape string = start_pipe_shape
			for 0 <= i && i < len(tiles) && 0 <= j && j < len(tiles[0]) {
				var succeed bool
				i, j, come_from, succeed = Next_tile(i, j, come_from, pipe_shape)
				if succeed {
					i_indexes = append(i_indexes, i)
					j_indexes = append(j_indexes, j)
					new_d_int, _ := strconv.Atoi(distances[len(distances)-1])
					new_d_int += 1
					var new_distance string = strconv.Itoa(new_d_int)
					distances = append(distances, new_distance)
				}
				if tiles[i][j] == "S" {
					return i_indexes, j_indexes, distances
				}
				i = i_indexes[len(i_indexes)-1]
				j = j_indexes[len(j_indexes)-1]
				pipe_shape = tiles[i][j]

			}
		}
	}
	return i_indexes, j_indexes, distances

}

func Find_farthest_distance(distances_string []string) int {
	var distances_int []int
	for _, distance_string := range distances_string {
		d, _ := strconv.Atoi(distance_string)
		distances_int = append(distances_int, d)
	}
	var i int = len(distances_string) / 2
	return distances_int[i]
}

func main() {
	var lines []string = File_to_string_table("input.txt")
	var tiles [][]string = Build_tiles_from_file(lines)
	_, _, distances := Extracts_main_loop(tiles)
	var farthest int = Find_farthest_distance(distances)
	fmt.Println(farthest)
}
