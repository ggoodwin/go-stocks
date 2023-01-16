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

func IsMarketOpen() bool {
	currentTime := GetESTTime()
	openTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 9, 30, 0, 0, time.UTC)
	closeTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 16, 00, 0, 0, time.UTC)
	earlyClose := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 13, 00, 0, 0, time.UTC)
	dayOfWeek := currentTime.Weekday()

	// Check Holidays
	if IsHoliday() == true {
		return false
	}

	// Check Early Close Days
	if IsEarlyClose() == true {
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
