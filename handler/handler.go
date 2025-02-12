package handler

import (
	"encoding/json"
	"strconv"

	"github.com/davidclarafigueiredo/SaleNotifier/model"
	"github.com/rs/zerolog/log"
)

// return a ResponseJSON struct from the json bytestream
func unmarshal(body []byte) model.ResponseJSON {
	var response model.ResponseJSON
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal().Err(err).Msg("Could not unmarshal json bytestream")
	}
	log.Debug().Msgf("Title ID: %s", response.Response.Docs[0].Title)
	return response
}

// returns the regular price of a game using the json bytestream
func GetPrice(body []byte) string {
	data := unmarshal(body)
	log.Debug().Msgf("Price for the title: %f\n", data.Response.Docs[0].PriceRegularF)
	return strconv.FormatFloat(data.Response.Docs[0].PriceRegularF, 'f', -1, 64)
}

// returns the price of a game with discount using the json bytestream if it does not have a discount it returns the regular price
func GetFormPrice(body []byte) string {
	data := unmarshal(body)

	if data.Response.Docs[0].PriceHasDiscountB {
		return strconv.FormatFloat(data.Response.Docs[0].PriceDiscountedF, 'f', -1, 64)
	}
	return GetPrice(body)
}

// returns the title of a game using the json bytestream
func GetGameTitle(body []byte) string {
	data := unmarshal(body)
	return data.Response.Docs[0].Title
}
