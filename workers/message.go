package workers

import (
	"fmt"
	"log"
	"time"
	"regexp"
	"context"
	"strings"
	"encoding/json"
	"encoding/base64"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"github.com/shopspring/decimal"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/zed-wong/leafer/api"
)

const (
	DEVUID="44d9717d-8cae-4004-98a1-f9ad544dcfb1"
	DEVCID="f9e1d053-b727-3969-8ae2-d219ff2535da"

	USDT="4d8c508b-91c5-375b-92b0-ee702ed2dac5"
	pUSD="31d2ea9c-95eb-3355-b65b-ba096853bc18"
	CNB="965e5c6e-434c-3fa9-b780-c50f43cd955c"
	LeafBotID="75f18fe8-b056-46d6-9c48-0214425e58ce"
	LeafURLPt="https://leaf.pando.im/#/vault/detail"
	UUIDPt="[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
	TIMEFORMAT="2006-01-02T15:04:05Z"
)

var (
	AcceptAsset = pUSD
	pricePerYear = decimal.NewFromInt(12)
)

type MessageWorker struct{
	client	*mixin.Client
	db	*gorm.DB
	mw	*MethodWorker
}

func NewMessageWorker(ctx context.Context, store *mixin.Keystore, dsn string, mw *MethodWorker) *MessageWorker{
        client, err := mixin.NewFromKeystore(store)
        if err != nil {
                panic(err)
        }
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
        rw := &MessageWorker{
                client: client,
		db: db,
		mw: mw,
        }
        return rw
}

func (rw *MessageWorker) handlePlainText(ctx context.Context, msg *mixin.MessageView, data []byte) {
	var (
		user *api.USER
		content *api.LANG
		lang	string
	)

	udata := strings.ToUpper(string(data))
	switch udata {
	case "HI", "HELLO", "HOLA", "HALLO",  "안녕하세요", "こんにちは":
		lang = "en"
		rw.respond(ctx, msg, mixin.MessageCategoryPlainText, []byte(api.EN.Hi))
	case "你好":
		lang = "zh"
		rw.respond(ctx, msg, mixin.MessageCategoryPlainText, []byte(api.ZH.Hi))
	case "C":
		rw.respond(ctx, msg, mixin.MessageCategoryPlainText, []byte(fmt.Sprintf("%s:%s", msg.UserID, msg.ConversationID)))
	default:
		log.Println(fmt.Sprintf("%s:%s",msg.UserID,string(data)))
	}

	// Add user if not exist
	if !api.UserExist(rw.db, msg.UserID) {
		User, _ := rw.client.ReadUser(ctx, msg.UserID)
		api.AddUser(rw.db, api.USER{
			UserID: msg.UserID,
			ConversationID: msg.ConversationID,
			IdentityID: User.IdentityNumber,
			Name: User.FullName,
			Avatar: User.AvatarURL,
			Lang: lang,
		})
		switch (lang) {
		case "en":
			content = &api.EN
		case "zh":
			content = &api.ZH
		default:
			content = &api.EN
		}
	} else {
		user = api.GetUser(rw.db, msg.UserID)
		content = api.GetContentByLang(user.Lang)
	}

	// Add vault if vault url
	if vaultID := GetVaultIDFrom(string(data)); len(vaultID) != 0{
		rw.addvault(ctx, msg, vaultID, content)
	}
}

func (rw *MessageWorker) handleSnapshot(ctx context.Context, msg *mixin.MessageView, data []byte) {
	var (
		t mixin.TransferInput
		c string
	)
	if msg.UserID == rw.client.ClientID {
		log.Println(msg.UserID,"=",rw.client.ClientID)
		return
	}

	json.Unmarshal(data, &t)
	memo0, err := base64.StdEncoding.DecodeString(t.Memo)
	memo := string(memo0)

	if len(memo) == 0 || (t.AssetID != pUSD && t.AssetID != USDT && t.AssetID != CNB) || err != nil{
		log.Println("memo:",memo)
		return
	}
	user := api.GetUser(rw.db, t.OpponentID)
	content := api.GetContentByLang(user.Lang)
	action := gjson.Get(string(memo), "action").String()

	switch (action) {
	case "add":
		var addState bool
		vaultID := gjson.Get(memo, "vaultid").String()
		if t.Amount.Equal(pricePerYear) && len(vaultID) != 0{
			addState = api.AddVault(rw.db, vaultID, user.UserID)
		}
		if addState{
			c = fmt.Sprintf(content.AddSuccess, time.Now().AddDate(1, 0, 0).Format("2006-01-02"))
		} else {
			c = content.AddFailed
		}

	case "renew":
		var renewState bool
		identityID := gjson.Get(memo, "identityid").String()
		if t.Amount.Equal(pricePerYear) && len(identityID) != 0{
			renewState = api.RenewVault(rw.db, identityID)
		}
		if renewState{
			c = fmt.Sprintf(content.RenewSuccess, time.Now().AddDate(1, 0, 0).Format("2006-01-02"))
		} else {
			c = content.RenewFailed
		}

	case "buy":
		method := gjson.Get(memo, "type").String()
		plan := gjson.Get(memo, "plan").String()
		num := gjson.Get(memo, "num").Int()
		price := t.Amount.String()

		if VerifyBuy(user.PhoneNumber, method, plan, price){
			if api.BuyService(rw.db, method, int(num), user.UserID){
				user = api.GetUser(rw.db, user.UserID)
				c = fmt.Sprintf(content.BuySuccess, user.SMSBalance, user.CallBalance)
			} else {
				c = content.BuyFailed
			}
		}

	case "test":
		method := gjson.Get(memo, "type").String()
		price := t.Amount.String()
		if method == "sms" && price == "0.1"{
			rw.mw.SMS(user.PhoneNumber, fmt.Sprintf(content.CensoredMsg, RandID(4)))
			c = content.SmsSent
		} else if method == "call" && price == "0.3"{
			rw.mw.Call(user.PhoneNumber)
			c = content.CallSent
		} else {
			c = content.TestFailed
		}
	default:
		c = content.NoAction
	}
	api.AddPayment(rw.db, t.OpponentID, t.TraceID, t.Amount.String(), action)
	rw.respond(ctx, msg, mixin.MessageCategoryPlainText, []byte(c))
	log.Println(fmt.Sprintf("%s(%s) paid %s", user.Name, user.IdentityID, t.Amount.String()))
}

func (rw *MessageWorker) handleAppCard(ctx context.Context, msg *mixin.MessageView, data []byte) {
	// Check if Leaf URL
	// yes  -> Get vaultID
	//	   Check if first vault
	//	   yes -> Add free vault
	//	   no  -> Pay for vault
	//
	// no 	-> return

	user := api.GetUser(rw.db, msg.UserID)
	if user == nil{
		//respond() no record, please oauth first
		return
	}
	content := api.GetContentByLang(user.Lang)

	if gjson.Get(string(data), "app_id").String() == LeafBotID{
		vaultUrl := gjson.Get(string(data), "action").String()
		if vaultID := GetVaultIDFrom(vaultUrl); len(vaultID) != 0{
			rw.addvault(ctx, msg, vaultID,content)
		}
	}
}

func (rw *MessageWorker) addvault(ctx context.Context, msg *mixin.MessageView, vaultID string, content *api.LANG){
	var c string
	if api.CheckFirstVault(rw.db, msg.UserID){
		if api.AddVault(rw.db, vaultID, msg.UserID) {
			c = content.AddFreeSuccess
		} else {
			c = content.AddFailed
		}
		rw.respond(ctx, msg, mixin.MessageCategoryPlainText, []byte(c))

	} else if api.IsAddedVault(rw.db, vaultID){
		c = content.AddExisted
		rw.respond(ctx, msg, mixin.MessageCategoryPlainText, []byte(c))

	} else {
		button := &mixin.AppButtonMessage{
			Label: content.PayAdding,
			Action: PayForAddVault(ctx, rw.client.ClientID, vaultID),
			Color: "#5979F0",
		}
		btngrp := &mixin.AppButtonGroupMessage{
			*button,
		}
		b, err := json.Marshal(btngrp)
		if err != nil {
			log.Println("handleAppCard()=>json.Marshal(btngrp) => ", err)
		}
		c = content.PayAddingMsg
		rw.respond(ctx, msg, mixin.MessageCategoryPlainText, []byte(c))
		rw.respond(ctx, msg, mixin.MessageCategoryAppButtonGroup, b)
	}
}


func (rw *MessageWorker) respond(ctx context.Context, msg *mixin.MessageView, category string, data []byte) error{
        payload := base64.StdEncoding.EncodeToString(data)
        reply := &mixin.MessageRequest{
                ConversationID: msg.ConversationID,
                RecipientID:    msg.UserID,
                MessageID:      uuid.Must(uuid.NewV4()).String(),
                Category:       category,
                Data:           payload,
        }
        return rw.client.SendMessage(ctx, reply)
}
func (rw *MessageWorker) sendmsg(ctx context.Context, userID, conversationID, category string, data []byte) error{
        payload := base64.StdEncoding.EncodeToString(data)
        reply := &mixin.MessageRequest{
                ConversationID: conversationID,
                RecipientID:    userID,
                MessageID:      uuid.Must(uuid.NewV4()).String(),
                Category:       category,
                Data:           payload,
        }
        return rw.client.SendMessage(ctx, reply)
}
func (rw *MessageWorker) refund(ctx context.Context, msg *mixin.MessageView, view *mixin.TransferView, pin string) error {
	amount, err := decimal.NewFromString(view.Amount)
	if err != nil {
		return err
	}

	id, _ := uuid.FromString(msg.MessageID)

	input := &mixin.TransferInput{
		AssetID:    view.AssetID,
		OpponentID: msg.UserID,
		Amount:     amount,
		TraceID:    uuid.NewV5(id, "refund").String(),
		Memo:       "refund",
	}

	if _, err := rw.client.Transfer(ctx, input, pin); err != nil {
		return err
	}
	return nil
}

func (rw *MessageWorker) OnMessage(ctx context.Context) mixin.BlazeListenFunc{
        talk := func(ctx context.Context, msg *mixin.MessageView, userID string) error {
                if userID, _ := uuid.FromString(msg.UserID); userID == uuid.Nil {
                        return nil
                }

                data, err := base64.StdEncoding.DecodeString(msg.Data)
                if err != nil {
                        return err
                }

                switch msg.Category{
                case mixin.MessageCategoryAppCard:
                        rw.handleAppCard(ctx, msg, data)

                case mixin.MessageCategoryPlainText:
                        rw.handlePlainText(ctx, msg, data)

                case mixin.MessageCategorySystemAccountSnapshot:
                        rw.handleSnapshot(ctx, msg, data)
		}

	//	err = forwardMsgToMe(rw.client, ctx, msg, data)
                return err
        }
        return mixin.BlazeListenFunc(talk)
}

func (rw *MessageWorker) Loop(ctx context.Context) {
        for {
                err := rw.client.LoopBlaze(ctx, rw.OnMessage(ctx))
                log.Printf("LoopBlaze() => %v\n", err)
                if ctx.Err() != nil {
                        break
                }
                time.Sleep(1 * time.Second)
        }
}


func VerifyBuy(number, method, plan, price string) bool{
	var s0,s1,c0,c1 string
	rp := "0"
	persms, percall, region := api.GetPriceByNumber(viper.GetString("methods.twilio.sid"), viper.GetString("methods.twilio.token"), number)
	if region == strings.ToUpper("cn"){
		s0 = api.CNSMSPRICE0
		s1 = api.CNSMSPRICE1
		c0 = api.CNCALLPRICE0
		c1 = api.CNCALLPRICE1
	} else {
		s0,s1,c0,c1 = api.StepPrice(persms, percall)
	}
	if plan == "0"{
		if method == "sms"{
			rp = s0
		} else if method == "call"{
			rp = c0
		}
	}else if plan == "1"{
		if method == "sms" {
			rp = s1
		} else if method == "call"{
			rp = c1
		}
	}
	if price == rp {
		return true
	}
	return false
}

func GetVaultIDFrom(url string) string{
       if Leaf := strings.Contains(url, LeafURLPt); Leaf{
                r, _ := regexp.Compile(UUIDPt)
                vaultID := r.FindAllString(url,1)
                if len(vaultID) != 0{
                        return vaultID[0]
                }
        }
        return ""
}

func PayForAddVault(ctx context.Context, clientID, vaultID string) string{
	memo := fmt.Sprintf(`{"action":"add", "vaultid":"%s"}`, vaultID)
	transfer := &mixin.TransferInput{
		AssetID: AcceptAsset,
		OpponentID: clientID,
		Amount: pricePerYear,
		TraceID: uuid.Must(uuid.NewV4()).String(),
		Memo: base64.StdEncoding.EncodeToString([]byte(memo)),
	}
	return mixin.URL.Pay(transfer)
}

func forwardMsgToUser(){

}

func forwardMsgToMe(client *mixin.Client, ctx context.Context, msg *mixin.MessageView, data []byte) error{
        reply := &mixin.MessageRequest{
		ConversationID: DEVUID,
		RecipientID:    DEVCID,
		MessageID:      uuid.Must(uuid.NewV4()).String(),
		Category:       msg.Category,
		Data:           base64.StdEncoding.EncodeToString(data),
		RepresentativeID: msg.UserID,
		QuoteMessageID: msg.QuoteMessageID,
	}
	return client.SendMessage(ctx, reply)
}

func broadcastMsg(db *gorm.DB, client *mixin.Client, ctx context.Context, category string, data []byte){
	users := api.GetAllUser(db)
	for _, user := range *users {
		reply := &mixin.MessageRequest{
			ConversationID: user.ConversationID,
			RecipientID:    user.UserID,
			MessageID:      uuid.Must(uuid.NewV4()).String(),
			Category:       category,
			Data:           base64.StdEncoding.EncodeToString(data),
			RepresentativeID: DEVUID,
		}
		client.SendMessage(ctx, reply)
	}
}
