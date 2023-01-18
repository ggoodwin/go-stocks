package nyse_stocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/**SECTION - Queries
 * * This section contains all the queries used in the program
 * * The queries are used to retrieve data from the Yahoo Finance API
 */

/**ANCHOR GetDetails
 * @param ticker string
 * @return *Result
 * * @desc returns the base summary of a ticker from Yahoo Finance API
 * @note returns nil if the ticker is not found
 * @note returns nil if there is an error
 */
func GetFullDetails(ticker string) *Result {
	res, err := http.Get(fmt.Sprintf(yahooURL, ticker))
	if err != nil {
		println(err)
		return nil
	}
	var out tLR
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

/**ANCHOR GetPriceAndPercentage
 * @param ticker string
 * @return string, string, string
 * * @desc returns the current price and percentage of a ticker from Yahoo Finance API
 * @note returns empty strings if the ticker is not found
 * @note returns empty strings if there is an error
 * @note returns the price of the stock as a string
 * @note returns the percentage movement of the stock as a string
 * @note returns the direction of movement of the stock as a string
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

//!SECTION
