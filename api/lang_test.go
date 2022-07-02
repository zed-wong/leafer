package api

import (
	"fmt"
	"testing"
	"github.com/zed-wong/leafer/api"
)

func TestCheckFirstVault(t *testing.T){
	content := api.GetContentByLang("")
	fmt.Printf("%+v",content)
}

