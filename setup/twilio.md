# Twilio setup

### To make a phone call, you will need:

1. Charge $20 (if this is your first time using twilio), then you can call anyone.

2. Go to google, search "text to mp3", enter the text you want twilio to say in call, convert to mp3 and download it to local. Listen it to see if something went wrong.

3. Upload your mp3 file to the cloud (I used google drive)

If you are using google drive too, you will need to:

- Upload mp3 to your account
- Share it with public access (get the link of file)
- Turn it into raw link
> origin: https://drive.google.com/file/d/XXX/view?usp=sharing 
> raw: https://drive.google.com/uc?id=XXX

Now you have raw link of the file.

BTW, the reason why I'm using mp3 instead of "Say" in TwiMl is that it requires some money every time you call the "Voice" API. Twilio will use "Amazon Polly Characters" to transform your text into audio every time.

4. Write a XML file in TwiMl format (https://www.twilio.com/docs/voice/twiml)

In my case, the file is like [this.](https://gist.github.com/zed-wong/e948506ef8138b55c3ebb73e51e50964)


5. Create a TwiML-Bin

In here https://www.twilio.com/console/twiml-bins

Then you will have a URL

6. Create a TwiML App

In here https://www.twilio.com/console/voice/twiml/apps

Input the TwiML-Bin URL into Voice or Messaging field (Depending on your need).

Create it, then you will have an Application SID.

7. Implement it into your method

For golang, it is like:

```go
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
    fmt.Println("Call Status: " + *resp.Status)
    fmt.Println("Call Sid: " + *resp.Sid)
    fmt.Println("Call Direction: " + *resp.Direction)
    return nil
}
```



---

## Reference:

1. https://testrtc.com/docs/setting-up-a-twilio-twiml-app-for-qualityrtc-ivr/
2. https://www.twilio.com/docs/voice/twiml
3. https://www.twilio.com/docs/voice/make-calls#
4. https://stackoverflow.com/questions/24834877/google-drive-raw-data
