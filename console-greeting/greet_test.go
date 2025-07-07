package main

import (
	"testing"
)

func TestGreetWithName(t *testing.T) {
	result := greet("Rahul")
	expected := "Hello, Rahul! Welcome to Go"

	if expected != result {
		t.Errorf(`actual : %s expected : %s`, result, expected)
	}
}

func TestGreetWithoutName(t *testing.T) {
	actual := greet("")
	expected := "Usage: go run main.go [name]"

	if expected != actual {
		t.Errorf(`actual : %s expected : %s`, actual, expected)
	}
}
