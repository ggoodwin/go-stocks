package stocks

import "time"

/** GetESTTime
 * * Gets EST time based on EST or EDT time for stock market
 */
func GetESTTime() time.Time {
	//Set the location for EST
	loc, _ := time.LoadLocation("EST")

	//Get current time in EST
	estTime := time.Now().In(loc)
	//fmt.Println("EST Time:", estTime)

	//Check if DST is in effect
	_, offset := estTime.Zone()
	if offset == -4*60*60 {
		//DST is in effect, subtract 4 hours
		estTime = estTime.Add(-4 * time.Hour)
	} else {
		//DST is not in effect, subtract 5 hours
		estTime = estTime.Add(-5 * time.Hour)
	}

	//Convert EST time to UTC
	//utcTime := estTime.UTC()
	return estTime
}

func IsMarketOpen() bool {
	currentTime := GetESTTime()
	openTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 9, 30, 0, 0, time.UTC)
	closeTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 16, 00, 0, 0, time.UTC)
	earlyClose := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 13, 00, 0, 0, time.UTC)
	dayOfWeek := currentTime.Weekday()

	// Check Holidays
	if IsHoliday() {
		return false
	}

	// Check Early Close Days
	if IsEarlyClose() && currentTime.After(openTime) && currentTime.Before(earlyClose) && dayOfWeek != time.Saturday && dayOfWeek != time.Sunday {
		return true
	}

	// Check Regular Market Hours
	if currentTime.After(openTime) && currentTime.Before(closeTime) && dayOfWeek != time.Saturday && dayOfWeek != time.Sunday {
		return true
	}

	return false
}
