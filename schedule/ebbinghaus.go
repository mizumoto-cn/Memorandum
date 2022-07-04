package schedule

import (
	"time"
)

// EbbinghausMultiplexer is a multiplexer that repeats a task
// following the Ebbinghaus curve.
// Here we set it to [1, 2, 4, 7, 15] days later.
var (
	eb_interval_default = []int{1, 2, 4, 7, 15}
	eb_times_default    = uint32(5)
)

type EbbinghausMultiplexer struct {
	Interval      []int  // default = eb_interval_default
	NumberOfTimes uint32 // default = eb_times_default
}

func NewEbbinghausMultiplexer(interval []int, times uint32) *EbbinghausMultiplexer {
	if len(interval) == 0 {
		interval = eb_interval_default
	}
	if times == 0 {
		times = eb_times_default
	}
	return &EbbinghausMultiplexer{
		Interval:      interval,
		NumberOfTimes: times,
	}
}

func (e *EbbinghausMultiplexer) Repeat(schedule *Schedule) error {
	if e.NumberOfTimes == 0 {
		return ErrZeroNumberOfTimes
	}

	// to be refactored
	oldsc := schedule.Schedule
	length := len(oldsc)
	newsc := make([]time.Time, length*int(e.NumberOfTimes+1))
	copy(newsc, oldsc)
	day := time.Duration(24) * time.Hour
	for i := length; i < length*int(e.NumberOfTimes+1); i++ {
		newsc[i] = newsc[i-length].Add(time.Duration(e.Interval[i%len(e.Interval)]) * day)
	}
	schedule.CreateSchedule(newsc...)
	return nil
}
