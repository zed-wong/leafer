package sms

import (
	"fmt"
	"errors"
        "strings"
        "crypto/md5"
        "encoding/hex"

        "github.com/go-resty/resty/v2"
)

const (
	smsbaoEndpoint = "https://api.smsbao.com/sms"
)

func Md5Encode(s string)string{
        hash := md5.Sum([]byte(s))
        return hex.EncodeToString(hash[:])
}

func SmsBao(user, pass, phone, message string) error{
        if len(phone) == 0{
                return errors.New("Phone Number Undefined")
        }
        client := resty.New()
	fmt.Println("SmsBao.phone:", phone)
        params := map[string]string{
                "u":user,
                "p":Md5Encode(pass),
                "m":phone,
                "c":message,
        }
        resp, err := client.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		Get(smsbaoEndpoint)

        if err != nil{
		return err
        } else{
                if resp.String()!="0"{
			return errors.New(fmt.Sprintf("Send Message Failed: %s", resp.String()))
                }
        }
	return nil
}

// send multi message (len(phones) < 100)
func SmsBaos(user, pass string, phones []string, message string) error{
        if len(phones) == 0{
                return errors.New("Zero Phone Number")
        }
        var phone string
        for i:=0; i<len(phones); i++{
                phone+=phones[i]+","
        }
        phone = strings.TrimSuffix(phone,",")
        client := resty.New()
        params := map[string]string{
                "u":user,
                "p":Md5Encode(pass),
                "m":phone,
                "c":message,
        }
        resp, err := client.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		Get(smsbaoEndpoint)

        if err != nil{
		return err
        } else{
                if resp.String()!="0"{
                        return errors.New(fmt.Sprintf("Send Message Failed: %s", resp.String()))
                }
        }
	return nil
}
