package main

import "fmt"

func main() {

	var builder = newNotificationBuilder()
	builder.SetTitle("New Notification")
	builder.SetIcon("icon-notification.png")
	builder.SetSubtitle("This a subtitle")
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
