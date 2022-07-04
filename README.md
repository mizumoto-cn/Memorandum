# Memorandum

Memorandum is a library with which you can set up your daily schedule, and get notified on Windows desktop through `gen2brain/beeep`.

## Quick Start

Take a glance at `/demo_main/demo_main.go`

try this:

```bash
go mod init demo_main
go get github.com/mizumoto-cn/memorandum
```

create a `main.go` file in the directory:

<!-- markdownlint-disable MD010 -->

```golang
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
```

```bash
go run ./main.go
```

and you'll receive three notifications on your desktop every 5 seconds.

## Todo

- Push Bullet demo to be added.
- Concurrency
- GUI

## Credits

<https://github.com/gen2brain/beeep>
<https://github.com/cschomburg/go-pushbullet>

## License

[`MGPL v1.2`](License/Mizumoto%20General%20Public%20License%20v1.2.md) as usual.
