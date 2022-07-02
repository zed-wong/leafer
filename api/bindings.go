package api

import(
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
	"github.com/tidwall/gjson"
	"github.com/go-resty/resty/v2"
)

// vault
func FetchVaultData(vaultID string) *VAULT{
	client := resty.New()
	respvat, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/vats/%s", vaultID))
	if err != nil {
		log.Println(err)
	}
	vault := gjson.Get(respvat.String(), "data")
	art := vault.Get("art").Float()		// art*rate = debt
	ink := vault.Get("ink").Float()		// Amount of coll(BTC)
	catID := vault.Get("collateral_id").String()
	identityID := vault.Get("identity_id").String()

	respcat, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/cats/%s", catID))
	if err != nil {
		log.Println(err)
	}
	collateral := gjson.Get(respcat.String(), "data")
	rate := collateral.Get("rate").Float()	// art*rate = debt
	price := collateral.Get("price").Float()// Current price
	mat := collateral.Get("mat").Float()	// Minium ratio
	gem := collateral.Get("gem").String()	// coll (BTC) UUID

	respasset, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/assets/%s", gem))
	if err != nil {
		log.Println(err)
	}
	avatar := gjson.Get(respasset.String(), "data.logo").String()	// Avatar url

	oracleresp, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/oracles/%s", gem))
	if err != nil {
		log.Println(err)
	}
	nextPrice := gjson.Get(oracleresp.String(), "data.next").Float()

	ratio := fmt.Sprintf("%.2f", (ink*price)/(art*rate)*100)
	nextRatio := fmt.Sprintf("%.2f", (ink*nextPrice)/(art*rate)*100)
	alertRatio := fmt.Sprintf("%.0f", mat*1.2*100)

	return &VAULT{
		VaultID: vaultID,
		IdentityID: identityID,
		Avatar: avatar,
		Ratio: ratio,
		NextRatio: nextRatio,
		AlertRatio: alertRatio,
		Triggered: false,
	}
}

func AddVault(db *gorm.DB, vaultID, userID string) bool{
	vault := FetchVaultData(vaultID)
	now := time.Now()
	vault.UserID = userID
	vault.AddAt = now.Format(TIMEFORMAT)
	vault.EndAt = now.AddDate(1,0,0).Format(TIMEFORMAT)

	return check(db.Create(&vault), "AddVault")
}

func IsAddedVault(db *gorm.DB, vaultID string) bool{
	var exist bool
	db.Model(&VAULT{}).
		Select("count(*) > 0").
		Where("vault_id = ?", vaultID).
		Find(&exist)
	return exist
}

func CheckFirstVault(db *gorm.DB, userID string) bool{
	var count int64
	db.Model(&VAULT{}).Where("user_id = ?",userID).Count(&count)
	if count == 0{
		return true
	}
	return false
}

func UpdateVaultRatio(db *gorm.DB, identityID, userID, ratio string) bool{
	if !isVaultOwner(db, identityID, userID){
		return false
	}
	return check(db.Model(&VAULT{}).Where("identity_id = ?", identityID).Update("alert_ratio", ratio), "UpdateVaultRatio")
}

func RenewVault(db *gorm.DB, identityID string) bool{
	var EndAt string
	vault := GetVault(db, identityID)
	endAt, _ := time.Parse(TIMEFORMAT, vault.EndAt)
	fmt.Println(endAt)
	if endAt.After(time.Now()){
		EndAt = endAt.AddDate(1,0,0).Format(TIMEFORMAT)
	} else {
		EndAt = time.Now().AddDate(1,0,0).Format(TIMEFORMAT)
	}
	fmt.Println(EndAt)

	return check(db.Model(&VAULT{}).Where("identity_id = ?", identityID).Update("end_at", EndAt), "RenewVault")
}

func DeleteVault(db *gorm.DB, identityID, userID string) bool{
	if !isVaultOwner(db, identityID, userID){
		return false
	}
	return check(db.Where("identity_id = ?", identityID).Delete(&VAULT{}), "DeleteVault")
}

func isVaultOwner(db *gorm.DB, identityID, userID string) bool{
	var vault VAULT
	db.Find(&vault, "identity_id = ?", identityID)
	if vault.UserID == userID{
		return true
	}
	return false
}

func check(result *gorm.DB, action string) bool{
	if result.Error != nil {
		log.Printf("%s() => %s", action, result.Error)
		return false
	} else if result.RowsAffected == 0 {
		return false
	}
	return true
}

// methods
func SetTg(db *gorm.DB, tgID, tgName, userID string) bool{
	return check(db.Model(&USER{}).Where("user_id", userID).Update("tg_id", tgID).Update("tg_name", tgName), "SetTg" )
}

func SetNumber(db *gorm.DB, method, number, userID string) bool{
	switch method{
	case "signal":
		return check(db.Model(&USER{}).Where("user_id", userID).Update("signal_number", number), "SetSignalNumber")
	case "phone":
		return check(db.Model(&USER{}).Where("user_id", userID).Update("phone_number", number), "SetPhoneNumber")
	}
	return false
}

func BuyService(db *gorm.DB, method string, num int, userID string) bool{
	var user USER
	db.Find(&user, "user_id = ?", userID)

	switch method{
	case "sms":
		return check(db.Model(&user).Where("user_id = ?", userID).Update("sms_balance", user.SMSBalance + num), "BuyService")
	case "call":
		return check(db.Model(&user).Where("user_id = ?", userID).Update("call_balance", user.CallBalance + num), "BuyService")
	}
	return false
}

func CostService(db *gorm.DB, method string, num int, userID string) bool{
	var user USER
	db.Find(&user, "user_id = ?", userID)

	switch method{
	case "sms":
		return check(db.Model(&user).Where("user_id = ?", userID).Update("sms_balance", user.SMSBalance - num), "CostService")
	case "call":
		return check(db.Model(&user).Where("user_id = ?", userID).Update("call_balance", user.CallBalance - num), "CostService")
	}
	return false
}

func AddHistory(db *gorm.DB, method, alertAt, userID string) bool{
	history := &AlertHistory{
		UserID: userID,
		Method: method,
		AlertAt: alertAt,
	}
	return check(db.Create(history), "AddHistory")
}

func AddPayment(db *gorm.DB, userID, traceID, amount, action string) bool{
	payment := &PAYMENT{
		UserID: userID,
		TraceID: traceID,
		Amount: amount,
		Action: action,
		Time: time.Now().Format(TIMEFORMAT),
	}
	return check(db.Create(payment), "AddPayment")
}

// Oauth
func UpdateOrAddUser(db *gorm.DB, user USER) bool{
	updateUser := &USER{
		UserID: user.UserID,
		ConversationID: user.ConversationID,
		IdentityID: user.IdentityID,
		Name: user.Name,
		Avatar: user.Avatar,
		Lang: user.Lang,
		LastActive: user.LastActive,
		AccessToken: user.AccessToken,
	}
	if len(user.UserID) != 0 {
		if UserExist(db, user.UserID){
			db.Model(&user).Where("user_id = ?", user.UserID).Updates(&updateUser)
		} else{
			db.Create(&user)
		}
		return true
	}
	return false
}

func AddUser(db *gorm.DB, user USER) bool{
	return check(db.Create(&user), "AddUser")
}

func VerifyToken(db *gorm.DB, token, userID string) bool{
	var user USER
	db.Find(&user, "user_id = ?", userID)
	if user.AccessToken == token {
		return true
	}
	return false
}

func GetUserData(db *gorm.DB, userID string) *USERDATA{
	var (
		user USER
		vault []VAULT
	)
	db.Find(&user, "user_id = ?", userID)
	db.Find(&vault, "user_id = ?", userID)

	return &USERDATA {
		User: user,
		Vault: vault,
	}
}

func GetPaymentState(db *gorm.DB, traceID string) bool{
	var exist bool
	db.Model(&PAYMENT{}).
		Select("count(*) > 0").
		Where("trace_id = ?", traceID).
		Find(&exist)
	return exist
}

// Workers
func GetAllVaults(db *gorm.DB) *[]VAULT{
	var vault *[]VAULT
	db.Find(&vault)
	return vault
}

func GetVault(db *gorm.DB, identityID string) *VAULT{
	var vault *VAULT
	db.Find(&vault, "identity_id = ?", identityID)
	return vault
}

func UpdateVault(db *gorm.DB, vault VAULT) bool{
//	result := db.Model(&VAULT{}).Where("vault_id = ?", vault.VaultID).Updates(VAULT{Ratio: vault.Ratio, Triggered: vault.Triggered})
	result := db.Table("vaults").Where("vault_id = ?", vault.VaultID).Updates(map[string]interface{}{"ratio": vault.Ratio, "triggered": vault.Triggered})
	return check(result, "UpdateVault")
}

func GetUser(db *gorm.DB, userID string) *USER{
	var user *USER
	db.Find(&user, "user_id = ?", userID)
	return user
}

func GetAllUser(db *gorm.DB) *[]USER{
	var users *[]USER
	db.Find(&users)
	return users
}

func UserExist(db *gorm.DB, userID string) bool{
	var exist bool
	db.Model(&USER{}).
		Select("count(*) > 0").
		Where("user_id = ?", userID).
		Find(&exist)
	return exist
}

func IsBought(db *gorm.DB, userID string) bool{
	var exist bool
	db.Model(&PAYMENT{}).
		Select("count(*) > 0").
		Where("action = 'buy' AND user_id = ?", userID).
		Find(&exist)
	if exist{
		return true
	}
	return false
}
