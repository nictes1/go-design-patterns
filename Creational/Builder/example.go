package main

import "fmt"

func main() {

	builder := NewNotificationBuilder()
	builder.SetTitle("New Notification")
	builder.SetIcon("icon-notification.png")
	builder.SetImage("image.png")
	builder.SetPriority(10)
	builder.SetMessage("This a notification message")
	builder.SetType("alert")

	notification, err := builder.Build()
	if err != nil {
		fmt.Println("Error building notification", err)
	} else {
		fmt.Println(notification)
	}

}
