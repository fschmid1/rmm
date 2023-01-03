package controller

import (
	"time"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
	"github.com/gregdel/pushover"
)

func SendMessage(user models.User, msg string, title string) error {
	message := &pushover.Message{
		Message:   msg,
		Title:     title,
		Priority:  pushover.PriorityHigh,
		Timestamp: time.Now().Unix(),
		Retry:     60 * time.Second,
		Sound:     pushover.SoundVibrate,
	}
	recipient := pushover.NewRecipient(user.PushToken)
	_, err := config.Pusher.SendMessage(message, recipient)
	return err
}
