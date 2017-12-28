package backoff

import (
	"math"
	"math/rand"
	"time"
)

func (b *Backoff) Do() time.Duration {
	b.Touch++
	return b.getTime()
}

func (b *Backoff) getTime() time.Duration {
	timeMin := b.MinTime
	timeMax := b.MaxTime
	if timeMin <= 0 {
		timeMin = 500 * time.Millisecond
	}
	if timeMax <= 0 {
		timeMax = 30 * time.Second
	}
	if timeMin >= timeMax {
		timeMax = timeMin + (1 * time.Second)
	}

	factorTime := float64(timeMin) * math.Pow(2, float64(b.Touch))

	duration := randInt64(int64(timeMin), int64(factorTime))

	if duration > int64(timeMax) {
		return timeMax
	}
	return time.Duration(duration)
}

func randInt64(min int64, max int64) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Int63n(max-min)
}

func (b *Backoff) Reset() {
	b.Touch = 0
}

func (b *Backoff) getTotalTouch() int {
	return b.Touch
}
