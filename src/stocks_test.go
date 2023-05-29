// Package Stocks is declaring that this Go file belongs to the `stocks` package, which can be
// imported and used in other Go files.
package stocks

import (
	"testing"
)

// TestGetFullDetails is a unit test function in Go that tests the GetFullDetails function by asserting expected
// values.
func TestGetFullDetails(t *testing.T) {

	// Call the function being tested
	result := GetFullDetails("AAPL")

	// Assert the expected values
	if result == nil {
		t.Error("Expected a non-nil result, but got nil.")
		return
	}
	if result.Symbol != "AAPL" {
		t.Errorf("Expected symbol to be 'AAPL', but got '%s'.", result.Symbol)
	}
	if result.ShortName != "Apple Inc." {
		t.Errorf("Expected companyName to be 'Apple Inc.', but got '%s'.", result.ShortName)
	}
}

// func BenchmarkGetFullDetails(b *testing.B) {
// 	result := GetFullDetails("AAPL")
// 	if result == nil {
// 		b.Error("Expected a non-nil result, but got nil.")
// 		return
// 	}
// 	if result.Symbol != "AAPL" {
// 		b.Errorf("Expected symbol to be 'AAPL', but got '%s'.", result.Symbol)
// 	}
// 	if result.ShortName != "Apple Inc." {
// 		b.Errorf("Expected companyName to be 'Apple Inc.', but got '%s'.", result.ShortName)
// 	}
// }
