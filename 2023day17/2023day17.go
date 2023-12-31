package main

import (
	"bufio"
	"fmt"
	"math"
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

type Task struct {
	x           int
	y           int
	direction   int
	consecutive int
}

type Entry struct {
	priority int
	count    int
	task     interface{}
}

type Priority_queue struct {
	heap         []Entry        //list of entries arranged in a heap
	entry_finder map[Task]Entry //mapping of tasks to entries
	//tasks = key ; []int = values
	counter int
}

const REMOVED = "<removed-task>" //placeholder for a removed task

func (queue *Priority_queue) Siftup(pos int) {
	for pos > 0 {
		var parent_pos int = (pos - 1) / 2
		if queue.heap[pos].priority > queue.heap[parent_pos].priority {
			var old_heap_pos Entry = queue.heap[pos]
			queue.heap[pos] = queue.heap[parent_pos]
			queue.heap[parent_pos] = old_heap_pos
			pos = parent_pos
		}
		if queue.heap[pos].priority <= queue.heap[parent_pos].priority {
			break
		}
	}
}

func (queue *Priority_queue) Heappush(entry Entry) {
	queue.heap = append(queue.heap, entry)
	queue.Siftup(len(queue.heap) - 1)
}

func (queue *Priority_queue) Add_task(task Task, priority int) {
	if queue.entry_finder == nil {
		entry_finder := make(map[Task]Entry)
		queue.entry_finder = entry_finder
	}
	_, is_in := queue.entry_finder[task]
	if is_in {
		queue.Remove_task(task)
	}
	queue.counter += 1
	var count int = queue.counter
	entry := Entry{priority, count, task}
	queue.entry_finder[task] = entry
	queue.Heappush(entry)
}

func (queue *Priority_queue) Remove_task(task Task) {
	entry := queue.entry_finder[task]
	delete(queue.entry_finder, task)
	entry.task = REMOVED
}

func (queue *Priority_queue) Siftdown(arg_pos int) {
	var pos int = arg_pos
	var endpos int = len(queue.heap)
	var startpos int = pos
	var newitem Entry = queue.heap[pos]
	var childpos int = 2*pos + 1 //Left child position
	for childpos < endpos {
		var rightpos int = childpos + 1
		if rightpos < endpos && queue.heap[childpos].priority >= queue.heap[rightpos].priority {
			childpos = rightpos
		}
		queue.heap[pos] = queue.heap[childpos]
		pos = childpos
		childpos = 2*pos + 1
	}
	queue.heap[pos] = newitem
	queue.Siftup1(startpos, pos)
}

func (queue *Priority_queue) Siftup1(startpos int, pos int) {
	var newitem Entry = queue.heap[pos]
	for pos > startpos {
		var parentpos int = (pos - 1) >> 1 //Right shift to divide by 2
		var parent Entry = queue.heap[parentpos]
		if newitem.priority < parent.priority {
			queue.heap[pos] = parent
			pos = parentpos
			continue
		}
		break
	}
	queue.heap[pos] = newitem
}

func (queue *Priority_queue) Heappop() Entry {
	if len(queue.heap) == 0 {
		fmt.Println("pop from an empty heap")
		var empty_entry Entry
		return empty_entry
	}
	var old_heap0 Entry = queue.heap[0]
	//Swap the root (smallest element) with the last element
	queue.heap[0] = queue.heap[len(queue.heap)-1]
	queue.heap[len(queue.heap)-1] = old_heap0
	//Pop the last element (which is now the smallest) and store it
	var smallest Entry = queue.heap[len(queue.heap)-1]
	queue.heap = queue.heap[:len(queue.heap)-1]
	//Restore the heap property by pushing the swapped element down
	queue.Siftdown(0)
	return smallest
}

func (queue *Priority_queue) Pop_task() Task {
	for len(queue.heap) != 0 {
		var last_entry Entry = queue.Heappop()
		if last_entry.task != REMOVED {
			delete(queue.entry_finder, last_entry.task.(Task))
			return last_entry.task.(Task)
		}
	}
	fmt.Println("pop from an empty priority queue")
	var empty_task Task
	return empty_task
}

func Calculates_path(heat_map [][]int) int {
	movement := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	var queue Priority_queue
	entry_finder := make(map[Task]Entry)
	queue.entry_finder = entry_finder
	for y, _ := range heat_map {
		for x, _ := range heat_map[0] {
			for direction := 0; direction < 4; direction++ {
				for consecutive := 1; consecutive < 4; consecutive++ {
					task := Task{x, y, direction, consecutive}
					queue.Add_task(task, 1000000)
				}
			}
		}
	}
	new_task := Task{0, 0, 1, 0}
	queue.Add_task(new_task, 0)
	total_heat := make(map[Task]int)
	total_heat[new_task] = 0
	for {
		t := queue.Pop_task()
		_, is_in := total_heat[t]
		if !is_in {
			total_heat[t] = 1000000
		}
		if t.x == len(heat_map[0])-1 && t.y == len(heat_map)-1 {
			return (total_heat[t])
		}
		neighbors := [][]int{{int(math.Abs(float64((t.direction + 1) % 4))), 1}, {int(math.Abs(float64((t.direction - 1) % 4))), 1}}
		if t.consecutive < 3 {
			forward_neighbor := []int{t.direction, t.consecutive + 1}
			neighbors = append(neighbors, forward_neighbor)
		}
		for _, neighbor := range neighbors {
			new_direction := neighbor[0]
			new_consecutive := neighbor[1]
			new_x := t.x + movement[new_direction][0]
			new_y := t.y + movement[new_direction][1]
			new_t := Task{new_x, new_y, new_direction, new_consecutive}
			if 0 <= new_x && new_x < len(heat_map[0]) && 0 <= new_y && new_y < len(heat_map) {
				new_heat := total_heat[t] + heat_map[new_y][new_x]
				if new_heat < total_heat[new_t] {
					total_heat[new_t] = new_heat
					queue.Add_task(new_t, new_heat)
				}
			}
		}
	}
}

func main() {
	var lines []string = File_to_string_table("input.txt")
	var heat_map [][]int = Builds_matrix(lines)
	var total_heat_loss int = Calculates_path(heat_map)
	fmt.Println(total_heat_loss)
}
