package main

import (
	"fmt"

	"github.com/emadghaffari/design-patterns/factory/packages"
)

func main() {

	if notifire, err := packages.GetNotifier(packages.Email); err == nil {
		notif := packages.Notification{
			To:      "emadghaffariii@gmail.com",
			Message: "Hi, this is a test for factory design-pattern",
		}
		resultStr, _ := notifire.SendNotification(notif)
		fmt.Println(resultStr)
	}

	if notifire, err := packages.GetNotifier(packages.SMS); err == nil {
		notif2 := packages.Notification{
			To:      "001664973115",
			Message: "Hi!",
		}
		resultStr, _ := notifire.SendNotification(notif2)
		fmt.Println(resultStr)
	}

}
