package schedule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func getSchedule() *Schedule {
// 	layout := time.RFC1123Z
// 	value := time.Now().Format(layout)
// 	sc, _ := NewSchedule("Asia/Tokyo", layout, value)
// 	return sc
// }

func setupInterval(n uint32, interval time.Duration) (*Schedule, Multiplexer) {
	sc := getSchedule()
	m := &IntervalMultiplexer{
		Interval:      interval,
		NumberOfTimes: n,
	}
	return sc, m
}

// assertExpectedRepeatNum asserts that the number of times is correct.

func TestIntervalCanRepeat(t *testing.T) {
	var n uint32 = 10
	interval := time.Duration(5) * time.Minute
	s, m := setupInterval(n, interval)
	pres := s.Next()
	assertExpectedRepeatNum(s, m, n, t)

	for s.Len() > 0 {
		next := s.Next()
		assert.Equal(t, interval, next.Sub(pres))
		pres = next
	}
}
