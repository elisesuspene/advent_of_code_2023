package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// string table to int table
func string_to_int(string_table []string) []int {
	var res []int
	for _, line := range string_table {
		number, _ := strconv.Atoi(line)
		res = append(res, number)
	}
	fmt.Println("The string table was successfully converted to an int table.")
	return res
}

func count_increases(table []int) int {
	var table_diff []int
	for i := 1; i < len(table); i++ {
		table_diff = append(table_diff, table[i]-table[i-1])
	}
	var sum int = 0
	for i := 0; i < len(table_diff); i++ {
		if table_diff[i] > 0 {
			sum += 1
		}
	}
	return sum
}

func main() {
	var lines []string = file_to_string_table("input.txt")
	var table []int = string_to_int(lines)
	fmt.Println(count_increases(table))
}
