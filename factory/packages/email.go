package packages

import "fmt"

// EmailNotification struct
type EmailNotification struct{}

// SendNotification metod
func (n EmailNotification) SendNotification(notification Notification) (string, error) {
	return fmt.Sprintf("`%s` was sent to email `%s` successfully!", notification.Message, notification.To), nil
}
