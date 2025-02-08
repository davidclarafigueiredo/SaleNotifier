package connect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

func Connect() {
	url := os.Getenv("REQUEST")
	res, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to site")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Could not read body")
	}
	if err := json.Unmarshal(body, &jsonbody); err != nil {
		log.Error().Err(err).Msg("Could not unmarshal json bytestream")
	}
	fmt.Printf("%s\n", jsonbody.Prices[0].RegularPrice.Amount)
}
