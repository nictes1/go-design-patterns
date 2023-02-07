package main

import "fmt"

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (b *NotificationBuilder) SetTitle(title string) {
	b.Title = title
}

func (b *NotificationBuilder) SetSubtitle(subtitle string) {
	b.Subtitle = subtitle
}

func (b *NotificationBuilder) SetMessage(message string) {
	b.Message = message
}

func (b *NotificationBuilder) SetImage(image string) {
	b.Image = image
}

func (b *NotificationBuilder) SetIcon(icon string) {
	b.Icon = icon
}

func (b *NotificationBuilder) SetPriority(priority int) {
	b.Priority = priority
}

func (b *NotificationBuilder) SetType(notType string) {
	b.NotType = notType
}

func (b *NotificationBuilder) Build() (*Notification, error) {
	if b.Icon != "" && b.Subtitle == "" {
		return nil, fmt.Errorf("You need to specify a subtitle when using an icon")
	}
	if b.Priority > 5 {
		return nil, fmt.Errorf("Priority must be 0 to 5")
	}

	return &Notification{
		title:    b.Title,
		subtitle: b.Subtitle,
		message:  b.Message,
		image:    b.Image,
		icon:     b.Icon,
		priority: b.Priority,
		notType:  b.NotType,
	}, nil
}
