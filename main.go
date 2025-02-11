package main

import (
	"fmt"

	"github.com/davidclarafigueiredo/SaleNotifier/config"
	"github.com/davidclarafigueiredo/SaleNotifier/connect"
	"github.com/davidclarafigueiredo/SaleNotifier/handler"
	"github.com/davidclarafigueiredo/SaleNotifier/scraper"
)

func main() {
	config.Init()
	fmt.Printf("%s\n", handler.GetPrice(connect.Connect()))
	//handler.SendMail()
	fmt.Printf("Game Title: %s\n", scraper.GetGameTitle())
	fmt.Printf("NSUID: %s\n", scraper.GetNSUID())
}
