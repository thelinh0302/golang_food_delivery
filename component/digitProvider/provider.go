package digitProvider

type Provider interface {
	GenerateOT1P() (*Otp, error)
}

type Otp struct {
	Otp string `json:"otp"`
}
