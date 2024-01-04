package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type Module struct {
	input             []int
	output            []int
	mod_type          string
	name              string
	status            string   //for a flip-flop module, initially off
	last_input_pulses []string //for a conjunction module, initially low pulses
}

func Parsing_input(lines []string) []Module {
	var mod_list []Module
	var output_names_list [][]string
	//adding a button module
	var empty_input []int
	var empty_output []int
	var empty_last_input_pulses []string
	mod_list = append(mod_list, Module{empty_input, empty_output, "", "button", "", empty_last_input_pulses})
	output_names_list = append(output_names_list, []string{"broadcaster"})
	//reading the file
	for _, line := range lines {
		var split_line []string = strings.Split(line, " -> ")
		var mod_type string
		var name string
		var status string
		if split_line[0] == "broadcaster" {
			name = "broadcaster" //broadcaster and button are names and have no type
		}
		if split_line[0] != "broadcaster" {
			mod_type = string(split_line[0][0])
			name = split_line[0][1:]
		}
		if mod_type == "%" {
			status = "off"
		}
		mod_list = append(mod_list, Module{empty_input, empty_output, mod_type, name, status, empty_last_input_pulses})
		var output_names []string
		var split_output_list []string = strings.Split(split_line[1], ", ")
		output_names = append(output_names, split_output_list...)
		output_names_list = append(output_names_list, output_names)

	}
	//getting the modules designated by their names in output_names
	for i, _ := range mod_list {
		var output_names []string = output_names_list[i]
		for _, output_name := range output_names {
			for k, output_mod := range mod_list {
				if output_mod.name == output_name {
					mod_list[i].output = append(mod_list[i].output, k)
				}
			}
		}
	}
	//getting the input module indexes of each module
	for l, mod := range mod_list {
		for _, output_mod_index := range mod.output {
			mod_list[output_mod_index].input = append(mod_list[output_mod_index].input, l)
		}
	}
	//initialising last_input_pulses for a conjunction module
	for i, _ := range mod_list {
		if mod_list[i].mod_type == "&" {
			for k := 0; k < len(mod_list[i].input); k++ {
				mod_list[i].last_input_pulses = append(mod_list[i].last_input_pulses, "low")
			}
		}
	}
	return mod_list
}

type Pulse struct {
	sender     int
	pulse_type string
}

func Flip_flop(mod_list []Module, mod_index int, sent_pulses []Pulse, impending_pulses []Pulse) ([]Module, []Pulse, []Pulse) {
	if len(impending_pulses) != 0 {
		var pulse Pulse = impending_pulses[0]
		if pulse.pulse_type == "low" {
			if mod_list[mod_index].status == "on" {
				mod_list[mod_index].status = "off"
				impending_pulses = append(impending_pulses, Pulse{mod_index, "low"})
				return mod_list, sent_pulses, impending_pulses
			}
			if mod_list[mod_index].status == "off" {
				mod_list[mod_index].status = "on"
				impending_pulses = append(impending_pulses, Pulse{mod_index, "high"})
			}
		}
	}
	return mod_list, sent_pulses, impending_pulses
}

func Conjunction(mod_list []Module, mod_index int, sent_pulses []Pulse, impending_pulses []Pulse) ([]Module, []Pulse, []Pulse) {
	var pulse Pulse = impending_pulses[0]
	for i, input_mod_index := range mod_list[mod_index].input {
		if mod_list[input_mod_index].name == mod_list[pulse.sender].name {
			mod_list[mod_index].last_input_pulses[i] = pulse.pulse_type
		}
	}
	var all_high bool = true
	for _, remembered_pulse := range mod_list[mod_index].last_input_pulses {
		if remembered_pulse != "high" {
			all_high = false
		}
	}
	if all_high {
		impending_pulses = append(impending_pulses, Pulse{mod_index, "low"})
	}
	if !all_high {
		impending_pulses = append(impending_pulses, Pulse{mod_index, "high"})
	}
	return mod_list, sent_pulses, impending_pulses
}

func Broadcast(mod_list []Module, mod_index int, sent_pulses []Pulse, impending_pulses []Pulse) ([]Module, []Pulse, []Pulse) {
	var pulse Pulse = impending_pulses[0]
	impending_pulses = append(impending_pulses, Pulse{mod_index, pulse.pulse_type})
	return mod_list, sent_pulses, impending_pulses
}

func Button(mod_list []Module) ([]Module, []Pulse, []Pulse) {
	var sent_pulses []Pulse
	impending_pulses := []Pulse{{0, "low"}}
	return mod_list, sent_pulses, impending_pulses
}

func Processes_pulses(mod_list []Module) ([]Module, []Pulse) {
	mod_list, sent_pulses, impending_pulses := Button(mod_list)
	for len(impending_pulses) != 0 {
		var pulse Pulse = impending_pulses[0]
		for _, receiver_index := range mod_list[pulse.sender].output {
			var receiver Module = mod_list[receiver_index]
			if receiver.mod_type == "%" {
				mod_list, sent_pulses, impending_pulses = Flip_flop(mod_list, receiver_index, sent_pulses, impending_pulses)
			}
			if receiver.mod_type == "&" {
				mod_list, sent_pulses, impending_pulses = Conjunction(mod_list, receiver_index, sent_pulses, impending_pulses)
			}
			if receiver.name == "broadcaster" {
				mod_list, sent_pulses, impending_pulses = Broadcast(mod_list, receiver_index, sent_pulses, impending_pulses)
			}
		}
		sent_pulses = append(sent_pulses, impending_pulses[0])
		impending_pulses = impending_pulses[1:]
	}
	return mod_list, sent_pulses
}

func Calculates_answer(mod_list []Module) int {
	var total_sent_pulses []Pulse
	for i := 0; i < 1000; i++ {
		var sent_pulses []Pulse
		mod_list, sent_pulses = Processes_pulses(mod_list)
		total_sent_pulses = append(total_sent_pulses, sent_pulses...)
	}
	var total_high int
	var total_low int
	for _, pulse := range total_sent_pulses {
		if pulse.pulse_type == "high" {
			total_high += len(mod_list[pulse.sender].output)
		}
		if pulse.pulse_type == "low" {
			total_low += len(mod_list[pulse.sender].output)
		}
	}
	return total_high * total_low
}

func main() {
	var lines []string = File_to_string_table("input.txt")
	var mod_list []Module = Parsing_input(lines)
	fmt.Println(Calculates_answer(mod_list))
}
