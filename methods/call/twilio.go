package call

import (
	"fmt"
	"errors"
        "github.com/twilio/twilio-go"
        openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

const (
	ApplicationSid = "AP90a597eae6"
)

func Twilio(from, sid, token, to string) error{
        client := twilio.NewRestClientWithParams(twilio.ClientParams{
                Username: sid,
                Password: token,
        })

        params := &openapi.CreateCallParams{}
        params.SetTo(to)
        params.SetFrom(from)
	params.SetApplicationSid(ApplicationSid)

        resp , err := client.ApiV2010.CreateCall(params)
        if err != nil {
                return err
	}
	if *resp.Status != "queued" {
		return errors.New(fmt.Sprintf("TwilioCall resp:", *resp))
	}
	return nil
}
