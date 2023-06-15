package email_template

import (
	"log"
	"net/smtp"
	"strings"
	texttemp "text/template"

	"github.com/USACE/instrumentation-api/api/config"
)

const CharSet = "UTF-8"

type EmailTemplateContent struct {
	TextSubject, TextBody *texttemp.Template
}

type EmailContent struct {
	TextSubject, TextBody string
	To                    []string
}

func CreateEmailTemplateContent(preformatted EmailContent) (*EmailTemplateContent, error) {
	tTextSubject, err := texttemp.New("textSubject").Parse(preformatted.TextSubject)
	if err != nil {
		return nil, err
	}
	tTextBody, err := texttemp.New("textBody").Parse(preformatted.TextBody)
	if err != nil {
		return nil, err
	}
	return &EmailTemplateContent{
		TextSubject: tTextSubject,
		TextBody:    tTextBody,
	}, nil
}

func FormatAlertConfigTemplates(templContent *EmailTemplateContent, data any) (*EmailContent, error) {
	var textSubject, textBody strings.Builder

	if err := templContent.TextSubject.Execute(&textSubject, data); err != nil {
		return nil, err
	}
	if err := templContent.TextBody.Execute(&textBody, data); err != nil {
		return nil, err
	}

	return &EmailContent{
		TextSubject: textSubject.String(),
		TextBody:    textBody.String(),
	}, nil
}

func ConstructAndSendEmail(ec *EmailContent, cfg *config.AlertCheckConfig, smtpConfig *config.SmtpConfig) error {
	if cfg.EmailSendMocked {
		log.Printf("mocking email '%s':\n%s", ec.TextSubject, ec.TextBody)
		return nil
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + ec.TextSubject + "\n"
	msg := []byte(subject + mime + "\n" + ec.TextSubject)

	if err := smtp.SendMail(smtpConfig.SmtpAddr, smtpConfig.SmtpAuth, cfg.EmailFrom, ec.To, msg); err != nil {
		return err
	}
	return nil
}
