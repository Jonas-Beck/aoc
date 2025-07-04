package main

import (
	_ "embed"
	"testing"
)

//go:embed day04a.txt
var filea string

//go:embed day04b.txt
var fileb string

func TestDay01a(t *testing.T) {
	actual := day04a(filea)
	expected := 24000.0

	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01b(t *testing.T) {
	actual := day04b(filea)
	expected := 45000.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
