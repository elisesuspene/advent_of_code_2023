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

func Gets_digits(string_table []string) []float64 {
	var num_table []float64
	for _, line := range string_table {
		runeSlice := []rune(line)
		var first_digit string = ""
		for i := 0; i < len(runeSlice); i++ {
			_, err_first := strconv.Atoi(string(runeSlice[i]))
			if err_first == nil && first_digit == "" { //meaning that i is the rank of an int
				first_digit = string(runeSlice[i])
			}
		}
		var last_digit string = ""
		for i := len(runeSlice) - 1; i >= 0; i-- {
			_, err_last := strconv.Atoi(string(runeSlice[i]))
			if err_last == nil && last_digit == "" { //meaning that i is the rank of an int
				last_digit = string(runeSlice[i])
			}
		}
		num, _ := strconv.ParseFloat(first_digit+last_digit, 64)
		num_table = append(num_table, num)
	}
	return num_table
}

func Sum_of_cal_values(num_table []float64) float64 {
	var s float64 = 0
	for _, num := range num_table {
		s += num
	}
	return s
}

func main() {
	int_table := File_to_string_table("input.txt")
	num_table := Gets_digits(int_table)
	fmt.Println(Sum_of_cal_values(num_table))
}
