package ses

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const sender = "rodrigodosanjosoliveira@gmail.com"

// sesEmailSender is the gateway to send emails using AWS SES
type sesEmailSender struct {
	sesClient *ses.SES
}

// NewSesEmailSender creates a new sesEmailSender
func NewSesEmailSender(sesClient *ses.SES) *sesEmailSender {
	return &sesEmailSender{
		sesClient: sesClient,
	}
}

// SendEmail sends an email to the given address
func (s *sesEmailSender) SendEmail(to string, subject string, body string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("use-east-1")},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	result, err := svc.SendEmail(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Email Sent to address: " + to)
	fmt.Println(result)
}
