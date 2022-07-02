package workers

import (
	"fmt"
	"log"
	"time"
	"context"
	"encoding/base64"
        "gorm.io/gorm"
        "gorm.io/driver/postgres"
	"github.com/spf13/viper"
	"github.com/gofrs/uuid"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/zed-wong/leafer/api"
)

type NotifierWorker struct{
	db	*gorm.DB
	mw	*MethodWorker
}

var (
	RENEWPERIOD = time.Hour * 24 * 30
	MINCHARGEBALANCE = 5
	MethodPageURL = "https://leafer.one/methods"
	ButtonColor = "#5979F0"
	AlertRenewAsset = AcceptAsset
)

func NewNotifierWorker(dsn string, mw *MethodWorker) *NotifierWorker{
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
                panic(err)
        }
        return &NotifierWorker{
                db: db,
		mw: mw,
        }
}

// Notify user to renew vault everyday in 30 days before expire
func (nw *NotifierWorker) AlertRenewVaults() {
	vaults := api.GetAllVaults(nw.db)
	if len(*vaults) == 0{
		return
	}
	for _, vault := range *vaults{
		if isUnderPeriod(&vault){
			user := api.GetUser(nw.db, vault.UserID)
			content := api.GetContentByLang(user.Lang)
			tm, _ :=  time.Parse(api.TIMEFORMAT, vault.EndAt)
			t := tm.Format("2006-01-02 15:04:05")
			c := fmt.Sprintf(content.PayRenewMsg, vault.IdentityID, t)
			nw.mw.Mixin(user.UserID, user.ConversationID, c)
			nw.mw.MixinBtn(user.UserID, user.ConversationID, content.PayRenew, PayForRenewVault(context.Background(), viper.GetString("methods.mixin.client_id"), vault.IdentityID), ButtonColor)
		}
	}
}

// Notify user to buy more sms and call when balance in (0 > balance >= 5)
func (nw *NotifierWorker) AlertChargeService() {
	users := api.GetAllUser(nw.db)
	if len(*users) == 0{
		return
	}
	for _, user := range *users{
		content := api.GetContentByLang(user.Lang)
		if api.IsBought(nw.db, user.UserID){
			if user.SMSBalance < MINCHARGEBALANCE || user.CallBalance < MINCHARGEBALANCE{
				c := fmt.Sprintf(content.ToMethodPageMsg, user.SMSBalance, user.CallBalance)
				nw.mw.Mixin(user.UserID, user.ConversationID, c)
				nw.mw.MixinBtn(user.UserID, user.ConversationID, content.ToMethodPage, MethodPageURL, ButtonColor)
			}
		}
	}
}

func isUnderPeriod(vault *api.VAULT) bool{
	endAt, err := time.Parse(api.TIMEFORMAT, vault.EndAt)
	if err != nil{
		log.Println("notifier.isUnderPeriod() =>", err)
	}
	return time.Now().Add(RENEWPERIOD).After(endAt)
}

func PayForRenewVault(ctx context.Context, clientID, identityID string) string{
        memo := fmt.Sprintf(`{"action":"renew", "identityid":"%s"}`, identityID)
        transfer := &mixin.TransferInput{
                AssetID: AlertRenewAsset,
                OpponentID: clientID,
                Amount: pricePerYear,
                TraceID: uuid.Must(uuid.NewV4()).String(),
                Memo: base64.StdEncoding.EncodeToString([]byte(memo)),
        }
        return mixin.URL.Pay(transfer)
}
