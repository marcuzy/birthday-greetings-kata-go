package main

import (
	"context"
	"os"
	"time"

	"github.com/marcuzy/birthday-greetings-kata-go/database/file"
	"github.com/marcuzy/birthday-greetings-kata-go/email"
	"github.com/marcuzy/birthday-greetings-kata-go/employee"
)

func main() {
	repo, err := file.NewEmployeeRepository("base.txt")
	if err != nil {
		panic(err)
	}

	emailService := email.NewEmployeeEmailService(
		os.Getenv("EMAILFROM"),
		os.Getenv("SMTP"),
	)

	grService := employee.NewEmployeeService(repo, emailService)

	err = grService.SendGreetings(context.TODO(), time.Now())

	if err != nil {
		panic(err)
	}
}
