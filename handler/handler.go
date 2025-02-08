package handler

import (
	"encoding/json"
	"errors"
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

var jsonbody Message

func Unmarshal(body []byte) error {
	if err := json.Unmarshal(body, &jsonbody); err != nil {
		log.Error().Err(err).Msg("Could not unmarshal json bytestream")
	}
	log.Debug().Msgf("Title ID: %s", jsonbody.Prices[0].DiscountPrice.Amount)
	if jsonbody.Prices[0].DiscountPrice.Amount == "" {
		err := errors.New("400: no discount price found")
		return err
	}
	return nil
}
