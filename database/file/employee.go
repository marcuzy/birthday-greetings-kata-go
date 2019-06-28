package file

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/marcuzy/birthday-greetings-kata-go/employee"
)

type employeeRepository struct {
	employees []employee.Employee
}

func NewEmployeeRepository(filename string) (employee.EmployeeRepositry, error) {
	r := &employeeRepository{}
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")[1:]

	for _, line := range lines {
		cols := strings.Split(line, ", ")
		dateOfBirth, _ := time.Parse("2006/01/02", cols[2])

		empl := employee.Employee{
			LastName:    cols[0],
			FirstName:   cols[1],
			DateOfBirth: dateOfBirth,
			Email:       cols[3],
		}

		r.employees = append(r.employees, empl)
	}

	return r, nil
}

func (r *employeeRepository) FindEmployeesBornOn(ctx context.Context, month time.Month, day int) ([]employee.Employee, error) {
	var res []employee.Employee

	for _, empl := range r.employees {
		if empl.DateOfBirth.Month() == month && empl.DateOfBirth.Day() == day {
			res = append(res, empl)
		}
	}

	return res, nil
}
