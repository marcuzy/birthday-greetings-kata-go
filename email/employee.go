package email

import (
	"context"
	"fmt"
	"net/smtp"
	"net/url"
	"strconv"
	"strings"

	emailLib "github.com/jordan-wright/email"
	"github.com/marcuzy/birthday-greetings-kata-go/employee"
)

type employeeEmailService struct {
	from string
	// Username string
	// Password string
	// Host     string
	// Port     int
	url *url.URL
}

func NewEmployeeEmailService(from string, URI string) employee.EmailService {
	parsedUrl, err := url.Parse(URI)
	if err != nil {
		panic(err)
	}
	return &employeeEmailService{
		url:  parsedUrl,
		from: from,
		// From:     args.From,
		// Username: args.Username,
		// Password: args.Password,
		// Host:     args.Host,
		// Port:     args.Port,
	}
}

func (s *employeeEmailService) Send(ctx context.Context, to, subject, text string) error {
	e := emailLib.NewEmail()
	e.From = s.from
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(text)

	parsedHost := strings.Split(s.url.Host, ":")
	host := parsedHost[0]
	port, _ := strconv.Atoi(s.url.Port())
	password, _ := s.url.User.Password()
	username := s.url.User.Username()
	err := e.Send(
		fmt.Sprintf("%s:%d", host, port),
		smtp.PlainAuth("", username, password, host),
	)

	return err
}
