package employee

import (
	"context"
	"time"
)

type EmployeeRepositry interface {
	FindEmployeesBornOn(ctx context.Context, month time.Month, day int) ([]Employee, error)
}
