package sms_sender

// SmsService is an interface for sending SMS.
type SmsService interface {
	SendSMS(receptor []string, message string) error
}