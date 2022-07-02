package workers_test

import (
	"fmt"
	"sync"
	"testing"
        "gorm.io/gorm"
        "gorm.io/driver/postgres"
	"github.com/spf13/viper"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/zed-wong/leafer/api"
        "github.com/zed-wong/leafer/workers"
)

func initDB() *gorm.DB{
        viper.SetConfigName("config")
        viper.AddConfigPath("../")
        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("viper.ReadInConfig() =>", err.Error())
        }

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", viper.GetString("database.host"), viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.dbname"), viper.GetString("database.port"), viper.GetString("database.sslmode"))
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
                panic(err)
        }
	return db
}

func initWorker() *workers.MethodWorker{
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

func TestMixin(t *testing.T){
	mw := initWorker()
	mw.Mixin("44d9717d-8cae-4004-98a1-f9ad544dcfb1", "f9e1d053-b727-3969-8ae2-d219ff2535da", fmt.Sprintf(api.EN.AlertMsg, "123", "205"))
}

func TestTg(t *testing.T){
	mw := initWorker()
	mw.Tg("4605", fmt.Sprintf(api.EN.AlertMsg, "2351", "215"))
}

func TestSignal(t *testing.T){
	mw := initWorker()
	mw.Signal("2092", fmt.Sprintf(api.ZH.AlertMsg, "51", "187"))
}

func TestSMS(t *testing.T){
	mw := initWorker()
	mw.SMS("+85901", fmt.Sprintf(api.ZH.CensoredMsg, workers.RandID(4)))
}

func TestCall(t *testing.T){
	mw := initWorker()
	mw.Call("+814901")
}

func TestAlertAllMethods(t *testing.T){
	var wg sync.WaitGroup
	mw := initWorker()
	db := initDB()
	user := api.GetUser(db, "44d9717d-8cae-4004-98a1-f9ad544dcfb1")
	vault := api.GetVault(db, "100")

	mw.AlertAllMethods(db, user, vault, api.ZH.AlertMsg, api.ZH.CensoredMsg, &wg)
}

func TestSplitNum(t *testing.T){
	number := "+863102331"
	t.Log(number[3:])
}
