<div align="center">
	<h1><img alt="stocks-go logo" src="https://github.com/ggoodwin/stocks-go/blob/master/nyse.png" height="300" /><br />
		Stock Info Go Library
	</h1>

<h2>Discord</h2>

[![GMan#0001](https://dcbadge.vercel.app/api/shield/179795086543028224)](https://discord.id/?prefill=179795086543028224) [![Go Package Devs](https://dcbadge.vercel.app/api/server/jwEYR2Dume)](https://discord.gg/jwEYR2Dume)

<h2>Go</h2>

[![Go Reference](https://pkg.go.dev/badge/ggoodwin/stocks-go.svg)](https://pkg.go.dev/github.com/ggoodwin/stocks-go) [![Go Version](https://img.shields.io/github/go-mod/go-version/ggoodwin/stocks-go)](https://go.dev/doc/go1.19)

<h2>Repo Info</h2>

![Size](https://img.shields.io/github/languages/code-size/ggoodwin/stocks-go) [![Last Commit](https://img.shields.io/github/last-commit/ggoodwin/stocks-go)](https://github.com/ggoodwin/stocks-go/commits/master) [![License](https://img.shields.io/github/license/ggoodwin/stocks-go)](https://github.com/ggoodwin/stocks-go/blob/master/LICENSE.md)

<h2>Code Reports</h2>

[![GoReportCard](https://goreportcard.com/badge/github.com/ggoodwin/stocks-go)](https://goreportcard.com/report/github.com/ggoodwin/stocks-go) [![CodeFactor](https://www.codefactor.io/repository/github/ggoodwin/stocks-go/badge)](https://www.codefactor.io/repository/github/ggoodwin/stocks-go) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/17f51d3e54264211b19220ce470783ae)](https://app.codacy.com/gh/ggoodwin/stocks-go/dashboard) [![Coverage Status](https://coveralls.io/repos/github/ggoodwin/stocks-go/badge.svg?branch=master)](https://coveralls.io/github/ggoodwin/stocks-go?branch=master)

<h2>Actions</h2>

[![Go](https://github.com/ggoodwin/stocks-go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/ggoodwin/stocks-go/actions/workflows/go.yml) [![lint](https://github.com/ggoodwin/stocks-go/actions/workflows/lint.yml/badge.svg?branch=master)](https://github.com/ggoodwin/stocks-go/actions/workflows/lint.yml) [![CodeQL](https://github.com/ggoodwin/stocks-go/actions/workflows/github-code-scanning/codeql/badge.svg?branch=master)](https://github.com/ggoodwin/stocks-go/actions/workflows/github-code-scanning/codeql) [![Vulnerabilities](https://github.com/ggoodwin/stocks-go/actions/workflows/vulnerabilities.yml/badge.svg?branch=master)](https://github.com/ggoodwin/stocks-go/actions/workflows/vulnerabilities.yml) [![Coveralls](https://github.com/ggoodwin/stocks-go/actions/workflows/coveralls.yml/badge.svg?branch=master)](https://github.com/ggoodwin/stocks-go/actions/workflows/coveralls.yml)

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
go get github.com/ggoodwin/stocks-go
```

### Importing

Add the import to your `.go` file

```go
import stocks "github.com/ggoodwin/stocks-go"
```

## 💰 Usage

### Stock Price and Percentage Change

Add this to your `.go` file

```go
// Parameter: Stock Symbol
// Returns: strings ex: `$123.45`, `1.23%`, `↑`
price, percent, direction := stocks.GetPriceAndPercentage("AAPL")
```

### Get full stock details

```go
// Parameter: Stock Symbol
// Returns `Result` struct
stock := stocks.GetFullDetails("AAPL")
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
[LICENSE]: https://github.com/ggoodwin/stocks-go/blob/master/LICENSE.md
[CHANGELOG]: https://github.com/ggoodwin/stocks-go/blob/master/CHANGELOG.md
[SECURITY]: https://github.com/ggoodwin/stocks-go/blob/master/SECURITY.md
[FORK]: https://github.com/ggoodwin/stocks-go/fork
[PULL REQUEST]: https://github.com/ggoodwin/stocks-go/compare
[CODE OF CONDUCT]: https://github.com/ggoodwin/stocks-go/blob/master/CODE_OF_CONDUCT.md
[CONTRIBUTING]: https://github.com/ggoodwin/stocks-go/blob/master/CONTRIBUTING.md
[GITHUB ISSUES]: https://github.com/ggoodwin/stocks-go/issues
[YAHOO API]: https://finance.yahoo.com/most-active
[GO]: https://go.dev/
[DOWNLOAD GO]: https://go.dev/dl/
