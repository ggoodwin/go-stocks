<div align="center">
	<h1><img alt="nyse-stocks logo" src="https://github.com/octolibs/nyse-stocks/blob/main/stock.png" height="300" /><br />
		New York Stock Exchange (NYSE) Go Library
	</h1>

[![Go Reference](https://pkg.go.dev/badge/octolibs/nyse-stocks.svg)](https://pkg.go.dev/github.com/octolibs/nyse-stocks) [![Go Version](https://img.shields.io/github/go-mod/go-version/octolibs/nyse-stocks)](https://go.dev/) [![GoReportCard](https://goreportcard.com/badge/github.com/octolibs/nyse-stocks)](https://goreportcard.com/report/github.com/octolibs/nyse-stocks) [![CodeFactor](https://www.codefactor.io/repository/github/octolibs/nyse-stocks/badge)](https://www.codefactor.io/repository/github/octolibs/nyse-stocks) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/octolibs/nyse-stocks/.github/workflows/go.yml)](https://github.com/octolibs/nyse-stocks/blob/main/.github/workflows/go.yml) ![Size](https://img.shields.io/github/languages/code-size/octolibs/nyse-stocks) [![Last Commit](https://img.shields.io/github/last-commit/octolibs/nyse-stocks)](https://github.com/octolibs/nyse-stocks/commits/main) [![License](https://img.shields.io/github/license/octolibs/nyse-stocks)](https://github.com/octolibs/nyse-stocks/blob/main/LICENSE)

</div>
<hr/>

## 🌟 How it works

We use the [`Yahoo`](https://finance.yahoo.com/most-active) API to gather real time stock data.

## 📦 Installation and Usage

### Go

Make sure you have `Go` installed on your machine.

You can check by running the following command in the `console`

```plain
go version
```

If you don't have `Go` installed, you can download it from [here](https://go.dev/dl/).

### Add to your project

Run the following command in the `console`, in the `project directory`, to install the library with `go get`

```plain
go get github.com/octolibs/nyse-stocks
```

### Importing

Add the import to your `.go` file

```go
import "github.com/octolibs/nyse-stocks"
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

- [`Go`](https://go.dev/)

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report an issue. If you'd
like, we welcome any contributions. Fork this library and submit a pull
request.

## ⚖️ License

This project is under the MIT License. See the [LICENSE](https://github.com/octolibs/nyse-stocks/blob/main/LICENSE) file for the full license text.

## 📜 Changes

Check out our [CHANGELOG](https://github.com/octolibs/nyse-stocks/blob/main/CHANGELOG.md)
