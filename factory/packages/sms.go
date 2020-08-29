package packages

import "fmt"

// SMSNotification struct
type SMSNotification struct{}

// SendNotification metod
func (n SMSNotification) SendNotification(notification Notification) (string, error) {
	return fmt.Sprintf("`%s` was sent to number `%s` successfully!", notification.Message, notification.To), nil
}
