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
	//convert to string
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

type Set struct {
	game_id int
	set_id  int
	cubes   string
}

func Parsing_to_sets(lines []string) [][]Set {
	var games [][]Set
	for _, line := range lines {
		var game []Set
		//separating "Game 1" from the results of Game 1
		var split_line []string = strings.Split(line, ":")
		var game_title string = split_line[0]
		var results string = split_line[1]
		//getting the game id
		var split_id []string = strings.Split(game_title, " ")
		var game_id_string string = split_id[1]
		game_id, _ := strconv.Atoi(game_id_string)
		//separating each set
		var split_sets []string = strings.Split(results, ";")
		for i, set_string := range split_sets {
			var set_id int = i
			var cubes string = set_string[1:] //because the first character is a blank space
			set := Set{game_id, set_id, cubes}
			game = append(game, set)
		}
		games = append(games, game)
	}
	return games
}

func Parsing_each_set(set Set) [][]string {
	var output [][]string
	var cubes string = set.cubes
	var split_colors []string = strings.Split(cubes, ",")
	for _, color := range split_colors {
		var words []string = strings.Split(color, " ")
		var parsed_color []string
		for _, word := range words {
			if word != "" && word != " " {
				parsed_color = append(parsed_color, word)
			}
		}
		output = append(output, parsed_color)
	}
	return output
}

func Possible_set(set Set) bool {
	var output bool = true
	var parsed_colors [][]string = Parsing_each_set(set)
	for _, color := range parsed_colors {
		var string_num string = color[0]
		var color_name string = color[1]
		num, _ := strconv.Atoi(string_num)
		var limit int
		if color_name == "red" {
			limit = 12
		}
		if color_name == "green" {
			limit = 13
		}
		if color_name == "blue" {
			limit = 14
		}
		if num > limit {
			output = false
		}
	}
	return output
}

func Calculates_answer(games [][]Set) int {
	var output int
	for i, game := range games {
		var possible_game bool = true
		for _, set := range game {
			if !Possible_set(set) {
				possible_game = false
			}
		}
		if possible_game {
			output = output + i + 1
		}
	}
	return output
}

func Max_numbers(game []Set) []int {
	game_output := []int{0, 0, 0} //order is red, green, blue
	for _, set := range game {
		var parsed_colors [][]string = Parsing_each_set(set)
		for _, color := range parsed_colors {
			var index int
			num, _ := strconv.Atoi(color[0])
			var color_name string = color[1]
			if color_name == "red" {
				index = 0
			}
			if color_name == "green" {
				index = 1
			}
			if color_name == "blue" {
				index = 2
			}
			if num > game_output[index] {
				game_output[index] = num
			}
		}
	}
	return game_output
}

func Calculates_answer_2(games [][]Set) int {
	var output int
	for _, game := range games {
		var power int
		var nums []int = Max_numbers(game)
		power = nums[0] * nums[1] * nums[2]
		output += power
	}
	return output
}

func main() {
	fmt.Println("first problem")
	var string_table []string = File_to_string_table("input.txt")
	var games [][]Set = Parsing_to_sets(string_table)
	var output int = Calculates_answer(games)
	fmt.Println(output)
	fmt.Println("second problem")
	var output2 int = Calculates_answer_2(games)
	fmt.Println(output2)
}
