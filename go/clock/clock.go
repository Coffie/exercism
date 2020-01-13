// Package clock is a program that can convert hours and minutes
// to a clock format and perform addition and subtractions of
// minutes.
package clock

import "fmt"

// Clock represents time with hours and minutes in 24h format
type Clock struct {
	hour int
	minute int
}

// New converts hours and minutes to total minutes both limited to 24h/1440m
// and returns a Clock with 24h format
func New(hour, minute int) Clock {
	// Oneliner for total, not as readable
	//total := (1440 + ((hour % 24) * 60 + (minute % 1440)) % 1440) % 1440
	hoursInMins := (hour % 24) * 60
	minsAdjusted := minute % 1440
	sum := (hoursInMins + minsAdjusted) % 1440
	sumAdjusted := (1440 + sum) % 1440
	return Clock{hour: sumAdjusted / 60, minute: sumAdjusted % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add takes minutes and adds them to the corresponding clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute + minutes)
}

// Subtract takes minutes and subtracts them to the corresponding clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute - minutes)
}
