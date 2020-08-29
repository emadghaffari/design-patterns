package packages

import (
	"fmt"
	"log"
	"testing"
)

func TestEmailSendNotification(t *testing.T) {
	notif := Notification{
		To:      "emadghaffariii@gmail.com",
		Message: "Hi",
	}

	notifier, err := GetNotifier(Email)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = notifier.SendNotification(notif)
	if err != nil {
		t.Error(err.Error())
	}

}

// ExampleEmailSendNotification func
func ExampleEmailSendNotification() {
	notif := Notification{
		To:      "emadghaffariii@gmail.com",
		Message: "Hi",
	}

	notifier, err := GetNotifier(Email)
	if err != nil {
		log.Fatalln(err.Error())
	}

	str, err := notifier.SendNotification(notif)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Print(str)
	// Output:
	// `Hi` was sent to email `emadghaffariii@gmail.com` successfully!
}
