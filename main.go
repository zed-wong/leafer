package main

import (
	"fmt"
	"log"
	"context"

	"github.com/spf13/viper"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/zed-wong/leafer/api"
	"github.com/zed-wong/leafer/workers"
)

func main(){
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("viper.ReadInConfig() =>", err.Error())
	}
	var (
		host = viper.GetString("api.host")
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", viper.GetString("database.host"), viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.dbname"), viper.GetString("database.port"), viper.GetString("database.sslmode"))
		clientID = viper.GetString("methods.mixin.client_id")
		sessionID = viper.GetString("methods.mixin.session_id")
		privateKey = viper.GetString("methods.mixin.private_key")
		pinToken = viper.GetString("methods.mixin.pin_token")
		appSecret = viper.GetString("methods.mixin.app_secret")

		tgToken = viper.GetString("methods.telegram.token")
		signal = viper.GetString("methods.signal.from")

		smsbaoUser = viper.GetString("methods.smsbao.user")
		smsbaoPass = viper.GetString("methods.smsbao.pass")

		dingxinCode = viper.GetString("methods.dingxin.app_code")
		dingxinTpl = viper.GetString("methods.dingxin.template_id")

		twilioFrom = viper.GetString("methods.twilio.from_number")
		twilioSid = viper.GetString("methods.twilio.sid")
		twilioToken = viper.GetString("methods.twilio.token")
	)

	ctx := context.Background()
	store := &mixin.Keystore{
		ClientID: clientID,
		SessionID: sessionID,
		PrivateKey: privateKey,
		PinToken: pinToken,
	}
	conf := &api.Config{
		ClientID: clientID,
		AppSecret: appSecret,
		TgToken : tgToken,
		TwilioSid: twilioSid,
		TwilioToken: twilioToken,
	}

	apiWorker := api.NewAPIWorker(dsn, conf)
	updateWorker := workers.NewUpdaterWorker(dsn)
	methodWorker := workers.NewMethodWorker(store, tgToken, signal, smsbaoUser, smsbaoPass, dingxinCode, dingxinTpl, twilioFrom, twilioSid, twilioToken)
	notifyWorker := workers.NewNotifierWorker(dsn, methodWorker)
	messsageWorker := workers.NewMessageWorker(ctx, store, dsn, methodWorker)
	InitCron(updateWorker, methodWorker, notifyWorker)

	go messsageWorker.Loop(ctx)
	apiWorker.Loop(host)
}
