package stocks

import (
	"time"
)

func IsEarlyClose() bool {
	currentDate := GetESTTime()

	/**
	 * * Same Day Holidays
	 */
	// Independence Day
	if currentDate.Month() == time.July && currentDate.Day() == 4 {
		return true
	}

	if currentDate.Year() == 2023 {
		//Thanksgiving
		if currentDate.Month() == time.November && currentDate.Day() == 23 {
			return true
		}
	} else if currentDate.Year() == 2024 {
		//Thanksgiving
		if currentDate.Month() == time.November && currentDate.Day() == 28 {
			return true
			//Christmas
		} else if currentDate.Month() == time.December && currentDate.Day() == 25 {
			return true
		}
	} else if currentDate.Year() == 2025 {
		//Thanksgiving
		if currentDate.Month() == time.November && currentDate.Day() == 27 {
			return true
			//Christmas
		} else if currentDate.Month() == time.December && currentDate.Day() == 25 {
			return true
		}
	}
	return false
}

func IsHoliday() bool {
	currentDate := GetESTTime()

	/**
	 * * Same Day Holidays
	 */
	// Juneteenth
	if currentDate.Month() == time.June && currentDate.Day() == 19 {
		return true
	}

	if currentDate.Year() == 2023 {
		// New Years
		if currentDate.Month() == time.January && currentDate.Day() == 2 {
			return true
			// MLK Jr. Day
		} else if currentDate.Month() == time.January && currentDate.Day() == 16 {
			return true
			// Washington's Birthday
		} else if currentDate.Month() == time.February && currentDate.Day() == 20 {
			return true
			// Good Friday
		} else if currentDate.Month() == time.April && currentDate.Day() == 7 {
			return true
			// Memorial Day
		} else if currentDate.Month() == time.May && currentDate.Day() == 29 {
			return true
			// Labor Day
		} else if currentDate.Month() == time.September && currentDate.Day() == 4 {
			return true
			// Christmas
		} else if currentDate.Month() == time.December && currentDate.Day() == 25 {
			return true
		}
	} else if currentDate.Year() == 2024 {
		// New Years
		if currentDate.Month() == time.January && currentDate.Day() == 1 {
			return true
			// MLK Jr. Day
		} else if currentDate.Month() == time.January && currentDate.Day() == 15 {
			return true
			// Washington's Birthday
		} else if currentDate.Month() == time.February && currentDate.Day() == 19 {
			return true
			// Good Friday
		} else if currentDate.Month() == time.March && currentDate.Day() == 29 {
			return true
			// Memorial Day
		} else if currentDate.Month() == time.May && currentDate.Day() == 27 {
			return true
			// Labor Day
		} else if currentDate.Month() == time.September && currentDate.Day() == 2 {
			return true
		}
	} else if currentDate.Year() == 2025 {
		// New Years
		if currentDate.Month() == time.January && currentDate.Day() == 1 {
			return true
			// MLK Jr. Day
		} else if currentDate.Month() == time.January && currentDate.Day() == 20 {
			return true
			// Washington's Birthday
		} else if currentDate.Month() == time.February && currentDate.Day() == 17 {
			return true
			// Good Friday
		} else if currentDate.Month() == time.April && currentDate.Day() == 18 {
			return true
			// Memorial Day
		} else if currentDate.Month() == time.May && currentDate.Day() == 26 {
			return true
			// Labor Day
		} else if currentDate.Month() == time.September && currentDate.Day() == 1 {
			return true
		}
	}
	return false
}
