package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func file_to_string_table(filename string) []string {
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

func gets_digits(string_table []string) []float64 {
	var num_table []float64
	for _, line := range string_table {
		var i int = 0
		_, char := utf8.DecodeRuneInString(line[i:])
		var first_digit string = fmt.Sprint(char)
		_, err_first := strconv.Atoi(first_digit)
		if err_first != nil && i < len(line)-1 {
			i += 1
			_, char = utf8.DecodeRuneInString(line[i:])
			first_digit = fmt.Sprint(char)
			_, err_first = strconv.Atoi(first_digit)
		}
		var j int = len(line) - 1
		_, char = utf8.DecodeRuneInString(line[j:])
		var last_digit string = fmt.Sprint(char)
		_, err_last := strconv.Atoi(last_digit)
		if err_last != nil && j > 0 {
			j = j - 1
			_, char = utf8.DecodeRuneInString(line[j:])
			last_digit = fmt.Sprint(char)
			_, err_last = strconv.Atoi(last_digit)
		}
		if j == 0 && err_last != nil {
			j = i
			_, char = utf8.DecodeRuneInString(line[j:])
			last_digit = fmt.Sprint(char)
			_, err_last = strconv.Atoi(last_digit)
		}
		num, _ := strconv.ParseFloat(first_digit+last_digit, 64)
		num_table = append(num_table, num)
	}
	return num_table
}

func sum_of_cal_values(num_table []float64) float64 {
	var s float64 = 0
	for _, num := range num_table {
		s += num
	}
	return s
}

func main() {
	int_table := file_to_string_table("input.txt")
	num_table := gets_digits(int_table)
	fmt.Println(sum_of_cal_values(num_table))
}
