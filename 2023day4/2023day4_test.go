package main

import (
	"testing"
)

func TestMatches_a_winning_num(t *testing.T) {
	var c1 Card = Makes_card("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 0)
	if !c1.Matches_a_winning_num(0) {
		t.Errorf("found that number at index 0 in Card 1 was not in winning numbers")
	}
	if c1.Matches_a_winning_num(2) {
		t.Errorf("found that number at index 2 in Card 1 was in winning numbers")
	}
}

func TestCounting_points(t *testing.T) {
	var c1 Card = Makes_card("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 0)
	res := c1.Counting_points()
	if res != 8 {
		t.Errorf("found that score for card1 is %d", res)
	}
	var c2 Card = Makes_card("Card 1: 1 | ", 1)
	res = c2.Counting_points()
	if res != 0 {
		t.Errorf("found that score for empty card is %d", res)
	}
}

func TestCalculates_instances(t *testing.T) {
	var c0 Card = Makes_card("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 0)
	var c1 Card = Makes_card("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 1)
	var card_array []Card
	card_array = append(card_array, c0, c1)
	res, _ := c0.Calculates_instances(card_array)
	card_array = res
	var expected_c1_instances int = 2
	if card_array[1].num_of_instances != expected_c1_instances {
		t.Errorf("did not update number of instances of card 2 after reading card 1")
	}
	if card_array[0].num_of_instances != 1 {
		t.Errorf("wrongfully updated num of instances of card 1")
	}
}
