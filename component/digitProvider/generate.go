package digitProvider

import (
	"crypto/rand"
)

const otpChars = "1234567890"

func GenerateOT1P() (*Otp, error) {
	buffer := make([]byte, 6)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < 6; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return &Otp{
		Otp: string(buffer),
	}, nil

}
