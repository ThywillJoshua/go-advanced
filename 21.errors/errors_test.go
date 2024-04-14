package main

import "testing"

type data struct {
	input    string
	expected bool
}

func TestTimeParse(t *testing.T) {
	table := []data{
		{"19:00:12", true},
		{"1:3:12", true},
		{"bad", false},
		{"19:-0:12", false},
		{"00:59:59", true},
		{"", false},
	}

	for _, data := range table {
		_, err := parseTime(data.input)
		if data.expected && err != nil {
			t.Errorf("%v: %v, error should be nil", data.input, err)
		}
	}
}
