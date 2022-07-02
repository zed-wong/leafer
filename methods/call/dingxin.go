package call

import(
	"fmt"
	"errors"
        "github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

const (
        Endpoint = "http://yuyin2.market.alicloudapi.com/dx/voice_notice"
)

func DingXin(phone, tplID, appcode string) error{
	if len(phone) == 0{
		return errors.New("Dingxin() => len(Phone) == 0")
	}
	client := resty.New()
	params := map[string]string{
		"param": "param",
		"phone": phone,
		"tpl_id": tplID,
	}

	resp, err := client.R().
		SetQueryParams(params).
		SetHeader("Authorization","APPCODE "+appcode).
		Post(Endpoint)
	if err != nil {
		return err
	}
	if gjson.Get(resp.String(), "return_code").String() != "00000"{
		return errors.New(fmt.Sprintf("DingXin error:", resp.String()))
	}

	return nil
}
