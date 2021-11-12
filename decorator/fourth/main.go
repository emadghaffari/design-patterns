package main

import "fmt"

type Notification interface {
	Send() string
}

type SMS struct {
	Notif Notification
}

func (s SMS) Send() string {
	return fmt.Sprintf("send sms: %s", s.Notif.Send())
}

type Slack struct {
	Notif Notification
}

func (s Slack) Send() string {
	return fmt.Sprintf("send slack: %s", s.Notif.Send())
}

type FaceBook struct {
	Notif Notification
}

func (s FaceBook) Send() string {
	return fmt.Sprintf("send facebook: %s", s.Notif.Send())
}

type Notifier struct {
	Msg string
}

func (n Notifier) Send() string {
	return n.Msg
}

func main() {
	notif := Notifier{Msg: "hello world"}

	m1 := SMS{Notif: notif}
	fmt.Println(m1.Send())

	m2 := FaceBook{Notif: m1}
	fmt.Println(m2.Send())

	m3 := Slack{Notif: m2}
	fmt.Println(m3.Send())
}
