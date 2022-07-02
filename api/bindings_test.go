package api

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//	"github.com/gofrs/uuid"
	"github.com/zed-wong/leafer/api"
)

func initDB() *gorm.DB {
	const (
		host     = "localhost"
		user     = ""
		password = ""
		dbname   = ""
		port     = ""
		sslmode  = "disable"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func TestGetPaymentState(t *testing.T) {
	db := initDB()
	traceID := "44d9717d-8cae-4004-98a1-f9ad544dcfb1"
	t.Log(api.GetPaymentState(db, traceID))
}

func TestSetTg(t *testing.T) {
	db := initDB()
	t.Log(api.SetTg(db, "460545", "zed_wong", "44d9717d-8cae-4004-98a1-f9ad544dcfb1"))
}

func TestSetNumber(t *testing.T) {
	db := initDB()
	result := api.SetNumber(db, "phone", "+12013323", "44d9717d-8cae-4004-98a1-f9ad544dcfb1")
	t.Log(result)
}

func TestUpdateRatio(t *testing.T) {
	db := initDB()
	api.UpdateVaultRatio(db, "1", "44d9717d-8cae-4004-98a1-f9ad544dcfb1", "360")
}
func TestGetUser(t *testing.T) {
	db := initDB()
	user := api.GetUser(db, "cfe2c9bf-6a38-42d0-951d-7a91e7926811")
	fmt.Println(user, "UserId:", user.UserID)

	/*
		users := api.GetAllUser(db)
		fmt.Println(users)
		userID := "44d9717d-8cae-4004-98a1-f9ad544dcfb1"
		data := api.GetUserData(db, userID)
		t.Logf("UserData: %+v", data)
	*/
}

func TestGetAllUser(t *testing.T) {
	//	db := initDB()
}
func TestCheckFirstVault(t *testing.T) {
	db := initDB()
	fmt.Println(api.CheckFirstVault(db, "44d9717d-8cae-4004-98a1-f9ad544dcfb1"))
}
func TestFetchData(t *testing.T) {
	id := "8f8f0e19-4630-358c-bbb3-34acf6ea2137"
	t.Logf("%+v", api.FetchVaultData(id))
}

func TestAddNewUser(t *testing.T) {
	db := initDB()
	user := api.USER{
		UserID:         "44d9717d-8cae-4004-98a1-f9ad544dcfb1",
		ConversationID: "44d9717d-8cae-4004-98a1-f9ad544dcfb1",
		IdentityID:     "28211",
		Name:           "KINS",
		Lang:           "zh",
		LastActive:     time.Now().Format(api.TIMEFORMAT),
		SignalNumber:   "+8613326300013",
		PhoneNumber:    "+8613326300223",
		SMSBalance:     0,
		CallBalance:    0,
	}
	t.Log("Added user:", api.AddUser(db, user))
}

func TestUserData(t *testing.T) {
	db := initDB()
	userID := "44d9717d-8cae-4004-98a1-f9ad544dcfb1"
	data := api.GetUserData(db, userID)
	t.Logf("UserData: %+v", data)
}

func TestAddVault(t *testing.T) {
	db := initDB()
	t.Log("AddVault:", api.AddVault(db, "52010fd0-6a9d-393c-abcb-ca997d950bf5", "44d9717d-8cae-4004-98a1-f9ad544dcfb1"))
	t.Log("AddVault:", api.AddVault(db, "321b4903-c291-39e4-9563-6b01c4826b08", "44d9717d-8cae-4004-98a1-f9ad544dcfb1"))
	t.Log("AddVault:", api.AddVault(db, "6fbd8796-3ffa-38ab-aa95-64584e6e57ca", "fcb87491-4fa0-4c2f-b387-262b63cbc112"))
	t.Log("AddVault:", api.AddVault(db, "321b4903-c291-39e4-9563-6b01c4826b08", "fcb87491-4fa0-4c2f-b387-262b63cbc112"))
}

func TestRenewVault(t *testing.T) {
	db := initDB()
	api.RenewVault(db, "1")
}

func TestIsBought(t *testing.T) {
	db := initDB()
	t.Log(api.IsBought(db, "e84f63bd-a2f5-4690-a560-a9858bc3209c"))
}
