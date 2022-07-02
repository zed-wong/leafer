package workers_test

import (
	"log"
	"fmt"
	"testing"
	"github.com/spf13/viper"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/zed-wong/leafer/workers"
)

func initdb() string{
        viper.SetConfigName("config")
        viper.AddConfigPath("../")
        err := viper.ReadInConfig()
        if err != nil {
                log.Fatalln("viper.ReadInConfig() =>", err.Error())
        }
        dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", viper.GetString("database.host"), viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.dbname"), viper.GetString("database.port"), viper.GetString("database.sslmode"))
        return dsn
}

func initmw() *workers.MethodWorker{
        viper.SetConfigName("config")
        viper.AddConfigPath("../")
        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("viper.ReadInConfig() =>", err.Error())
        }

        var (
                clientID = viper.GetString("methods.mixin.client_id")
                sessionID = viper.GetString("methods.mixin.session_id")
                privateKey = viper.GetString("methods.mixin.private_key")
                pinToken = viper.GetString("methods.mixin.pin_token")
       //         appSecret = viper.GetString("methods.mixin.app_secret")

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

        store := &mixin.Keystore{
                ClientID: clientID,
                SessionID: sessionID,
                PrivateKey: privateKey,
                PinToken: pinToken,
        }
        mw := workers.NewMethodWorker(store, tgToken, signal, smsbaoUser, smsbaoPass, dingxinCode, dingxinTpl, twilioFrom, twilioSid, twilioToken)
        return mw
}

func initnw() *workers.NotifierWorker{
	return workers.NewNotifierWorker(initdb(), initmw())
}

func TestAlertRenewVaults(t *testing.T){
	nw := initnw()
	nw.AlertRenewVaults()
}

func TestAlertChargeService(t *testing.T){
	nw := initnw()
	nw.AlertChargeService()
}
