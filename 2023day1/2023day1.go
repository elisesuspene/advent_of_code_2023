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

func Equal_slices(slice1 []float64, slice2 []float64) bool {
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

func Gets_digits(string_table []string) []float64 {
	var num_table []float64
	for _, line := range string_table {
		runeSlice := []rune(line)
		var first_digit string = ""
		for i := 0; i < len(runeSlice); i++ {
			if first_digit == "" {
				_, err_first := strconv.Atoi(string(runeSlice[i]))
				if err_first == nil { //meaning that i is the rank of an int
					first_digit = string(runeSlice[i])
				}
			}
		}
		var last_digit string = ""
		for i := len(runeSlice) - 1; i >= 0; i-- {
			if last_digit == "" {
				_, err_last := strconv.Atoi(string(runeSlice[i]))
				if err_last == nil { //meaning that i is the rank of an int
					last_digit = string(runeSlice[i])
				}
			}
		}
		num, _ := strconv.ParseFloat(first_digit+last_digit, 64)
		num_table = append(num_table, num)
	}
	return num_table
}

func StartOfSpelledOutDigit(line string, i int) (bool, string) {
	spelled_out := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var res bool = false
	var digit_found string = ""
	runeSlice := []rune(line)
	for _, digit := range spelled_out {
		var n int = len(digit)
		if i+n <= len(runeSlice) {
			var substring string = string(runeSlice[i : i+n])
			if substring == digit {
				res = true
				digit_found = digit
			}
		}
	}
	return res, digit_found
}

func Converts_spelled_out_to_int(input string) string {
	spelled_out_table := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	num_table := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, spelled_out := range spelled_out_table {
		var equal bool = true
		if len(spelled_out) == len(input) {
			for j, _ := range spelled_out {
				if spelled_out[j] != input[j] {
					equal = false
				}
			}
		}
		if equal {
			return fmt.Sprint(num_table[i])
		}
	}
	return ""
}

func EndOfSpelledOutDigit(line string, i int) (bool, string) {
	spelled_out := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var res bool = false
	var digit_found string = ""
	runeSlice := []rune(line)
	for _, digit := range spelled_out {
		var n int = len(digit)
		if i-n+1 >= 0 {
			var substring string = string(runeSlice[i-n+1 : i+1])
			if substring == digit {
				res = true
				digit_found = digit
			}
		}
	}
	return res, digit_found
}

func Gets_digits2(string_table []string) []float64 {
	var num_table []float64
	for _, line := range string_table {
		runeSlice := []rune(line)
		var first_digit string = ""
		for i := 0; i < len(runeSlice); i++ {
			if first_digit == "" {
				_, err_first := strconv.Atoi(string(runeSlice[i]))
				is_start, digit_found := StartOfSpelledOutDigit(line, i)
				if is_start {
					first_digit = Converts_spelled_out_to_int(digit_found)
				}
				if err_first == nil && first_digit == "" { //meaning that i is the rank of an int
					first_digit = string(runeSlice[i])
				}
			}
		}
		var last_digit string = ""
		for i := len(runeSlice) - 1; i >= 0; i-- {
			if last_digit == "" {
				_, err_last := strconv.Atoi(string(runeSlice[i]))
				is_end, digit_found := EndOfSpelledOutDigit(line, i)
				if is_end {
					last_digit = Converts_spelled_out_to_int(digit_found)
				}
				if err_last == nil && last_digit == "" { //meaning that i is the rank of an int
					last_digit = string(runeSlice[i])
				}
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
	fmt.Println("first problem :")
	int_table := File_to_string_table("input.txt")
	num_table := Gets_digits(int_table)
	fmt.Println(Sum_of_cal_values(num_table))
	fmt.Println("second problem :")
	num_table = Gets_digits2(int_table)
	fmt.Println(Sum_of_cal_values(num_table))
}
