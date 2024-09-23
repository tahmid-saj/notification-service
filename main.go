package main

import "notification-service/email"


func main() {
	// VerifyEmail
	// email.VerifyEmail("tahmid.saj@gmail.com", "tahmid.saj@gmail.com")

	email.SendEmail("tahmid.saj@gmail.com", "tahmid.saj@gmail.com", "subject", "text1", "test2", "UTF-8")
}