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
func TestSearchBook(t *testing.T) {
	store := bookstore{
		Books: []book{
			{
				ID:             3,
				Title:          "The Conjuring",
				Stock:          5,
				Price:          600.00,
				Category:       "Horror",
				CountSale:      0,
				CountPriceSale: 0,
			},
		},
	}
	result := store.SearchBook(3)
	expected := true
	if result != expected {
		t.Errorf("result was %v, expected %v", result, expected)
	}
}

func TestHasStock(t *testing.T) {
	store := bookstore{
		Books: []book{
			{
				ID:             2,
				Title:          "Harry Potter",
				Stock:          8,
				Price:          800.00,
				Category:       "Fantasy",
				CountSale:      0,
				CountPriceSale: 0,
			},
		},
	}
	result := store.HasStock(2)
	expected := true
	if result != expected {
		t.Errorf("result was %v, expected %v", result, expected)
	}
}

func TestCountBooksByCategory(t *testing.T) {
	store := bookstore{
		Books: []book{
			{
				ID:             1,
				Title:          "The Lord of the Rings",
				Stock:          3,
				Price:          200.00,
				Category:       "Fantasy",
				CountSale:      0,
				CountPriceSale: 0,
			},
			{
				ID:             2,
				Title:          "Harry Potter",
				Stock:          8,
				Price:          800.00,
				Category:       "Fantasy",
				CountSale:      0,
				CountPriceSale: 0,
			},
			{
				ID:             3,
				Title:          "The Conjuring",
				Stock:          5,
				Price:          600.00,
				Category:       "Horror",
				CountSale:      0,
				CountPriceSale: 0,
			},
		},
	}

	result := store.CountBooksByCategory("Fantasy")
	expected := 2
	if result != expected {
		t.Errorf("result was %v, expected %v", result, expected)
	}
}

func TestBestSellingBook(t *testing.T) {
	store := bookstore{
		Books: []book{
			{
				ID:             1,
				Title:          "The Lord of the Rings",
				Stock:          3,
				Price:          200.00,
				Category:       "Fantasy",
				CountSale:      6,
				CountPriceSale: 0,
			},
			{
				ID:             2,
				Title:          "Harry Potter",
				Stock:          8,
				Price:          800.00,
				Category:       "Fantasy",
				CountSale:      0,
				CountPriceSale: 0,
			},
			{
				ID:             3,
				Title:          "The Conjuring",
				Stock:          5,
				Price:          600.00,
				Category:       "Horror",
				CountSale:      9,
				CountPriceSale: 0,
			},
		},
	}
	result := store.BestSellingBook()
	expected := "The Conjuring"
	if result != expected {
		t.Errorf("result was %v, expected %v", result, expected)
	}
}

