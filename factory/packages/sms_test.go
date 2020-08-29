package packages

import (
	"fmt"
	"log"
	"testing"
)

func TestSMSSendNotification(t *testing.T) {
	notif := Notification{
		To:      "00193116497",
		Message: "Hi",
	}

	notifier, err := GetNotifier(SMS)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = notifier.SendNotification(notif)
	if err != nil {
		t.Error(err.Error())
	}

}

// ExampleSMSSendNotification func
func ExampleSMSSendNotification() {
	notif := Notification{
		To:      "00193116497",
		Message: "Hi",
	}

	notifier, err := GetNotifier(SMS)
	if err != nil {
		log.Fatalln(err.Error())
	}

	str, err := notifier.SendNotification(notif)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Print(str)
	// Output:
	// `Hi` was sent to number `00193116497` successfully!
}
