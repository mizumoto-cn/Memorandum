package notify

import (
	"github.com/gen2brain/beeep"
	"github.com/mizumoto-cn/memorandum/pushbullet"
)

// Notifier is a notification interface.
// To be updated: consider github.com/nikosr/notify
type Notifier interface {
	Notify(title, message string) error
}

// DesktopNotifier wraps beeep
type DesktopNotifier struct{}

// Notify creates a beeep.Notify
func (n *DesktopNotifier) Notify(title, message string) error {
	return beeep.Notify(title, message, "arc/1.png")
}

// PushbulletNotifier contains necessary information for pb API
type PushbulletNotifier struct {
	client     *pushbullet.Client
	deviceTags []string
}

// NewPushbulletNotifier creates a new PushbulletNotifier
// https://docs.pushbullet.com/#api-overview
func NewPushbulletNotifier(apiKey string, deviceTags ...string) *PushbulletNotifier {
	return &PushbulletNotifier{
		client:     pushbullet.New(apiKey),
		deviceTags: deviceTags,
	}
}

// Notify will push notifications to specified receivers.
func (pb *PushbulletNotifier) Notify(title, message string) error {
	for _, tag := range pb.deviceTags {
		device, err := pb.client.Device(tag)
		if err != nil {
			return err
		}
		err = device.PushNote(title, message)
		if err != nil {
			return err
		}
	}
	return nil
}
