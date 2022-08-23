package messprovider

import (
	"context"
	"fmt"
	"github.com/twilio/twilio-go"
)
import openapi "github.com/twilio/twilio-go/rest/api/v2010"

var (
	client *twilio.RestClient
)

type MessProvider interface {
	SendMessage(ctx context.Context, toPhone string, msg any) error
}
type messProvider struct {
	accountId string
	tokenId   string
	fromPhone string
}

func NewSendMessage(accountId string, tokenId string, fromPhone string) *messProvider {
	return &messProvider{
		accountId: accountId,
		tokenId:   tokenId,
		fromPhone: fromPhone,
	}
}

func (p *messProvider) SendMessage(ctx context.Context, toPhone string, msg string) error {
	fmt.Printf("error creating and sending message")
	params := openapi.CreateMessageParams{}
	params.SetTo(p.fromPhone)
	params.SetFrom(toPhone)
	params.SetBody(msg)

	_, err := client.Api.CreateMessage(&params)

	if err != nil {
		return nil
	}
	return nil
}
