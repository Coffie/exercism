// Package leap is used to check if a year is a leap year
package leap

// IsLeapYear takes a year as input and check whether it is
// a leap year or not.
func IsLeapYear(year int) bool {
	if year % 100 == 0 && year % 400 != 0 {
		return false
	}
	if year % 4 == 0 {
		return true
	}
	return false
}
