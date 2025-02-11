package main

import (
	"github.com/davidclarafigueiredo/SaleNotifier/actions"
	"github.com/davidclarafigueiredo/SaleNotifier/config"
)

func main() {
	config.Init()
	//fmt.Printf("%s\n", handler.GetPrice(connect.Connect()))
	//handler.SendMail()

	actions.CreateWishlistEntries()
	actions.SaleChecker()
}
