// Package space is a tool to be used for calculating
// the age of a person on different planets in our
// solar system.
package space

const (
	// Seconds in one year on Earth
	EarthSeconds = float64(31557600)
)

type Planet string

var (
	// planets is a mapping of seconds per year for a
	// given planet
	planets = map[Planet]float64{
		"Earth":   EarthSeconds,
		"Mercury": 0.2408467 * EarthSeconds,
		"Venus":   0.61519726 * EarthSeconds,
		"Mars":    1.8808158 * EarthSeconds,
		"Jupiter": 11.862615 * EarthSeconds,
		"Saturn":  29.447498 * EarthSeconds,
		"Uranus":  84.016846 * EarthSeconds,
		"Neptune": 164.79132 * EarthSeconds,
	}
)

// Age takes in seconds and a Planet and returns the age of
// in years on the given planet.
func Age(seconds float64, planet Planet) float64 {
	planetYear := planets[planet]
	age := seconds / planetYear
	return age
}
