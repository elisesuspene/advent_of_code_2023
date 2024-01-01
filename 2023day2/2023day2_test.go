package main

import (
	"testing"
)

func TestParsing_to_sets(t *testing.T) {
	lines := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
	var games [][]Set = Parsing_to_sets(lines)
	if len(games) != 5 {
		t.Errorf("len(games) is %d instead of 5", len(games))
	}
	if games[0][0].game_id != 1 {
		t.Errorf("games[0].game_id is %d instead of 1", games[0][0].game_id)
	}
}

func TestParsing_each_set(t *testing.T) {
	lines := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
	var games [][]Set = Parsing_to_sets(lines)
	var set Set = games[0][0]
	var parsed_set [][]string = Parsing_each_set(set)
	if len(parsed_set) != 2 {
		t.Errorf("len(parsed_set) is %d instead of 2", len(parsed_set))
	}
	if parsed_set[0][0] != "3" {
		t.Errorf("parsed_set[0][0] is %s instead of 3", parsed_set[0][0])
	}
	if parsed_set[0][1] != "blue" {
		t.Errorf("parsed_set[0][1] is %s instead of blue", parsed_set[0][1])
	}
	var game3 []Set = games[2]
	if len(game3) != 3 {
		t.Errorf("len(game3) is %d instead of 3", len(game3))
	}
	set = game3[0]
	parsed_set = Parsing_each_set(set)
	if len(parsed_set) != 3 {
		t.Errorf("len(parsed_set) is %d instead of 3", len(parsed_set))
	}
	if parsed_set[0][0] != "8" {
		t.Errorf("parsed_set[0][0] is %s instead of 8", parsed_set[0][0])
	}
	if parsed_set[0][1] != "green" {
		t.Errorf("parsed_set[0][1] is %s instead of green", parsed_set[0][1])
	}
}

func TestPossible_set(t *testing.T) {
	lines := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
	var games [][]Set = Parsing_to_sets(lines)
	var set Set = games[0][0]
	var output bool = Possible_set(set)
	if output == false {
		t.Errorf("did not work for games[0][0]")
	}
	set = games[2][0]
	output = Possible_set(set)
	if output == true {
		t.Errorf("did not work for games[2][0]")
	}
}

func TestCalculates_answer(t *testing.T) {
	lines := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
	var games [][]Set = Parsing_to_sets(lines)
	var output int = Calculates_answer(games)
	if output != 8 {
		t.Errorf("function gave %d instead of 8", output)
	}
}
