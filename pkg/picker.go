package picker

// TODO in order to not mess with randomness, this should use its own instance of random number generator.

import (
	"math/rand"
	"time"
)

// Color represents a single RGB color
type Color struct {
	Red, Green, Blue int
}

// Colorset uses a
type Colorset struct {
}

// RandomColor generates a truly random color with 3 random numbers 0-255.
func RandomColor() Color {
	return Color{
		Red:   rand.Intn(256),
		Green: rand.Intn(256),
		Blue:  rand.Intn(256),
	}
}

// RandomColors generates a "cloud" in 3D color-space of colors.
func RandomColorset() {
	return Colorset{}
}

func init() {
	// TODO I guess maybe we should leave this to the user
	rand.Seed(time.Now().UTC().UnixNano())
}
