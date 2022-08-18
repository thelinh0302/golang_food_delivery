package otpServices

import (
	"context"
	twilio "github.com/twilio/twilio-go"
)
import openapi "github.com/twilio/twilio-go/rest/api/v2010"

type otpServices struct {
	accountId string
	tokenId   string
	toPhone   string
	client    *twilio.RestClient
}

func otpServicesProvider(accountId string, tokenId string, client *twilio.RestClient) *otpServices {
	otpProvider := &otpServices{
		accountId: accountId,
		tokenId:   tokenId,
	}
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: otpProvider.accountId,
		Password: otpProvider.tokenId,
	})
	return otpProvider
}

func (provider *otpServices) sendMessage(ctx context.Context, fromPhone string, msg string) (res any, err error) {
	params := openapi.CreateMessageParams{}
	params.SetTo(provider.toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	_, err = provider.client.Api.CreateMessage(&params)

	if err != nil {
		return nil, err
	}
	return res, nil
}
