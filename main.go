package main

import "fmt"

func main() {
	bookOne := book{
		ID:             1,
		Title:          "The Lord of the Rings",
		Stock:          3,
		Price:          200.00,
		Category:       "Fantasy",
		CountSale:      0,
		CountPriceSale: 0,
	}

	bookTwo := book{
		ID:             2,
		Title:          "Harry Potter",
		Stock:          8,
		Price:          800.00,
		Category:       "Fantasy",
		CountSale:      0,
		CountPriceSale: 0,
	}

	bookThree := book{
		ID:             3,
		Title:          "The Conjuring",
		Stock:          5,
		Price:          600.00,
		Category:       "Horror",
		CountSale:      0,
		CountPriceSale: 0,
	}

	store := bookstore{}
	store.AddBook(bookOne)
	store.AddBook(bookTwo)
	store.ListBook()
	store.SellBook(2)
	store.SellBook(1)
	countsBookByCategory := store.CountBooksByCategory("Fantasy")
	fmt.Println("There are", countsBookByCategory, "Fantasy books")
	store.ListBook()
	bestSellingBook := store.BestSellingBook()
	fmt.Println("The best-selling book is", bestSellingBook)
	sales := store.TotalSales()
	totalMoneySales := store.TotalMoneySales()
	fmt.Println("Total book sales:", sales, "and total money earned is:", totalMoneySales, "$")
	store.ModifyBook(2, 900.00)
	store.DeleteBook(1)
	store.AddBook(bookThree)
	store.ListBook()
}

