package infrastructure

import (
	"fmt"
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	// "github.com/kavenegar/kavenegar-go"
)

// KavenegarSmsService is an implementation of SmsService for Kavenegar.
type KavenegarSmsService struct {
	apiKey string
	sender string
}

// NewKavenegarSmsService creates a new KavenegarSmsService.
func NewKavenegarSmsService() *KavenegarSmsService {
	config := config.LoadConfig()
	return &KavenegarSmsService{
		apiKey: config.KAVENEGAR_API_KEY,
		sender: config.KAVENEGAR_SENDER,

	}
}

// SendSMS sends an SMS using the Kavenegar API.
func (k *KavenegarSmsService) SendSMS(receptor []string, message string) error {
	time.Sleep(time.Second * 2)
	fmt.Printf("Sending SMS to %v: %v\n", receptor, message)
	// api := kavenegar.New(k.apiKey)

	// res, err := api.Message.Send(k.sender, receptor, message, nil)
	// if err != nil {
	// 	// Enhanced error handling
	// 	switch err := err.(type) {
	// 	case *kavenegar.APIError:
	// 		return fmt.Errorf("API error: %v", err.Error())
	// 	case *kavenegar.HTTPError:
	// 		return fmt.Errorf("HTTP error: %v", err.Error())
	// 	default:
	// 		return fmt.Errorf("unknown error: %v", err.Error())
	// 	}
	// }

	// // Log or return response details
	// for _, r := range res {
	// 	fmt.Println("MessageID  = ", r.MessageID)
	// 	fmt.Println("Status     = ", r.Status)
	// }
	return nil
}