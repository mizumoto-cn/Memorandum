package main

import (
	"fmt"
	"time"

	"github.com/mizumoto-cn/memorandum"
	"github.com/mizumoto-cn/memorandum/notify"
	"github.com/mizumoto-cn/memorandum/schedule"
)

func main() {

	// Schedule an event
	format := time.RFC1123Z
	s, _ := schedule.NewSchedule("Local", format, time.Now().Format(format))

	// Repeat it every 5 seconds
	r := schedule.IntervalMultiplexer{
		Interval:      time.Duration(5) * time.Second,
		NumberOfTimes: 3,
	}
	r.Repeat(s)
	// Create the reminder
	rem := memorandum.Reminder{
		Schedule: s,
		Notifier: &notify.DesktopNotifier{},
	}

	err := rem.Remind("Reminder", "I'm a reminder made in Go!")
	fmt.Println(err)
}
