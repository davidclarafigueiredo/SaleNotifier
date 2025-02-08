package main

import (
	"fmt"

	"github.com/davidclarafigueiredo/SaleNotifier/config"
	"github.com/davidclarafigueiredo/SaleNotifier/connect"
	"github.com/davidclarafigueiredo/SaleNotifier/handler"
)

func main() {
	config.Init()
	fmt.Printf("%s", handler.GetPrice(connect.Connect()))
}
