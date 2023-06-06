package email_template

import (
	htmltemp "html/template"
	"log"
	"strings"
	texttemp "text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

const CharSet = "UTF-8"

type EmailTemplateContent struct {
	TextSubject, TextBody *texttemp.Template
	HtmlBody              *htmltemp.Template
}

type EmailContent struct {
	TextSubject, TextBody, HtmlBody string
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
	tHtmlBody, err := htmltemp.New("htmlBody").Parse(preformatted.HtmlBody)
	if err != nil {
		return nil, err
	}
	return &EmailTemplateContent{
		TextSubject: tTextSubject,
		TextBody:    tTextBody,
		HtmlBody:    tHtmlBody,
	}, nil
}

func FormatAlertConfigTemplates(templContent *EmailTemplateContent, data any) (*EmailContent, error) {
	var textSubject, textBody, htmlBody strings.Builder

	if err := templContent.TextSubject.Execute(&textSubject, data); err != nil {
		return nil, err
	}
	if err := templContent.TextBody.Execute(&textBody, data); err != nil {
		return nil, err
	}
	if err := templContent.HtmlBody.Execute(&htmlBody, data); err != nil {
		return nil, err
	}

	return &EmailContent{
		TextSubject: textSubject.String(),
		TextBody:    textBody.String(),
		HtmlBody:    htmlBody.String(),
	}, nil
}

func ConstructAndSendEmail(svc *ses.SES, ec *EmailContent, toAddresses []*string, sender string, mock bool) error {
	if mock {
		log.Printf("mocking email '%s':\n%s", ec.TextSubject, ec.TextBody)
		return nil
	}
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(ec.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(ec.TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(ec.TextSubject),
			},
		},
		Source: aws.String(sender),
	}

	if _, err := svc.SendEmail(input); err != nil {
		return err
	}
	return nil
}
