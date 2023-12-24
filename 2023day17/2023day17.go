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
	//convert to []string
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

func Builds_matrix(lines []string) [][]int {
	var output [][]int
	for i, _ := range lines {
		var output_row []int
		for j, _ := range lines[i] {
			runeSlice := []rune(lines[i])
			heat_loss := string(runeSlice[j])
			num, _ := strconv.Atoi(heat_loss)
			output_row = append(output_row, num)
		}
		output = append(output, output_row)
	}
	return output
}

type Step struct {
	i            int
	j            int // (i, j) is the indices of the new city block entered
	line_dir     int // can be -1, 0 or 1
	col_dir      int
	n_dir        int //number of steps already taken in that direction
	heat_loss    int
	path_to_step [][]int
}

func Puts_value_in_queue(value Step, queue []Step) []Step {
	var i int = 0
	var val_was_added bool = false
	for i < len(queue)-1 {
		if !val_was_added && queue[i].heat_loss <= value.heat_loss && queue[i+1].heat_loss > value.heat_loss {
			var new_queue []Step
			for j := 0; j <= i; j++ {
				new_queue = append(new_queue, queue[j])
			}
			new_queue = append(new_queue, value)
			for k := i + 1; k < len(queue); k++ {
				new_queue = append(new_queue, queue[k])
			}
			queue = new_queue
			val_was_added = true
		}
		i++
	}
	if !val_was_added {
		if len(queue) == 1 && queue[len(queue)-1].heat_loss > value.heat_loss {
			var new_queue []Step
			var last Step = queue[len(queue)-1]
			new_queue = append(new_queue, value)
			new_queue = append(new_queue, last)
			queue = new_queue
		}
		if len(queue) >= 1 && queue[len(queue)-1].heat_loss <= value.heat_loss {
			queue = append(queue, value)
		}
		if len(queue) == 0 {
			queue = append(queue, value)
		}
	}
	return queue
}

func In_slice(val Step, slice []Step) bool {
	var is_in bool = false
	for _, step := range slice {
		if step.i == val.i && step.j == val.j {
			is_in = true
		}
	}
	return is_in
}

func Calculates_path(matrix [][]int) Step {
	var visited []Step
	var start_i int = 0
	var start_j int = 0
	var end_i int = len(matrix) - 1
	var end_j int = len(matrix[0]) - 1
	var queue []Step
	//starting the path
	poss1 := Step{start_i, start_j + 1, 0, 1, 1, matrix[start_i][start_j+1], [][]int{{start_i, start_j}, {start_i, start_j + 1}}}
	poss2 := Step{start_i + 1, start_j, 1, 0, 1, matrix[start_i+1][start_j], [][]int{{start_i, start_j}, {start_i + 1, start_j}}}
	queue = Puts_value_in_queue(poss1, queue)
	queue = Puts_value_in_queue(poss2, queue)
	var possible_ends []Step
	//looping
	for len(queue) != 0 {
		var queue_top Step = queue[0]
		queue = queue[1:]
		if queue_top.i == end_i && queue_top.j == end_j {
			possible_ends = Puts_value_in_queue(queue_top, possible_ends)
		}
		if !In_slice(queue_top, visited) {
			visited = append(visited, queue_top)
			if queue_top.n_dir < 3 {
				var new_i int = queue_top.i + queue_top.line_dir
				var new_j int = queue_top.j + queue_top.col_dir
				if new_i < len(matrix) && new_j < len(matrix[0]) && new_i >= 0 && new_j >= 0 {
					coords := []int{new_i, new_j}
					path_to_new := append(queue_top.path_to_step, coords)
					new_step := Step{new_i, new_j, queue_top.line_dir, queue_top.col_dir, queue_top.n_dir + 1, matrix[new_i][new_j], path_to_new}
					queue = Puts_value_in_queue(new_step, queue)
				}
			}
			for new_line_dir := -1; new_line_dir < 2; new_line_dir++ {
				for new_col_dir := -1; new_col_dir < 2; new_col_dir++ {
					if new_line_dir-new_col_dir == -1 || new_line_dir-new_col_dir == 1 {
						if (new_line_dir != queue_top.line_dir || new_col_dir != queue_top.col_dir) && (new_line_dir != -queue_top.line_dir || new_col_dir != -queue_top.col_dir) {
							var new_i int = queue_top.i - new_line_dir
							var new_j int = queue_top.j + new_col_dir
							if new_i < len(matrix) && new_j < len(matrix[0]) && new_i >= 0 && new_j >= 0 {
								coords := []int{new_i, new_j}
								path_to_new := append(queue_top.path_to_step, coords)
								new_step := Step{new_i, new_j, new_line_dir, new_col_dir, 1, queue_top.heat_loss + matrix[new_i][new_j], path_to_new}
								queue = Puts_value_in_queue(new_step, queue)
							}
						}
					}
				}
			}
		}

	}
	return possible_ends[0]
}

func main() {
	var lines []string = File_to_string_table("input.txt")
	var mat [][]int = Builds_matrix(lines)
	var last_step Step = Calculates_path(mat)
	fmt.Println(last_step.heat_loss)
}
