package models

import "notification-service/email"

type EmailVerificationInput struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}

type EmailInput struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	HTMLBody  string `json:"htmlBody"`
	TextBody  string `json:"textBody"`
	CharSet   string `json:"charSet"`
}

// email
func VerifyEmail(emailVerificationInput *EmailVerificationInput) (*Response, error) {
	resVerifyEmail, err := email.VerifyEmail((*emailVerificationInput).Sender, (*emailVerificationInput).Recipient)
	if err != nil {
		return &Response{
			Ok: false,
			Response: false,
		}, err
	}

	return &Response{
		Ok: true,
		Response: resVerifyEmail,
	}, nil
}

func SendEmail(emailInput *EmailInput) (*Response, error) {
	resSendEmail, err := email.SendEmail((*emailInput).Sender, (*emailInput).Recipient, (*emailInput).Subject,
		(*emailInput).HTMLBody, (*emailInput).TextBody, (*emailInput).CharSet)
	if err != nil {
		return &Response{
			Ok: false,
			Response: false,
		}, err
	}

	return &Response{
		Ok: true,
		Response: resSendEmail,
	}, nil
}