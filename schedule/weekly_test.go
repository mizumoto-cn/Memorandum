package schedule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getSchedule() *Schedule {
	layout := time.RFC1123Z
	value := time.Now().Format(layout)

	sc, _ := NewSchedule("Asia/Tokyo", layout, value)
	return sc
}

func setupWeekly(n uint32, weekdays ...time.Weekday) (*Schedule, Multiplexer) {
	sc := getSchedule()
	m := &WeeklyMultiplexer{
		Days:          weekdays,
		NumberOfTimes: n,
	}
	return sc, m
}

func assertExpectedRepeatNum(s *Schedule, m Multiplexer, n uint32, t *testing.T) {
	l := s.Len()
	err := m.Repeat(s)

	assert.NoError(t, err)
	assert.Equal(t, l*int((n+1)), s.Len())
}

func TestWeeklyCanRepeat(t *testing.T) {
	var n uint32 = 10
	var pres time.Time
	day := time.Monday
	s, m := setupWeekly(n, day)
	assertExpectedRepeatNum(s, m, n, t)

	i := 0
	for s.Len() > 0 {
		pres = s.Next()
		if i < 1 {
			i++
			continue
		}
		wd := pres.Weekday()
		assert.Equal(t, day, wd)
	}
}
