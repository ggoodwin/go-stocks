// Package Stocks is declaring that this Go file belongs to the `stocks` package, which can be
// imported and used in other Go files.
package stocks_test

// TestGetPriceAndPercentage is a unit test function that tests the GetPriceAndPercentage function in the stocks package and
// checks if it returns expected values.
// func TestGetPriceAndPercentage(t *testing.T) {

// 	// Call the function being tested
// 	price, percent, direction := stocks.GetPriceAndPercentage("AAPL")

// 	// Assert the expected values
// 	if price == "" {
// 		t.Errorf("Expected a result from price, but got an empty result.")
// 	}
// 	if percent == "" {
// 		t.Errorf("Expected a result from percent, but got an empty result.")
// 	}
// 	if direction == "" {
// 		t.Errorf("Expected a result from direction, but got an empty result.")
// 	}
// }

// TestGetFullDetails is a unit test function in Go that tests the GetFullDetails function by asserting expected
// values.
// func TestGetFullDetails(t *testing.T) {

// 	// Call the function being tested
// 	result := stocks.GetFullDetails("AAPL")

// 	// Assert the expected values
// 	if result == nil {
// 		t.Error("Expected a non-nil result, but got nil.")
// 		return
// 	}
// 	if result.Symbol != "AAPL" {
// 		t.Errorf("Expected symbol to be 'AAPL', but got '%s'.", result.Symbol)
// 	}
// 	if result.ShortName != "Apple Inc." {
// 		t.Errorf("Expected companyName to be 'Apple Inc.', but got '%s'.", result.ShortName)
// 	}
// }
