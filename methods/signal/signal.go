package signal

import (
	"fmt"
	"errors"
	"github.com/go-resty/resty/v2"
)

const (
	//Endpoint = "http://localhost:8080"
	Endpoint = "http://"
)

func Message(from, to, data string) error{
	if len(from) == 0 {
		return errors.New("From Number Undefined")
	}
	if len(to) == 0 {
		return errors.New("To Number Undefined")
	}
	if len(data) == 0 {
		return errors.New("Data Undefined")
	}

	client := resty.New()
	_ , err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte( fmt.Sprintf(`{"message": "%s", "number": "%s", "recipients": ["%s"]}`, data, from, to))).
		Post(fmt.Sprintf("%s/v2/send", Endpoint))
	if err != nil {
		return err
	}
	return nil
}
