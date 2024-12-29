package main

import (
	_ "embed"
	"testing"
)

//go:embed day06a.txt
var filea string

//go:embed day06b.txt
var fileb string

func TestDay01a(t *testing.T) {
	actual := day06a(filea)
	expected := 24000.0

	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01b(t *testing.T) {
	actual := day06b(filea)
	expected := 45000.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
