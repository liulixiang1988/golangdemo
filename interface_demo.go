package main

import (
	"fmt"
)

type empty interface {
}

type PhoneConnector struct {
	name string
}

func main() {
	phoneConnector := PhoneConnector{name: "Phone Connector"}
	Disconnect(phoneConnector)
}

func Disconnect(usb empty) {
	if phoneConnector, ok := usb.(PhoneConnector); ok {
		fmt.Println("Phone disconnect:", phoneConnector.name)
	}

	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Phone disconnect:", v.name)
	default:
		fmt.Println("Unknow device")
	}
}
