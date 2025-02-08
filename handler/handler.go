package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type Message struct {
	Personalized bool   `json:"personalized"`
	Country      string `json:"country"`
	Prices       []struct {
		TitleID      int64  `json:"title_id"`
		SalesStatus  string `json:"sales_status"`
		RegularPrice struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
			RawValue string `json:"raw_value"`
		} `json:"regular_price"`
		DiscountPrice struct {
			Amount        string    `json:"amount"`
			Currency      string    `json:"currency"`
			RawValue      string    `json:"raw_value"`
			StartDatetime time.Time `json:"start_datetime"`
			EndDatetime   time.Time `json:"end_datetime"`
		} `json:"discount_price"`
		GoldPoint struct {
			BasicGiftGp     string `json:"basic_gift_gp"`
			BasicGiftRate   string `json:"basic_gift_rate"`
			ConsumeGp       string `json:"consume_gp"`
			ExtraGoldPoints []any  `json:"extra_gold_points"`
			GiftGp          string `json:"gift_gp"`
			GiftRate        string `json:"gift_rate"`
		} `json:"gold_point"`
	} `json:"prices"`
}

type prodcut struct {
	Title string
	Price int
}

var jsonbody Message
var game prodcut

func unmarshal(body []byte) bool {
	if err := json.Unmarshal(body, &jsonbody); err != nil {
		log.Error().Err(err).Msg("Could not unmarshal json bytestream")
	}
	log.Debug().Msgf("Title ID: %s", jsonbody.Prices[0].DiscountPrice.Amount)
	return jsonbody.Prices[0].DiscountPrice.Amount != ""
}

func GetPrice(body []byte) string {
	if unmarshal(body) {
		game = prodcut{Title: "Title", Price: 0}
		fmt.Printf("Title: %s\n", game.Title)
		return jsonbody.Prices[0].DiscountPrice.Amount
	}
	return jsonbody.Prices[0].RegularPrice.Amount
}
