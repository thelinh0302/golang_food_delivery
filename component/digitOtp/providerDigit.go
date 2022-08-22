package digitOtp

import "time"

type Provider interface {
	GenerateOTP() (token string, err error)
}

type OTP struct {
	Otp     string    `json:"otp"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}
