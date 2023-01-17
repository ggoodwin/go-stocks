<div align="center">
	<h1><img alt="Stocks logo" src="https://avatars.githubusercontent.com/u/122575388?s=400&u=6803206c4e7618aa7c24a7dab7384947d5d0b29c&v=4" height="300" /><br />
		Go Stocks Library
	</h1>

[![Go Reference](https://pkg.go.dev/badge/octolibs/stocks.svg)](https://pkg.go.dev/github.com/octolibs/stocks) [![Go Version](https://img.shields.io/github/go-mod/go-version/octolibs/stocks)](https://go.dev/) [![GoReportCard](https://goreportcard.com/badge/github.com/octolibs/stocks)](https://goreportcard.com/report/github.com/octolibs/stocks) [![CodeFactor](https://www.codefactor.io/repository/github/octolibs/stocks/badge)](https://www.codefactor.io/repository/github/octolibs/stocks) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/octolibs/stocks/.github/workflows/go.yml)](https://github.com/octolibs/stocks/blob/main/.github/workflows/go.yml) ![Size](https://img.shields.io/github/languages/code-size/octolibs/stocks) [![Last Commit](https://img.shields.io/github/last-commit/octolibs/stocks)](https://github.com/octolibs/stocks/commits/main) [![License](https://img.shields.io/github/license/octolibs/stocks)](https://github.com/octolibs/stocks/blob/main/LICENSE)

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
go get github.com/octolibs/stocks
```

### Importing

Add the import to your `.go` file

```go
import "github.com/octolibs/stocks"
```

## 💰 Usage

### Stock Price and Percentage Change

Add this to your `.go` file

```go
//Returns strings ex: `$123.45`, `1.23%`, `↑`
price, percent, direction := stocks.GetPriceAndPercentage("AAPL")
```

### Get full stock details

```go
//Returns `Result` struct
stock := stocks.GetFullDetails("AAPL")
```

### Is Market Open?

```go
//Returns `bool`
isOpen := stocks.IsMarketOpen()
```

### Is Holiday?

```go
//Returns `bool`
isHoliday := stocks.IsHoliday()
```

### Is Early Close?

```go
//Returns `bool`
isEarlyClose := stocks.IsEarlyClose()
```

## 💻 Dependencies

- [`Go`](https://go.dev/)

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report an issue. If you'd
like, we welcome any contributions. Fork this library and submit a pull
request.

## ⚖️ License

This project is under the MIT License. See the [LICENSE](https://github.com/octolibs/stocks/blob/main/LICENSE) file for the full license text.

## 📜 Changes

Check out our [CHANGELOG](https://github.com/octolibs/stocks/blob/main/CHANGELOG.md)
