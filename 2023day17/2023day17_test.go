package main

import (
	"testing"
)

func TestBuilds_matrix(t *testing.T) {
	lines := []string{"12", "34"}
	mat := Builds_matrix(lines)
	if mat[0][0] != 1 {
		t.Errorf("mat[0][0] is %d instead of %d", mat[0][0], 1)
	}
	if mat[1][0] != 3 {
		t.Errorf("mat[1][0] is %d instead of %d", mat[1][0], 3)
	}
	if mat[0][1] != 2 {
		t.Errorf("mat[0][1] is %d instead of %d", mat[0][1], 2)
	}
}

func TestAdd_task(t *testing.T) {
	var queue Priority_queue
	task1 := Task{0, 0, 0, 1}
	task2 := Task{0, 0, 0, 2}
	queue.Add_task(task1, 1)
	queue.Add_task(task2, 2)
	if len(queue.heap) != 2 {
		t.Errorf("len(queue.heap) is %d instead of %d", len(queue.heap), 2)
	}
	if queue.heap[0].task.(Task).x != 0 {
		t.Errorf("queue.heap[0].task.x is %d instead of %d", queue.heap[0].task.(Task).x, 0)
	}
	if queue.heap[0].priority != 2 {
		t.Errorf("queue.heap[0].priority is %d instead of %d", queue.heap[0].priority, 2)
	}
}

func TestPop_task(t *testing.T) {
	var queue Priority_queue
	task1 := Task{0, 0, 0, 1}
	task2 := Task{0, 0, 0, 2}
	queue.Add_task(task1, 1)
	queue.Add_task(task2, 2)
	output := queue.Pop_task()
	if output.x != task2.x {
		t.Errorf("output.x is %d instead of %d", output.x, task2.x)
	}
	if output.direction != task2.direction {
		t.Errorf("output.direction is %d instead of %d", output.direction, task2.direction)
	}
}

func TestCalculates_path(t *testing.T) {
	lines := []string{"12", "34"}
	heat_map := Builds_matrix(lines)
	var heat_loss int = Calculates_path(heat_map)
	if heat_loss != 6 {
		t.Errorf("heat_loss is %d instead of %d", heat_loss, 6)
	}
	lines1 := []string{"2413432311323", "3215453535623", "3255245654254", "3446585845452", "4546657867536", "1438598798454", "4457876987766", "3637877979653", "4654967986887", "4564679986453", "1224686865563", "2546548887735", "4322674655533"}
	heat_map = Builds_matrix(lines1)
	heat_loss1 := Calculates_path(heat_map)
	if heat_loss1 != 102 {
		t.Errorf("heat_loss is %d instead of %d", heat_loss1, 102)
	}
}
