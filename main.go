package main

import (
	"github.com/davidclarafigueiredo/SaleNotifier/config"
	"github.com/davidclarafigueiredo/SaleNotifier/connect"
	"github.com/davidclarafigueiredo/SaleNotifier/handler"
)

func main() {
	config.Init()
	handler.Unmarshal(connect.Connect())
}
