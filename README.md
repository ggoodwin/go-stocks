<div align="center">
	<h1><img alt="go-nyse-stocks logo" src="https://github.com/ggoodwin/go-nyse-stocks/blob/master/stock.png" height="300" /><br />
		New York Stock Exchange (NYSE) Go Library
	</h1>

[![Go Reference](https://pkg.go.dev/badge/ggoodwin/go-nyse-stocks.svg)](https://pkg.go.dev/github.com/ggoodwin/go-nyse-stocks) [![Go Version](https://img.shields.io/github/go-mod/go-version/ggoodwin/go-nyse-stocks)](https://go.dev/) [![GoReportCard](https://goreportcard.com/badge/github.com/ggoodwin/go-nyse-stocks)](https://goreportcard.com/report/github.com/ggoodwin/go-nyse-stocks) [![CodeFactor](https://www.codefactor.io/repository/github/ggoodwin/go-nyse-stocks/badge)](https://www.codefactor.io/repository/github/ggoodwin/go-nyse-stocks) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ggoodwin/go-nyse-stocks/.github/workflows/go.yml)](https://github.com/ggoodwin/go-nyse-stocks/blob/master/.github/workflows/go.yml) ![Size](https://img.shields.io/github/languages/code-size/ggoodwin/go-nyse-stocks) [![Last Commit](https://img.shields.io/github/last-commit/ggoodwin/go-nyse-stocks)](https://github.com/ggoodwin/go-nyse-stocks/commits/master) [![License](https://img.shields.io/github/license/ggoodwin/go-nyse-stocks)](https://github.com/ggoodwin/go-nyse-stocks/blob/master/LICENSE.md)

</div>
<hr/>

## 🌟 How it works

We use the [Yahoo API] to gather real-time stock data.

## 📦 Installation and Usage

### Go

Make sure you have `Go` installed on your machine.

You can check by running the following command in the `console`

```plain
go version
```

If you don't have `Go` installed, you can [Download Go] and install it.

### Add to your project

Run the following command in the `console`, in the `project directory`, to install the library with `go get`

```plain
go get github.com/ggoodwin/go-nyse-stocks
```

### Importing

Add the import to your `.go` file

```go
import nyse_stocks "github.com/ggoodwin/go-nyse-stocks"
```

## 💰 Usage

### Stock Price and Percentage Change

Add this to your `.go` file

```go
// Parameter: Stock Symbol
// Returns: strings ex: `$123.45`, `1.23%`, `↑`
price, percent, direction := nyse_stocks.GetPriceAndPercentage("AAPL")
```

### Get full stock details

```go
// Parameter: Stock Symbol
// Returns `Result` struct
stock := nyse_stocks.GetFullDetails("AAPL")
```

### Result struct

```go
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
```

## 💻 Dependencies

- [Go]

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report the issue using our [Github Issues] or check out the [Security] details if it is security related.

If you'd like, I welcome any contributions. Please read the [Contributing] document then [Fork] this library and submit a [Pull Request]. Make sure to click `compare across forks` to see your fork.

## ⚖️ License

This project is under the MIT License. See the [License] file for the full license text.

## 📜 Changes

Check out our [Changelog]

## 👍🏻 Code of Conduct

Please read my [Code of Conduct] before contributing or engaging in discussions.

<!-- Links -->
[LICENSE]: https://github.com/ggoodwin/go-nyse-stocks/blob/master/LICENSE.md
[CHANGELOG]: https://github.com/ggoodwin/go-nyse-stocks/blob/master/CHANGELOG.md
[SECURITY]: https://github.com/ggoodwin/go-nyse-stocks/blob/master/SECURITY.md
[FORK]: https://github.com/ggoodwin/go-nyse-stocks/fork
[PULL REQUEST]: https://github.com/ggoodwin/go-nyse-stocks/compare
[CODE OF CONDUCT]: https://github.com/ggoodwin/go-nyse-stocks/blob/master/CODE_OF_CONDUCT.md
[CONTRIBUTING]: https://github.com/ggoodwin/go-nyse-stocks/blob/master/CONTRIBUTING.md
[GITHUB ISSUES]: https://github.com/ggoodwin/go-nyse-stocks/issues
[YAHOO API]: https://finance.yahoo.com/most-active
[GO]: https://go.dev/
[DOWNLOAD GO]: https://go.dev/dl/
