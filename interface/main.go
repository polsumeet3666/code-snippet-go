package main

import "fmt"

// User struct
type User struct {
	Name     string
	Email    string
	Notifier UserNotifier
}

// Notify implented in struct
func (user *User) Notify(message string) {
	user.Notifier.SendMessage(user, message)
}

// UserNotifier interface
type UserNotifier interface {
	SendMessage(user *User, message string) error
}

//EmailNotifier struct
type EmailNotifier struct {
}

// SendMessage implement interface by emailNotifier
func (emailNotifier EmailNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Email Notification for %s and message : %s\n", user.Name, message)
	return err
}

//SMSNotifier struct
type SMSNotifier struct {
}

// SendMessage implemented by sms notifier
func (notifier SMSNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("SMS Notification for %s and message : %s\n", user.Name, message)
	return err
}

func main() {
	user1 := &User{Name: "sumeet1", Email: "sumeet1@emai.com", Notifier: EmailNotifier{}}
	user2 := &User{Name: "sumeet2", Email: "sumeet2@emai.com", Notifier: SMSNotifier{}}

	user1.Notify("h1")
	user2.Notify("h2")
}
