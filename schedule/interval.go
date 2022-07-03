// Interval Multiplexer
package schedule

import (
	"time"
)

type IntervalMultiplexer struct {
	Interval      time.Duration
	NumberOfTimes uint32
}

// Repeat take a schedule object
// and alter it into NumberOfTimes and Interval
func (m *IntervalMultiplexer) Repeat(schedule *Schedule) error {
	if m.NumberOfTimes == 0 {
		return ErrZeroNumberOfTimes
	}

	// to be refactored
	oldsc := schedule.Schedule
	length := len(oldsc)
	newsc := make([]time.Time, length*int(m.NumberOfTimes+1))
	copy(newsc, oldsc)

	for i := length; i < length*int(m.NumberOfTimes+1); i++ {
		newsc[i] = newsc[i-length].Add(m.Interval)
	}
	schedule.CreateSchedule(newsc...)
	return nil
}
