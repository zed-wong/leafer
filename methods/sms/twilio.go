package sms

import (
        "github.com/twilio/twilio-go"
        openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Twilio(from, sid, token, to, message string) error{
        client := twilio.NewRestClientWithParams(twilio.ClientParams{
                Username: sid,
                Password: token,
        })

        params := &openapi.CreateMessageParams{}
        params.SetTo(to)
        params.SetFrom(from)
        params.SetBody(message)
        _ , err := client.ApiV2010.CreateMessage(params)
        if err != nil {
		return err
        }
	return nil
}
