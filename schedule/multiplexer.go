package schedule

// Interface multiplexer defines how to repeat a task.
// Either on certain days of week or fixed intervals.
// Ebbinghaus curve shall be implemented later.
type Multiplexer interface {
	Repeat(schedule *Schedule) error
}

// assert implementation check
var _ Multiplexer = &WeeklyMultiplexer{}

var _ Multiplexer = &IntervalMultiplexer{}

var _ Multiplexer = &EbbinghausMultiplexer{}
