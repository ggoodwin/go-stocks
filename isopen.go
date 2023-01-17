package stocks

import "time"

/** GetESTTime
 * * Gets EST time based on EST or EDT time for stock market
 */
func GetESTTime() time.Time {
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)
	return now
}

/** IsMarketOpen
 * * Checks if the market is open
 */
func IsMarketOpen() bool {
	loc, _ := time.LoadLocation("America/New_York")
	currentTime := GetESTTime()
	openTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 9, 30, 0, 0, loc)
	closeTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 16, 00, 0, 0, loc)
	earlyClose := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 13, 00, 0, 0, loc)
	dayOfWeek := currentTime.Weekday()

	// Check Holidays
	// if IsHoliday() {
	// 	return false
	// }

	// Check Early Close Days
	if IsEarlyClose() {
		if currentTime.After(openTime) && currentTime.Before(earlyClose) && dayOfWeek != time.Saturday && dayOfWeek != time.Sunday {
			return true
		} else {
			return false
		}
	}

	// Check Regular Market Hours
	if currentTime.After(openTime) && currentTime.Before(closeTime) && dayOfWeek != time.Saturday && dayOfWeek != time.Sunday {
		return true
	}

	return false
}
