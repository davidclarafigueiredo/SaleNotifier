package main

import (
	"github.com/davidclarafigueiredo/SaleNotifier/config"
	"github.com/davidclarafigueiredo/SaleNotifier/connect"
)

func main() {
	config.Init()
	connect.Connect()
}
