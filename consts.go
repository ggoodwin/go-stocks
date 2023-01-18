package nyse_stocks

/**SECTION - Constants
 * * This section contains all the constants used in the program
 */

/**ANCHOR - yahooUrl
 * * @desc the url used to query the Yahoo Finance API
 * @note the url is formatted with the ticker symbol
 * @note the url is used to retrieve the base summary of a ticker
 * @note the url is used to retrieve the current price and percentage of a ticker
 */
const (
	// YahooURL = Yahoo Finance
	yahooURL = "https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s"
)

//!SECTION
