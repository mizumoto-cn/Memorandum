package schedule

import (
	"container/heap"
	"time"
)

type Schedule struct {
	Schedule []time.Time
}

// assert that Schedule implements heap.Interface
var _ heap.Interface = &Schedule{}

// Implement heap.Interface
func (s *Schedule) Len() int {
	return len(s.Schedule)
}

func (s *Schedule) Less(i, j int) bool {
	return s.Schedule[i].Before(s.Schedule[j])
}

func (s *Schedule) Swap(i, j int) {
	s.Schedule[i], s.Schedule[j] = s.Schedule[j], s.Schedule[i]
}

// Push -> heap.Append
func (s *Schedule) Push(x interface{}) {
	s.Schedule = append(s.Schedule, x.(time.Time))
}

// Pop -> heap.Pop
func (s *Schedule) Pop() interface{} {
	n := len(s.Schedule)
	x := s.Schedule[n-1]
	s.Schedule = s.Schedule[:n-1]
	return x
}

// NewSchedule creates a new Schedule
func NewSchedule(location, layout, value string) (*Schedule, error) {
	s := &Schedule{}
	err := s.SetTime(location, layout, value)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// SetTime sets a time to the Schedule
func (s *Schedule) SetTime(location, layout, value string) error {
	if location == "" {
		location = "Asia/Tokyo"
	}
	loc, err := time.LoadLocation(location)
	if err != nil {
		return err
	}
	if loc == nil {
		return ErrInvalidLocation
	}
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return err
	}
	s.CreateSchedule(t)
	return nil
}

// CreateSchedule creates a new Schedule
func (s *Schedule) CreateSchedule(params ...time.Time) {
	s.Schedule = params
	heap.Init(s)
}

// Next returns the time of the next event
func (s *Schedule) Next() time.Time {
	return heap.Pop(s).(time.Time)
}
