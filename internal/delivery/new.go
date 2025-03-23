package delivery

import "github.com/AmirHossein82x/doctor-appointment/internal/logger"



func TestLog() {
	log := logger.SetUpLogger()
	log.Error("test for line nine")
}