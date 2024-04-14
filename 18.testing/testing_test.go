package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	data := "email@example.com"
	if !isValidEmail(data) {
		t.Errorf("isValidEmail(%v) = false, want true", data)
	}
}

// Using Test Tables to test multiple values
type Data struct {
	input string
	want  bool
}

func TestIsValidEmailV2(t *testing.T) {
	table := []Data{
		{"email@mail.com", true},
		{"invalid-email", true},
	}

	for _, item := range table {
		result := isValidEmail(item.input)
		if result != item.want {
			t.Errorf("isValidEmail(%v) = %v, want is %v", item.input, result, item.want)
		}
	}
}
