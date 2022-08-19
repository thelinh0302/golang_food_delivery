package otpProvider

import (
	"context"
	twilio "github.com/twilio/twilio-go"
)
import openapi "github.com/twilio/twilio-go/rest/api/v2010"

type otpServices struct {
	accountId string
	tokenId   string
	fromPhone string
}

var (
	client *twilio.RestClient
)

func NewOtpServicesProvider(accountId string, tokenId string, fromPhone string) *otpServices {
	otpProvider := &otpServices{
		accountId: accountId,
		tokenId:   tokenId,
		fromPhone: fromPhone,
	}
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: otpProvider.accountId,
		Password: otpProvider.tokenId,
	})
	return nil
}

func (provider *otpServices) SendMessage(ctx context.Context, toPhone string, msg string) (res any, err error) {
	params := openapi.CreateMessageParams{}
	params.SetTo(provider.fromPhone)
	params.SetFrom(toPhone)
	params.SetBody(msg)

	_, err = client.Api.CreateMessage(&params)

	if err != nil {
		return nil, err
	}
	return res, nil
}
