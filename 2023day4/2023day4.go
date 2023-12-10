package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	i                int //starts at 0
	winning          []int
	had              []int
	num_of_instances int
}

func (c Card) Calculates_instances(card_array []Card) ([]Card, bool) {
	var winning_nums int
	var changed_something bool = false
	for k, _ := range c.had {
		if c.Matches_a_winning_num(k) {
			winning_nums += 1
		}
	}
	for i := c.i + 1; i <= c.i+winning_nums; i++ {
		if i < len(card_array) {
			card_array[i].num_of_instances += c.num_of_instances
			changed_something = true
		}
	}
	return card_array, changed_something
}

func (c Card) Matches_a_winning_num(k int) bool {
	var winning []int = c.winning
	var had []int = c.had
	var res bool = false
	for _, winning_num := range winning {
		if winning_num == had[k] {
			res = true
		}
	}
	return res
}

func (c Card) Counting_points() int {
	var score int = 0
	var first_was_counted bool = false
	var had []int = c.had
	for k, _ := range had {
		if c.Matches_a_winning_num(k) {
			if first_was_counted {
				score = 2 * score
			}
			if !first_was_counted {
				score += 1
				first_was_counted = true
			}
		}
	}
	return score
}

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

/**makes original cards : their number of instances is 1*/
func Makes_card(line string, i int) Card {
	var output_winning []int
	var output_had []int
	var output_i int = i
	//getting rid of "Card 1:"
	var line_table []string = strings.Split(line, ":")
	var line_without_Card string = line_table[1]
	//separating winning and have numbers
	var lines_w_or_h []string = strings.Split(line_without_Card, "|")
	var winning_string string = lines_w_or_h[0]
	var had_string string = lines_w_or_h[1]
	//separating numbers
	var winning_string_table []string = strings.Split(winning_string, " ")
	var had_string_table []string = strings.Split(had_string, " ")
	for _, winning_num := range winning_string_table {
		num, err := strconv.Atoi(winning_num)
		if err == nil {
			output_winning = append(output_winning, num)
		}
	}
	for _, had_num := range had_string_table {
		num, err := strconv.Atoi(had_num)
		if err == nil {
			output_had = append(output_had, num)
		}
	}
	var output Card = Card{output_i, output_winning, output_had, 1}
	return output
}

func main() {
	fmt.Println("first problem")
	var res int
	var string_table []string = File_to_string_table("input.txt")
	var card_array []Card
	for i, line := range string_table {
		var c Card = Makes_card(line, i)
		card_array = append(card_array, c)
		var points int = c.Counting_points()
		res += points
	}
	fmt.Println(res)
	fmt.Println("second problem")
	for _, c := range card_array {
		new_card_array, _ := c.Calculates_instances(card_array)
		card_array = new_card_array
	}
	var total_cards int = 0
	for _, c := range card_array {
		total_cards += c.num_of_instances
	}
	fmt.Println(total_cards)
}
