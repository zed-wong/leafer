package workers

import (
	"log"
	"time"
	"sync"
	"strings"
        "math/rand"

        "gorm.io/gorm"
	"github.com/zed-wong/leafer/api"
	"github.com/zed-wong/leafer/methods/mixinMessenger"
	"github.com/zed-wong/leafer/methods/telegram"
	"github.com/zed-wong/leafer/methods/signal"
	"github.com/zed-wong/leafer/methods/sms"
	"github.com/zed-wong/leafer/methods/call"
	"github.com/fox-one/mixin-sdk-go"
)

type MethodWorker struct{
	client		*mixin.Client
	telegram	string
	signal		string
	sms		*SMSConf
	call		*CallConf
}

type SMSConf struct{
	smsbao  *SMSBAO
	twilio	*TWILIO
}

type CallConf struct{
	dingxin *DINGXIN
	twilio  *TWILIO
}

type SMSBAO struct{
	user	string
	pass	string
}

type DINGXIN struct{
	appcode	string
	tplid	string
}

type TWILIO struct{
	from	string
	sid	string
	token	string
}

func NewMethodWorker(store *mixin.Keystore, tg, signal, smsbaoUser, smsbaoPass, dingxinCode, dingxinTpl, twilioFrom, twilioSid, twilioToken string) *MethodWorker{
        client, err := mixin.NewFromKeystore(store)
        if err != nil {
                panic(err)
        }

	return &MethodWorker{
		client: client,
		telegram: tg,
		signal: signal,
		sms: &SMSConf{
			smsbao: &SMSBAO{
				user: smsbaoUser,
				pass: smsbaoPass,
			},
			twilio: &TWILIO{
				from: twilioFrom,
				sid: twilioSid,
				token: twilioToken,
			},
		},
		call: &CallConf{
			dingxin: &DINGXIN{
				appcode: dingxinCode,
				tplid: dingxinTpl,
			},
			twilio: &TWILIO{
				from: twilioFrom,
				sid: twilioSid,
				token: twilioToken,
			},
		},
	}
}

func (mw *MethodWorker) Mixin(userID, conversationID, data string){
	if err := mixinMessenger.SendMixinMsg(mw.client, userID, conversationID, []byte(data)); err != nil{
		log.Println("mw.Mixin() =>", err)
	}
}

func (mw *MethodWorker) MixinBtn(userID, conversationID, label, action, color string){
	if err := mixinMessenger.SendMixinBtn(mw.client, userID, conversationID, label, action, color); err != nil{
		log.Println("mw.MixinBtn() =>", err)
	}
}

func (mw *MethodWorker) Tg(chatID, data string){
	tw := telegram.NewTelegramWorker(mw.telegram)
	if err := tw.SendTgMsg(chatID, data); err != nil{
		log.Println("mw.Tg() =>", err)
	}
}

func (mw *MethodWorker) Signal(to, data string){
	if err := signal.Message(mw.signal, to, data); err != nil{
		log.Println("mw.Signal() =>", err)
	}
}

func (mw *MethodWorker) SMS(number, data string){
	number = strings.TrimSpace(number)
	if isChinese(number){
		number = number[3:]
		if err := sms.SmsBao(mw.sms.smsbao.user, mw.sms.smsbao.pass, number, data); err != nil{
			log.Println("mw.Sms.Smsbao() =>", err)
		}
	} else {
		if err := sms.Twilio(mw.sms.twilio.from, mw.sms.twilio.sid, mw.sms.twilio.token, number, data); err != nil{
			log.Println("mw.Sms.Twilio() =>", err)
		}
	}
}

func (mw *MethodWorker) Call(number string){
	number = strings.TrimSpace(number)
	if isChinese(number){
                number = number[3:]
		if err := call.DingXin(number, mw.call.dingxin.tplid, mw.call.dingxin.appcode); err != nil{
			log.Println("mw.Call.DingXin() =>", err)
		}
	} else {
		if err := call.Twilio(mw.call.twilio.from, mw.call.twilio.sid, mw.call.twilio.token, number); err != nil{
			log.Println("mw.Call.Twilio() =>", err)
		}
	}
}

func isChinese(number string) bool{
	number = strings.TrimSpace(number)
	if number[0:3] == "+86" && len(number) == 14{
		return true
	}
	return false
}

// Concurrency
func (mw *MethodWorker) MixinCon(userID, conversationID, data string, wg *sync.WaitGroup){wg.Add(1); mw.Mixin(userID, conversationID, data); log.Println("MixinDone"); wg.Done() }
func (mw *MethodWorker) TgCon(chatID, data string, wg *sync.WaitGroup){ wg.Add(1); mw.Tg(chatID, data); log.Println("TgDone"); wg.Done() }
func (mw *MethodWorker) SignalCon(to, data string, wg *sync.WaitGroup){ wg.Add(1); mw.Signal(to, data); log.Println("SignalDone"); wg.Done() }
func (mw *MethodWorker) SMSCon(number, data string, wg *sync.WaitGroup){ wg.Add(1); mw.SMS(number, data); log.Println("SMSDone"); wg.Done() }
func (mw *MethodWorker) CallCon(number string, wg *sync.WaitGroup){ wg.Add(1); mw.Call(number); log.Println("CallDone"); wg.Done() }

func (mw *MethodWorker) AlertAllMethods(db *gorm.DB, user *api.USER, vault *api.VAULT, data, censored string, wg *sync.WaitGroup) {
	usedMethods := "M"
	go mw.MixinCon(user.UserID, user.ConversationID, data, wg)
	if len(user.TgID) != 0 {
		go mw.TgCon(user.TgID, data, wg)
		usedMethods += "t"
	}
	if len(user.SignalNumber) != 0 {
		go mw.SignalCon(user.SignalNumber, data, wg)
		usedMethods += "s"
	}
	if user.SMSBalance > 0{
		go mw.SMSCon(user.PhoneNumber, censored, wg)
		api.CostService(db, "sms", 1, user.UserID)
		usedMethods += "m"
	}
	if user.CallBalance > 0{
		go mw.CallCon(user.PhoneNumber, wg)
		api.CostService(db, "call", 1, user.UserID)
		usedMethods += "c"
	}
	api.AddHistory(db, usedMethods, time.Now().Format(api.TIMEFORMAT), user.UserID)
	wg.Wait()
}

func RandID(n int)string{
        var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
        rand.Seed(time.Now().UnixNano())
        b := make([]rune, n)
        for i := range b {
                b[i] = letters[rand.Intn(len(letters))]
        }
        return string(b)
}

