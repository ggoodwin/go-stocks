# Go Stocks Library

[![GoReportCard](https://goreportcard.com/badge/github.com/octolibs/stocks)](https://goreportcard.com/report/github.com/octolibs/stocks)

## How it works

We use the `Yahoo` stock API.

## Installation and Usage

### Installing

Run the following command in the `console`, in the `project directory`, to install the library with `go get`

```plain
go get github.com/octolibs/stocks
```

### Importing

```go
import "github.com/octolibs/stocks"
```

#### Stock Price and Percentage Change

```go
//Returns `float64` numbers
price, percent := stocks.GetPriceAndPercentage("AAPL")
```

## Issues and Contributing

If you find an issue with this library, please report an issue. If you'd
like, we welcome any contributions. Fork this library and submit a pull
request.
