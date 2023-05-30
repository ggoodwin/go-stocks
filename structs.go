// Package stocks is defining a Go package named "stocks" which contains various types and structs
// related to financial data for stocks and cryptocurrencies.
package stocks

// tLR type contains a field called QuoteResponse of type quoteResponse with a JSON tag.
// @property {quoteResponse} QuoteResponse - `QuoteResponse` is a struct type that contains information
// related to a stock quote response. It has a single field `quoteResponse` of type `quoteResponse`.
// The `json:"quoteResponse"` tag is used to specify the name of the field when encoding or decoding
// JSON data.
type tLR struct {
	QuoteResponse quoteResponse `json:"quoteResponse"`
}

// Result type defines a struct for storing financial data related to a stock or cryptocurrency.
// @property {string} Symbol - The stock symbol or ticker symbol of the company.
// @property {string} ShortName - The short name or abbreviation of the company or asset being traded.
// @property {string} QuoteSourceName - The name of the source providing the quote data for the
// security.
// @property {string} Language - The language used for the information provided in the result.
// @property {string} Region - The region where the stock is traded.
// @property {string} QuoteType - The type of quote being provided for the security, such as "EQUITY"
// or "ETF".
// @property {bool} Triggerable - A boolean value indicating whether or not the security is triggerable
// for trading. A triggerable security is one that can be traded based on certain conditions being met,
// such as a specific price or volume level.
// @property {string} Currency - The currency in which the stock is traded.
// @property {string} Exchange - The stock exchange where the stock is traded.
// @property {string} MessageBoardID - This property represents the unique identifier for the message
// board associated with the stock or asset. It can be used to retrieve or post messages related to the
// stock or asset on the message board.
// @property {string} ExchangeTimezoneName - The name of the timezone where the exchange is located.
// @property {string} ExchangeTimezoneShortName - The abbreviated name of the timezone where the stock
// exchange is located. For example, "EST" for Eastern Standard Time.
// @property {int} GmtOffSetMilliseconds - The number of milliseconds offset from GMT (Greenwich Mean
// Time) for the stock's exchange. This is used to convert the exchange's local time to GMT.
// @property {string} Market - The market in which the security is traded.
// @property {bool} EsgPopulated - A boolean value indicating whether the company has ESG
// (Environmental, Social, and Governance) data available.
// @property {int64} FirstTradeDateMilliseconds - The Unix timestamp (in milliseconds) of the first
// trade date for the stock or asset.
// @property {float64} RegularMarketChange - The change in the regular market price of the security
// from the previous trading day's closing price.
// @property {float64} RegularMarketChangePercent - The percentage change in the regular market price
// compared to the previous day's closing price.
// @property {int} RegularMarketTime - The time of the last trade in the regular market hours,
// expressed in Unix epoch time (number of seconds since January 1, 1970).
// @property {float64} RegularMarketPrice - The current market price of the stock or asset.
// @property {float64} RegularMarketDayHigh - The highest price at which a security was traded during
// the regular trading hours of the market.
// @property {string} RegularMarketDayRange - RegularMarketDayRange is a string property that
// represents the range of the regular market day for a stock or security. It typically includes the
// lowest and highest prices at which the security was traded during the regular market hours. For
// example, "50.00 - 55.00" would indicate that the
// @property {float64} RegularMarketDayLow - The lowest price at which a security was traded during
// regular market hours.
// @property {int64} RegularMarketVolume - The total number of shares or contracts traded for the
// security during the regular trading hours of the market.
// @property {float64} RegularMarketPreviousClose - The previous day's closing price of the stock or
// asset.
// @property {string} FullExchangeName - The full name of the exchange where the stock is traded.
// @property {float64} RegularMarketOpen - The opening price of the stock or asset in the regular
// market session.
// @property {int64} AverageDailyVolume3Month - The average daily trading volume of the security over
// the past 3 months.
// @property {int64} AverageDailyVolume10Day - The average daily trading volume over the past 10 days
// for the given stock or asset.
// @property {int} StartDate - The start date of the stock or asset in Unix timestamp format (number of
// seconds since January 1, 1970).
// @property {string} CoinImageURL - The URL of the image associated with the coin.
// @property {float64} FiftyTwoWeekLowChange - The change in the stock's price from its 52-week low.
// @property {float64} FiftyTwoWeekLowChangePercent - The percentage change from the current market
// price to the 52-week low price.
// @property {string} FiftyTwoWeekRange - The range of the stock price over the past 52 weeks, from the
// lowest price to the highest price.
// @property {float64} FiftyTwoWeekHighChange - The change in the stock's price from its 52-week high
// to its current price.
// @property {float64} FiftyTwoWeekHighChangePercent - The percentage change in the stock's price from
// its 52-week high.
// @property {float64} FiftyTwoWeekLow - The lowest price that a stock has traded at in the past 52
// weeks.
// @property {float64} FiftyTwoWeekHigh - The highest price that a stock has traded at during the past
// 52 weeks (1 year).
// @property {float64} FiftyDayAverage - The 50-day moving average of the stock's price. This is the
// average price of the stock over the past 50 trading days.
// @property {float64} FiftyDayAverageChange - The change in the fifty day moving average price of the
// stock.
// @property {float64} FiftyDayAverageChangePercent - The percentage change in the fifty-day moving
// average of the stock price.
// @property {float64} TwoHundredDayAverage - The two hundred day moving average of the stock's price.
// @property {float64} TwoHundredDayAverageChange - The change in the two hundred day average price of
// the stock.
// @property {float64} TwoHundredDayAverageChangePercent - The percentage change in the two hundred day
// average price of the stock.
// @property {int64} MarketCap - The market capitalization of the stock, which is the total value of
// all outstanding shares of the company's stock. It is calculated by multiplying the current stock
// price by the total number of outstanding shares.
// @property {int} SourceInterval - The frequency at which the data is updated, in seconds.
// @property {int} ExchangeDataDelayedBy - The number of minutes that the exchange data is delayed by.
// @property {bool} Tradeable - A boolean value indicating whether the security is currently tradeable
// on the exchange.
// @property {string} MarketState - The current state of the market for the given symbol, such as "PRE"
// for pre-market, "REGULAR" for regular trading hours, or "POST" for after-hours trading.
type Result struct {
	Symbol                            string  `json:"symbol"`
	ShortName                         string  `json:"shortName"`
	QuoteSourceName                   string  `json:"quoteSourceName"`
	Language                          string  `json:"language"`
	Region                            string  `json:"region"`
	QuoteType                         string  `json:"quoteType"`
	Triggerable                       bool    `json:"triggerable"`
	Currency                          string  `json:"currency"`
	Exchange                          string  `json:"exchange"`
	MessageBoardID                    string  `json:"messageBoardId"`
	ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
	Market                            string  `json:"market"`
	EsgPopulated                      bool    `json:"esgPopulated"`
	FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
	RegularMarketChange               float64 `json:"regularMarketChange"`
	RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
	RegularMarketTime                 int     `json:"regularMarketTime"`
	RegularMarketPrice                float64 `json:"regularMarketPrice"`
	RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
	RegularMarketDayRange             string  `json:"regularMarketDayRange"`
	RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
	RegularMarketVolume               int64   `json:"regularMarketVolume"`
	RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
	FullExchangeName                  string  `json:"fullExchangeName"`
	RegularMarketOpen                 float64 `json:"regularMarketOpen"`
	AverageDailyVolume3Month          int64   `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day           int64   `json:"averageDailyVolume10Day"`
	StartDate                         int     `json:"startDate"`
	CoinImageURL                      string  `json:"coinImageUrl"`
	FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
	MarketCap                         int64   `json:"marketCap"`
	SourceInterval                    int     `json:"sourceInterval"`
	ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
	Tradeable                         bool    `json:"tradeable"`
	MarketState                       string  `json:"marketState"`
}

// quoteResponse type contains a list of Result objects and an optional error field.
// @property {[]Result} Result - Result is a slice of Result structs that contain the response data for
// a quote request.
// @property Error - The "Error" property is an interface{} type which can hold any type of value,
// including an error message or a null value. It is used to indicate if there was an error while
// processing the request or not. If there was an error, the value of this property will be set to the
type quoteResponse struct {
	Result []Result    `json:"result"`
	Error  interface{} `json:"error"`
}
