package mailing

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/vanng822/go-premailer/premailer"

	mail "github.com/xhit/go-simple-mail/v2"
)

type SMTPClient struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption mail.Encryption
	FromEmail  string
	FromName   string
}

const mailHtmlTemplateAddress = "./templates/email.html.gohtml"
const mailPlainTemplateAddress = "./templates/email.plain.gohtml"

func (c *SMTPClient) SendMessage(msg *Message) error {
	if msg.From == "" {
		msg.From = c.FromEmail
	}
	if msg.FromName == "" {
		msg.FromName = c.FromName
	}

	email, err := c.buildEmail(msg)

	if err != nil {
		return err
	}

	smtpClient, err := c.buildSMTPClient()

	if err != nil {
		return err
	}

	err = email.Send(smtpClient)

	if err != nil {
		return err
	}

	return nil
}

func (c *SMTPClient) buildHtmlMessage(msg *Message) (string, error) {
	tpl, err := template.New("html-template").ParseFiles(mailHtmlTemplateAddress)

	if err != nil {
		return "", err
	}

	var tplString bytes.Buffer

	tpl.ExecuteTemplate(&tplString, "body", msg.DataMap)

	prm, err := premailer.NewPremailerFromBytes(tplString.Bytes(), &premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   true,
		KeepBangImportant: true,
	})

	if err != nil {
		return "", err
	}

	return prm.Transform()
}

func (c *SMTPClient) buildPlainTextMessage(msg *Message) (string, error) {
	tpl, err := template.New("plain-template").ParseFiles(mailPlainTemplateAddress)

	if err != nil {
		return "", err
	}

	var tplString bytes.Buffer

	err = tpl.ExecuteTemplate(&tplString, "body", msg.DataMap)

	if err != nil {
		return "", err
	}

	return tplString.String(), nil
}

func (c *SMTPClient) buildSMTPClient() (*mail.SMTPClient, error) {
	mailSrv := mail.NewSMTPClient()
	mailSrv.Host = c.Host
	mailSrv.Port = c.Port
	mailSrv.Username = c.Username
	mailSrv.Password = c.Password
	mailSrv.Encryption = c.Encryption
	mailSrv.KeepAlive = false
	mailSrv.ConnectTimeout = 15 * time.Second
	mailSrv.SendTimeout = 20 * time.Second

	return mailSrv.Connect()
}

func (c *SMTPClient) buildEmail(msg *Message) (*mail.Email, error) {
	formatted, err := c.buildHtmlMessage(msg)

	if err != nil {
		return nil, err
	}

	plain, err := c.buildPlainTextMessage(msg)

	if err != nil {
		return nil, err
	}

	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("%s <%s>", msg.FromName, msg.From)).
		SetSubject(msg.Subject).
		SetBody(mail.TextPlain, plain).
		AddAlternative(mail.TextHTML, formatted).
		AddTo(msg.To)

	for _, attachment := range msg.Attachments {
		email.AddAttachment(attachment)
	}
	return email, nil
}

func NewSMTPClient(fromEmail, fromName string) *SMTPClient {
	u, err := url.Parse(os.Getenv("SMTP_URL"))

	if err != nil {
		log.Panic("SMTP url is not configured!")
	}

	if len(fromName) == 0 {
		fromName = "Microservices Administrator"
	}

	if len(fromEmail) == 0 {
		fromEmail = "gomicroservices@example.com"
	}

	pass, ok := u.User.Password()

	if !ok {
		pass = ""
	}

	port, err := strconv.Atoi(u.Port())
	if err != nil {
		port = 0
	}

	return &SMTPClient{
		Host:       u.Hostname(),
		Port:       port,
		Username:   u.User.Username(),
		Password:   pass,
		Encryption: mail.EncryptionNone,
		FromEmail:  fromEmail,
		FromName:   fromName,
	}
}
