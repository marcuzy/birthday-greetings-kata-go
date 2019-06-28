package employee

import "time"

type Employee struct {
	LastName    string
	FirstName   string
	DateOfBirth time.Time
	Email       string
}
