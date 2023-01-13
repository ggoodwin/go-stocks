# Go Stocks Library

[![Go Version](https://img.shields.io/github/go-mod/go-version/octolibs/stocks)](https://go.dev/) [![GoReportCard](https://goreportcard.com/badge/github.com/octolibs/stocks)](https://goreportcard.com/report/github.com/octolibs/stocks) [![CodeFactor](https://www.codefactor.io/repository/github/octolibs/stocks/badge)](https://www.codefactor.io/repository/github/octolibs/stocks) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/octolibs/stocks/.github/workflows/go.yml)](https://github.com/octolibs/stocks/blob/main/.github/workflows/go.yml) ![Size](https://img.shields.io/github/languages/code-size/octolibs/stocks) [![Last Commit](https://img.shields.io/github/last-commit/octolibs/stocks)](https://github.com/octolibs/stocks/commits/main) [![License](https://img.shields.io/github/license/octolibs/stocks)](https://github.com/octolibs/stocks/blob/main/LICENSE)

## OctoLibs

![octolibs-brand](https://avatars.githubusercontent.com/u/122575388?s=400&u=6803206c4e7618aa7c24a7dab7384947d5d0b29c&v=4)

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
