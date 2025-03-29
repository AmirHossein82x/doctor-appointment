package ports

type SmsService interface {
	SendSMS([]string, string) error
}
