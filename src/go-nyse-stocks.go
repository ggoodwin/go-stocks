package stocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// The function retrieves full details of a stock using its ticker symbol from Yahoo Finance API and
// returns the result.
func GetFullDetails(ticker string) *Result {
	res, err := http.Get(fmt.Sprintf(yahooURL, ticker))
	if err != nil {
		println(err)
		return nil
	}
	var out tLR
	d := json.NewDecoder(res.Body)
	err = d.Decode(&out)
	if err != nil || len(out.QuoteResponse.Result) == 0 {
		if err != nil {
			println(err)
		}
		return nil
	}
	return &out.QuoteResponse.Result[0]
}

// The function returns the price, percentage change, and direction of a given stock ticker.
func GetPriceAndPercentage(ticker string) (string, string, string) {
	det := GetFullDetails(ticker)
	if det == nil {
		return "", "", ""
	}

	trimmedPrice := strconv.FormatFloat(det.RegularMarketPrice, 'f', 2, 64)
	trimmedPercent := strconv.FormatFloat(det.RegularMarketChangePercent, 'f', 2, 64)

	var direction string
	if det.RegularMarketChangePercent < 0 {
		//We are down
		direction = "↓"
	} else {
		//We are up
		direction = "↑"
	}

	return "$" + trimmedPrice, trimmedPercent + "%", direction
}
