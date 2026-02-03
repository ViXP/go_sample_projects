package server

import (
	"fmt"
	"log"
	"mailer-service/internal/mailing"
)

type RPCProcedures struct{}

func (r *RPCProcedures) SendMail(payload *EmailRequestPayload, reply *string) error {
	message := mailing.NewMessage(payload.To, payload.Subject, payload.Body)
	smtpClient := mailing.NewSMTPClient(payload.From, payload.FromName)

	err := smtpClient.SendMessage(message)

	if err != nil {
		log.Print(err)
		return err
	}

	logMsg := fmt.Sprintf("Email is sent successfully to %s\n", payload.To)
	log.Print(logMsg)
	*reply = logMsg

	return nil
}
