package stocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/** GetDetails
 * * returns the base summary of a ticker from Yahoo Finance API
 */
func GetFullDetails(ticker string) *Result {
	res, err := http.Get(fmt.Sprintf(yahooURL, ticker))
	if err != nil {
		println(err)
		return nil
	}
	var out TLR
	d := json.NewDecoder(res.Body)
	d.Decode(&out)
	if err != nil || len(out.QuoteResponse.Result) == 0 {
		if err != nil {
			println(err)
		}
		return nil
	}
	return &out.QuoteResponse.Result[0]
}

/** GetPriceAndPercentage
* * retrieves the current trading price of a ticker from Yahoo
 */
func GetPriceAndPercentage(ticker string) (string, string, string) {
	det := GetFullDetails(ticker)
	if det == nil {
		return "", "", ""
	}

	trimmedPrice := strconv.FormatFloat(det.RegularMarketPrice, 'f', 2, 64)
	trimmedPercent := strconv.FormatFloat(det.RegularMarketChangePercent, 'f', 2, 64)

	var direction string = ""
	if det.RegularMarketChangePercent < 0 {
		//We are down
		direction = "↓"
	} else {
		//We are up
		direction = "↑"
	}

	return "$" + trimmedPrice, trimmedPercent + "%", direction
}
