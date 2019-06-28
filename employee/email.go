package employee

import "context"

type EmailService interface {
	Send(ctx context.Context, to, subject, text string) error
}
