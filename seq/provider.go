package seq

import "time"

// Provider is a provider for values for a seq.
type Provider interface {
	Len() int
	GetValue(int) float64
}

// TimeProvider is a provider for values for a seq.
type TimeProvider interface {
	Len() int
	GetValue(int) time.Time
}
