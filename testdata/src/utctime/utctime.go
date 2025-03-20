//nolint:revive,staticcheck // ignore for analysis testing.
package utctime

import "time"

func ValidUsage() {
	_ = time.Now().UTC()            // OK
	t := time.Now().UTC()           // OK
	someFunc(time.Now().UTC())      // OK
	if time.Now().UTC().Before(t) { // OK
	}
}

func InvalidUsage() {
	_ = time.Now()            // want "time.Now\\(\\) must be followed by .UTC\\(\\)"
	t := time.Now()           // want "time.Now\\(\\) must be followed by .UTC\\(\\)"
	someFunc(time.Now())      // want "time.Now\\(\\) must be followed by .UTC\\(\\)"
	if time.Now().Before(t) { // want "time.Now\\(\\) must be followed by .UTC\\(\\)"
	}
}

func someFunc(_ time.Time) {
}
