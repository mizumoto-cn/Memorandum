package memorandum

import (
	"time"

	"github.com/mizumoto-cn/memorandum/notify"
	"github.com/mizumoto-cn/memorandum/schedule"
)

type Reminder struct {
	Schedule *schedule.Schedule
	Notifier notify.Notifier
}

// Remind() will go through the schedule and notify the notifier.
func (r *Reminder) Remind(title, msg string) error {
	s := *r.Schedule
	for s.Len() > 0 {
		if d := time.Until(s.Next()); d > time.Duration(0) {
			time.Sleep(d)
		}
		err := r.Notifier.Notify(title, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
