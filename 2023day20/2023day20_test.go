package main

import (
	"testing"
)

func TestParsing_input(t *testing.T) {
	lines := []string{"broadcaster -> a, b, c", "%a -> b", "%b -> c", "%c -> inv", "&inv -> a"}
	var mod_list []Module = Parsing_input(lines)
	if len(mod_list) != 6 {
		t.Errorf("len(mod_list) is %d instead of 6", len(mod_list))
	}
	if mod_list[0].name != "button" {
		t.Errorf("mod_list[0].name is %s instead of button", mod_list[0].name)
	}
	if len(mod_list[1].output) != 3 {
		t.Errorf("len(mod_list[1].output) is %d instead of 3", len(mod_list[1].output))
	}
	if mod_list[1].output[1] != 3 {
		t.Errorf("mod_list[1].output[1] is %d instead of 3", mod_list[1].output[1])
	}
	if mod_list[mod_list[2].input[0]].name != "broadcaster" {
		t.Errorf("mod_list[mod_list[2].input[0]].name is %s instead of broadcaster", mod_list[mod_list[2].input[0]].name)
	}
}

func TestProcesses_pulses(t *testing.T) {
	lines := []string{"broadcaster -> a, b, c", "%a -> b", "%b -> c", "%c -> inv", "&inv -> a"}
	var mod_list []Module = Parsing_input(lines)
	var sent_pulses []Pulse
	mod_list, sent_pulses = Processes_pulses(mod_list)
	var total_high int
	var total_low int
	for _, pulse := range sent_pulses {
		if pulse.pulse_type == "high" {
			total_high += 1
		}
		if pulse.pulse_type == "low" {
			total_low += 1
		}
	}
	if total_high != 4 {
		t.Errorf("total_high is %d instead of 4", total_high)
	}
	if total_low != 6 {
		t.Errorf("total_low is %d instead of 6", total_low)
	}
	if len(mod_list) != 6 {
		t.Errorf("len(mod_list) is %d instead of 6", len(mod_list))
	}
}

func TestCalculates_answer(t *testing.T) {
	lines := []string{"broadcaster -> a, b, c", "%a -> b", "%b -> c", "%c -> inv", "&inv -> a"}
	var mod_list []Module = Parsing_input(lines)
	var output int = Calculates_answer(mod_list)
	if output != 32000000 {
		t.Errorf("output is %d instead of 32000000", output)
	}
	lines1 := []string{"broadcaster -> a", "%a -> inv, con", "&inv -> b", "%b -> con", "&con -> output"}
	mod_list1 := Parsing_input(lines1)
	output1 := Calculates_answer(mod_list1)
	if output1 != 11687500 {
		t.Errorf("output1 is %d instead of 11687500", output1)
	}
}
