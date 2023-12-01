package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// reads file, converts it to a string table
func read_file(filename string) []string {
	//opening file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The file was opened successfully")
	//converting to a string table
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	fmt.Println("The file was successfully converted to a string table.")
	return lines
}

func counts_two_dirs(string_table []string) (horiz_pos int, depth_pos int) {
	horiz_pos = 0 //initialising
	depth_pos = 0
	for _, line := range string_table {
		if string(line[1:4]) == "down" { //down 5 means that the depth increases by 5
			number, _ := strconv.Atoi(string(line[4:len(line)]))
			depth_pos += number
		}
		if string(line[0:7]) == "forward" {
			number, _ := strconv.Atoi(string(line[7:len(line)]))
			horiz_pos += number
		}
	}
	return horiz_pos, depth_pos
}

func main() {
	var lines []string = read_file("input.txt")
	horiz_pos, depth_pos := counts_two_dirs(lines)
	fmt.Println(horiz_pos * depth_pos)
}
