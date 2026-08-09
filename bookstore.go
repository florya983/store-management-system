package main

import "fmt"

type bookstore struct{
	 Books []book
	 Sales int
	 TotalMoney float32 	 
}

func (store bookstore) SearchBook(ID int)bool{
	for i:=0;i<len(store.Books);i++{
		if store.Books[i].ID==ID{
			return true
		}
	}
	return false
}

func (store bookstore) HasStock(ID int)bool{
	for i:=0;i<len(store.Books);i++{
		if store.Books[i].ID==ID{
			if store.Books[i].Stock==0{
				return false
			}
		}
	}
	return true
}

func (store *bookstore) SellBook(ID int){
	if !store.SearchBook(ID){
		fmt.Println("Book not exists")
		return 
	}
	if !store.HasStock(ID){
		fmt.Println("Book out of stock")
		return
	}
	for i:=0;i<len(store.Books);i++{
		if store.Books[i].ID==ID{
			store.Sales=store.Sales+1
		    store.TotalMoney=store.TotalMoney+store.Books[i].Price
			store.Books[i].CountSale=store.Books[i].CountSale+1
		 	store.Books[i].CountPriceSale=store.Books[i].CountPriceSale+ store.Books[i].Price
			store.Books[i].Stock--	
			return
		    }
	    }

}
func (store *bookstore) DeleteBook (ID int){
	if !store.SearchBook(ID){
		fmt.Println("Book not exists")
		return 
	}
	for i:=0;i<len(store.Books);i++{
		if store.Books[i].ID==ID{
			store.Books=append(store.Books[:i],store.Books[i+1:]...)
		}
    }		
}
func (store bookstore) CountBooksByCategory(categoryBook string)int{
	count:=0
	for i:=0;i<len(store.Books);i++{
		if store.Books[i].Category==categoryBook{
			count++
		}
	}
	return count

}


func (store bookstore) BestSellingBook() string{
	max:=0
	title:=""
	for i:=0;i<len(store.Books);i++{
		if store.Books[i].CountSale>max{
			max=store.Books[i].CountSale
			title=store.Books[i].Title
		}
    }
	return title
}


func (store *bookstore) ModifyBook(ID int,price float32){
	if !store.SearchBook(ID){
		fmt.Println("Book not exists")
		return 
	}
	for i:=0;i<len(store.Books);i++{
		if store.Books[i].ID == ID{
			store.Books[i].Price=price
		}
	}
	return
}

func (store *bookstore) AddBook(newbook book){
	store.Books=append(store.Books,newbook)
	
}

func (store bookstore) TotalSales()int{
	count:=0
	for i:=0;i<len(store.Books);i++{
		 count=count+store.Books[i].CountSale
					
	}
	return count
}

func (store bookstore) TotalMoneySales()float32{
	count:=float32(0)
	for i:=0;i<len(store.Books);i++{
			count=count+store.Books[i].CountPriceSale
		}
	return count					
}

func (store bookstore) ListBook(){
	for i:=0;i<len(store.Books);i++{
	fmt.Printf("ID: %d\nTitle: %s\nPrice: %.2f\nStock: %d\nCategory: %s\nCountSale: %d\nCountPriceSale: %.2f\n",
    store.Books[i].ID,
    store.Books[i].Title,
    store.Books[i].Price,
    store.Books[i].Stock,
    store.Books[i].Category,
    store.Books[i].CountSale,
    store.Books[i].CountPriceSale,)
	}
}
