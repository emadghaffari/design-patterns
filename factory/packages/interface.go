package packages

import (
	"fmt"
)

// SMS, Email const for factory design-pattern
const (
	SMS   = 1
	Email = 2
)

// Notification struct
type Notification struct {
	To      string
	Message string
}

// Notifier interface
type Notifier interface {
	SendNotification(Notification) (string, error)
}

// GetNotifier func
// for choice between SMS and Email
func GetNotifier(nType int) (Notifier, error) {
	switch nType {
	case SMS:
		return new(SMSNotification), nil
	case Email:
		return new(EmailNotification), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Notifier type %d not recognized.", nType))
	}
}
