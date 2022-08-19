package otpProvider

import "context"

type OtpProvider interface {
	SendMessage(ctx context.Context, toPhone string, msg string) (res any, err error)
}
