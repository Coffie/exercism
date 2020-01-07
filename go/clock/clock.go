package clock

import "fmt"

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

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute + minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute - minutes)
}
