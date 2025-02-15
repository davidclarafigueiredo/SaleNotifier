package actions

import (
	"fmt"

	"github.com/davidclarafigueiredo/SaleNotifier/connect"
	"github.com/davidclarafigueiredo/SaleNotifier/handler"
	"github.com/davidclarafigueiredo/SaleNotifier/scraper"
	"github.com/shopspring/decimal"
)

// returns a boolean if game is part of the wishlist (i think?)
func Contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func ComparePrice(url string, apiUrl string) bool {

	price, _ := decimal.NewFromString(handler.GetPrice(connect.Connect(apiUrl)))
	discountPrice, _ := decimal.NewFromString(scraper.GetPrice(url))

	if !price.Equal(discountPrice) {
		fmt.Println("Price: ", price)
		fmt.Println("Discount Price: ", discountPrice)
		// fmt.Printf("%s is on sale for %s", scraper.GetGameTitle(url), handler.GetFormPrice(connect.Connect(apiUrl)))
		return true
	}
	return false
}
