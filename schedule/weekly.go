package schedule

import (
	"time"
)

// WeeklyMultiplexer is a multiplexer that repeats a task on certain days of week.
type WeeklyMultiplexer struct {
	// Days on which the task shall be repeated.
	Days []time.Weekday

	// Number of times the task shall be repeated.
	NumberOfTimes uint32
}

func (m *WeeklyMultiplexer) Repeat(schedule *Schedule) error {
	if m.NumberOfTimes == 0 {
		return ErrZeroNumberOfTimes
	}

	// to be refactored
	oldsc := schedule.Schedule
	length := len(oldsc)
	newsc := make([]time.Time, length*int(m.NumberOfTimes+1))
	copy(newsc, oldsc)

	day := time.Duration(24) * time.Hour
	weekdayMap := m.weekdayMap()

	for i := length; i < length*int(m.NumberOfTimes+1); i++ {
		t := newsc[i-length].Add(day)
		newsc[i] = t.Add(weekdayMap[int(t.Weekday())])
	}
	schedule.CreateSchedule(newsc...)
	return nil
}

func (m *WeeklyMultiplexer) weekdayMap() map[int]time.Duration {
	weekdays := [7]int{}

	// Set the days of week on which tasks are set's value to 1.
	for _, v := range m.Days {
		weekdays[v] = 1
	}

	weekdaysMap := make(map[int]time.Duration)

	for i, v := range weekdays {
		if v == 1 {
			weekdaysMap[i] = 0
		} else {
			j := 1
			for weekdays[(i+j)%7] == 0 {
				j++
			}
			weekdaysMap[i] = time.Duration(j*24) * time.Hour
		}
	}
	return weekdaysMap
}
