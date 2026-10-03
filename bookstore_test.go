package main

import "testing"

func TestSearchBookNegativeID(t *testing.T) {
	store := bookstore{}
	result := store.SearchBook(-2)
	expected := false
	if result != expected {
		t.Errorf("result was %v, expected %v", result, expected)
	}

}

func TestDeleteBookNegativeID(t *testing.T) {
	store := bookstore{}
	store.DeleteBook(-4)
	if len(store.Books) != 0 {
		t.Errorf("the bookstore should have 0 books, got %d", len(store.Books))
	}

}
