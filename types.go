package jaen

import "time"

type Backoff struct {
	Touch   int
	MaxTime time.Duration
	MinTime time.Duration
}
