package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dict{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != nil {
		t.Errorf("got %q want %q", got, want)
	}
}
