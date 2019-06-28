package employee

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type EmployeeService interface {
	SendGreetings(ctx context.Context, date time.Time) error
}

type employeeService struct {
	repo  EmployeeRepositry
	email EmailService
}

func NewEmployeeService(repo EmployeeRepositry, email EmailService) EmployeeService {
	return &employeeService{
		repo:  repo,
		email: email,
	}
}

func (s *employeeService) SendGreetings(ctx context.Context, date time.Time) error {
	_, month, day := date.Date()
	employees, err := s.repo.FindEmployeesBornOn(ctx, month, day)

	if err != nil {
		return err
	}

	var emailErrors []error

	for _, employee := range employees {
		subject := "Happy birthday!"
		text := fmt.Sprintf("Happy birthday, dear %s!", employee.FirstName)
		err := s.email.Send(ctx, employee.Email, subject, text)

		if err != nil {
			emailErrors = append(emailErrors, err)
		}
	}

	if len(emailErrors) > 0 {
		return errors.New("Some emails were sent with errors")
	}

	return nil
}
