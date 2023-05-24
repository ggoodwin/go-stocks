# Getting Started

This page is dedicated to helping you get started with go-nyse-stocks.
Once you've done that please don't forget to submit it to the
[Awesome go-nyse-stocks] list :).

**First, lets cover a few topics so you can make the best choices on how to
move forward from here.**

## Requirements

go-nyse-stocks requires Go version 1.19 or higher.  It has been tested to compile and
run successfully on Debian Linux 8, FreeBSD 10, and Windows 7.  It is expected
that it should work anywhere Go 1.19 or higher works. If you run into problems
please let me know :).

You must already have a working Go environment setup to use go-nyse-stocks.  If you
are new to Go and have not yet installed and tested it on your computer then
please visit the [Go Install Page] first then I highly
recommend you walk though [A Tour of Go] to
help get your familiar with the Go language.  Also checkout the relevant Go plugin
for your editor &mdash; they are hugely helpful when developing Go code.

* Visual Studio &mdash; [vscode-go]
* Vim &mdash; [vim-go]
* Sublime &mdash; [GoSublime]

## Install go-nyse-stocks

Like any other Go package the fist step is to `go get` the package.  This will
always pull the latest tagged release from the master branch. Then run
`go install` to compile and install the libraries on your system.

### Windows

Simply run the following command in a command prompt or powershell instance to download, compile, and install the package.

Make sure you are in your project's directory before running the command.

```plain
go get github.com/ggoodwin/go-nyse-stocks
```

### Linux/BSD/MacOS

Run go get to download the package to your GOPATH/src folder.

```sh
go get github.com/ggoodwin/go-nyse-stocks
```

Finally, compile and install the package into the GOPATH/pkg folder. This isn't
absolutely required but doing this will allow the Go plugin for your editor to
provide autocomplete for all go-nyse-stocks functions.

```sh
cd $GOPATH/src/github.com/ggoodwin/go-nyse-stocks
go install
```

## Use go-nyse-stocks

### Importing

Add the import to your `.go` file

```go
import "github.com/ggoodwin/go-nyse-stocks"
```

### Commands

#### Stock Price and Percentage Change

Add this to your `.go` file

```go
// Parameter: Stock Symbol
// Returns: strings ex: `$123.45`, `1.23%`, `↑`
price, percent, direction := go_nyse_stocks.GetPriceAndPercentage("AAPL")
```

#### Get full stock details

```go
// Parameter: Stock Symbol
// Returns `Result` struct
stock := go_nyse_stocks.GetFullDetails("AAPL")
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

<!-- Links -->
[Awesome go-nyse-stocks]: https://github.com/ggoodwin/go-nyse-stocks/wiki/Awesome-go-nyse-stocks
[Go Install Page]: https://golang.org/doc/install
[A Tour of Go]: https://tour.golang.org/welcome/1
[vim-go]: https://github.com/fatih/vim-go
[GoSublime]: https://github.com/DisposaBoy/GoSublime
[vscode-go]: https://github.com/golang/vscode-go
