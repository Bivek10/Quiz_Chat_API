package services

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
	"google.golang.org/api/gmail/v1"
)

type GmailService struct {
	gmailService *gmail.Service
	logger       infrastructure.Logger
}

func NewGmailService(gmailService *gmail.Service, logger infrastructure.Logger) GmailService {
	return GmailService{
		gmailService: gmailService,
		logger:       logger,
	}
}

func (g GmailService) SendEmail(params models.EmailParams) (bool, error) {
	to := params.To
	emailBody, err := utils.ParseTemplate(params.BodyTemplate, params.BodyData)
	if err != nil {
		return false, errors.New("unable to parse email body template")
	}
	fmt.Println("gmail serviec lien 32")
	var msgString string
	emailTo := "To: " + to + "\r\n"
	msgString = emailTo
	subject := "Subject: " + params.SubjectData + "\n"
	msgString = msgString + subject
	msgString = msgString + "\n" + emailBody
	var msg []byte

	fmt.Println("gmail serviec lien 41")

	if params.Lang != nil && *params.Lang != "en" {
		msgStringJP, _ := utils.ToISO2022JP(msgString)
		msg = []byte(msgStringJP)
	} else {
		msg = []byte(msgString)
		fmt.Println("gmail serviec lien 47")
	}
	message := gmail.Message{
		Raw: base64.URLEncoding.EncodeToString([]byte(msg)),
	}
	fmt.Println("gmail serviec lien 53")
	_, err = g.gmailService.Users.Messages.Send("me", &message).Do()
	fmt.Println("gmail serviec lien 55")
	if err != nil {
		fmt.Println("gmail serviec lien 56")
		return false, err
	}
	fmt.Println("gmail serviec lien 57")
	return true, nil
}
