// binding for website
package api

import (
	"fmt"
	"log"
	"time"
	"strings"
	"context"
        "gorm.io/gorm"
        "github.com/tidwall/gjson"
        "github.com/gin-gonic/gin"
        "github.com/go-resty/resty/v2"
	"github.com/shopspring/decimal"
	"github.com/fox-one/mixin-sdk-go"
)

var (
	CNSMSPRICE0="0.5"
	CNCALLPRICE0="0.9"
	CNSMSPRICE1="5"
	CNCALLPRICE1="9"

	PRICESTEP0=decimal.NewFromFloat(0.025)
	REALPRICE0=decimal.NewFromFloat(0.05)
	PRICESTEP1=decimal.NewFromFloat(0.05)
	REALPRICE1=decimal.NewFromFloat(0.1)
	PRICESTEP2=decimal.NewFromFloat(0.1)
	REALPRICE2=decimal.NewFromFloat(0.15)
	PRICESTEP3=decimal.NewFromFloat(0.15)
	REALPRICE3=decimal.NewFromFloat(0.3)
	PRICESTEP4=decimal.NewFromFloat(0.3)
	REALPRICE4=decimal.NewFromFloat(0.5)

	PLAN1NUM = "8"
	PLAN2NUM = "100"
	PLAN1TIMES=decimal.NewFromInt(10)
	PLAN2TIMES=decimal.NewFromInt(100)
)

func MixinOauthHandler(db *gorm.DB, clientID, appSecret string) gin.HandlerFunc{
	return func(c *gin.Context){
		code := c.Query("code")
		conversationID := c.Query("conversation_id")
		lang := c.Query("lang")

		if len(code) == 64{
			ctx := context.Background()
			accessToken, _, err := mixin.AuthorizeToken(ctx, clientID, appSecret, code, "")
			if err != nil{
				fmt.Println(err)
			}
			if len(accessToken) != 0 {
				user, err := mixin.UserMe(ctx, accessToken)
				if err != nil{
					fmt.Println(err)
				}
				if len(conversationID) == 0 {
					conversationID = mixin.UniqueConversationID(clientID , user.UserID)
				}
				lastActive := time.Now().Format(TIMEFORMAT)
				User := &USER{
					UserID: user.UserID,
					ConversationID: conversationID,
					IdentityID: user.IdentityNumber,
					Name: user.FullName,
					Avatar: user.AvatarURL,
					Lang: lang,
					LastActive: lastActive,
					AccessToken: accessToken,
					SignalNumber: user.Phone,
					PhoneNumber: user.Phone,
				}
				UpdateOrAddUser(db, *User)
				c.IndentedJSON(200, GetUserData(db, user.UserID))
			} else{
				c.JSON(500, gin.H{
					"code": 500,
					"error": "Internal Server Error",
				})
			}
		} else{
			c.JSON(401, gin.H{
				"code": 401,
				"error": "Invalid Code",
			})
		}
	}
}

func getTokenFromHeader(c *gin.Context) (string, string){
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"code":  401,
			"error": "Invalid Authorization Token",
		})
		c.Abort()
		return "", ""
	}
	parts := strings.SplitN(authorization, " ", 2)
	if parts[0] != "Bearer" || len(parts[1]) == 0 {
		c.JSON(401, gin.H{
			"code":  401,
			"error": "Invalid Authorization Token",
		})
		c.Abort()
		return "", ""
	}
	uid := c.Request.Header.Get("UserID")

	return parts[1], uid
}

func authMiddleware(c *gin.Context, db *gorm.DB) bool{
	token, userID := getTokenFromHeader(c)
	if len(token) == 0 || len(userID) == 0{
		log.Println("token == 0")
		return false
	}
	if VerifyToken(db, token, userID){
		return true
	}
	return false
}

// HTTP GET
func PollHandler(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		if authMiddleware(c, db){
			_, userID := getTokenFromHeader(c)
			c.IndentedJSON(200, GetUserData(db, userID))
		}else {
			c.IndentedJSON(401, gin.H{
				"code": 401,
				"error": "Invalid token",
			})
		}
	}
}

// HTTP GET
func PollBuyHandler(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		traceID := c.Query("trace_id")
		if len(traceID) == 0{
			c.IndentedJSON(401, gin.H{
				"code": 401,
				"error": "Invalid Trace",
			})
			return
		}
		state := GetPaymentState(db, traceID)
		c.IndentedJSON(200, gin.H{
			"code": 200,
			"state": state,
		})
	}
}

// HTTP PUT
func UpdateRatioHandler(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		if authMiddleware(c, db){
			_, userID := getTokenFromHeader(c)
			body, err := c.GetRawData()
			if err != nil{
				log.Println("c.GetRawData() => ",err)
			}
			identityID := gjson.Get(string(body), "identity_id").String()
			ratio := gjson.Get(string(body), "ratio").String()
			if UpdateVaultRatio(db, identityID, userID, ratio){
				c.JSON(200, gin.H{
					"code":  200,
					"error": "UpdateRatio success",
				})
			} else {
				c.JSON(401, gin.H{
					"code":  401,
					"error": "UpdateRatio failed",
				})
			}
		}
	}
}

// HTTP DELETE
func DeleteVaultHandler(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		if authMiddleware(c, db){
			_, userID := getTokenFromHeader(c)
			body, err := c.GetRawData()
			if err != nil{
				log.Println("c.GetRawData() => ",err)
			}

			identityID := gjson.Get(string(body), "identity_id").String()
			if DeleteVault(db, identityID, userID){
				c.JSON(200, gin.H{
					"code":  200,
					"error": "DeleteVault success",
				})
			} else {
				c.JSON(401, gin.H{
					"code":  401,
					"error": "DeleteVault failed",
				})
			}
		}
	}
}

// HTTP PUT 
func TelegramHandler(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		if authMiddleware(c, db){
			_, userID := getTokenFromHeader(c)
			body, err := c.GetRawData()
			if err != nil{
				log.Println("c.GetRawData() => ",err)
			}
			tgID := gjson.Get(string(body), "tg_id").String()
			tgName := gjson.Get(string(body), "tg_name").String()
			if SetTg(db, tgID, tgName, userID){
				c.JSON(200, gin.H{
					"code":  200,
					"error": "SetNumber success",
				})
			} else {
				c.JSON(500, gin.H{
					"code":  500,
					"error": "SetTelegram error",
				})
			}
		}
	}
}

// HTTP PUT
func NumberHandler(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		if authMiddleware(c, db){
			_, userID := getTokenFromHeader(c)
			body, err := c.GetRawData()
			if err != nil{
				log.Println("c.GetRawData() => ",err)
			}

			action := gjson.Get(string(body), "action").String()
			number := gjson.Get(string(body), "number").String()

			if SetNumber(db, action, number, userID){
				c.JSON(200, gin.H{
					"code":  200,
					"error": "SetNumber success",
				})
			} else {
				c.JSON(500, gin.H{
					"code":  500,
					"error": "SetNumber error",
				})
			}
		}
	}
}

// HTTP GET
func PriceHandler(twilioSid, twilioToken string) gin.HandlerFunc{
	if len(twilioSid) == 0 || len(twilioToken) == 0 {
		panic("Twilio token invalid")
	}
	return func(c *gin.Context){
		var smsprice0,smsprice1,callprice0,callprice1 string
		number := c.Query("number")
		if number == ""{
			c.JSON(400, gin.H{
				"code": 400,
				"error": "Invalid parameters",
			})
			return
		}

		perSmsPrice, perCallPrice, region := GetPriceByNumber(twilioSid, twilioToken, number)

		if perSmsPrice == "" || perCallPrice == ""{
			c.JSON(400, gin.H{
				"code": 400,
				"error": "Invalid number ",
			})
			return
		}

		if region == "CN"{
			smsprice0, smsprice1, callprice0, callprice1 = CNSMSPRICE0, CNSMSPRICE1, CNCALLPRICE0, CNCALLPRICE1
		} else {
			smsprice0, smsprice1, callprice0, callprice1 = StepPrice(perSmsPrice, perCallPrice)
		}

		result := &[]PriceResult{
			PriceResult {
				Type: "sms",
				Plan: "0",
				Num:  PLAN1NUM,
				Price: smsprice0,
			},
			PriceResult {
				Type: "sms",
				Plan: "1",
				Num:  PLAN2NUM,
				Price: smsprice1,
			},
			PriceResult {
				Type: "call",
				Plan: "0",
				Num:  PLAN1NUM,
				Price: callprice0,
			},
			PriceResult {
				Type: "call",
				Plan: "1",
				Num:  PLAN2NUM,
				Price: callprice1,
			},
		}

		c.JSON(200, *result)
	}
}

func GetPriceByNumber(twilioSid, twilioToken, number string) (string, string, string){
	callurl := fmt.Sprintf("https://pricing.twilio.com/v2/Voice/Numbers/%s", number)
	client := resty.New()
	callresp, _ := client.R().
	SetBasicAuth(twilioSid, twilioToken).
	Get(callurl)

	region := gjson.Get(callresp.String(), "iso_country").String()

	smsurl := fmt.Sprintf("https://pricing.twilio.com/v1/Messaging/Countries/%s", strings.ToUpper(region))
	smsresp, _ := client.R().
	SetBasicAuth(twilioSid, twilioToken).
	Get(smsurl)

	perSmsPrice := gjson.Get(smsresp.String(), "outbound_sms_prices.0.prices.0.current_price").String()
	perCallPrice := gjson.Get(callresp.String(), "outbound_call_prices.0.current_price").String()
	return perSmsPrice, perCallPrice, region
}

func StepPrice(perSMS, perCall string) (string, string, string, string){
	var realsms, realcall decimal.Decimal
	var smsprice0, smsprice1, callprice0, callprice1 string
	sms, _ := decimal.NewFromString(perSMS)
	call, _ := decimal.NewFromString(perCall)

	if sms.LessThanOrEqual(PRICESTEP0){
		realsms = REALPRICE0
	} else if sms.LessThanOrEqual(PRICESTEP1){
		realsms = REALPRICE1
	} else if sms.LessThanOrEqual(PRICESTEP2){
		realsms = REALPRICE2
	} else if sms.LessThanOrEqual(PRICESTEP3){
		realsms = REALPRICE3
	} else if sms.LessThanOrEqual(PRICESTEP4){
		realsms = REALPRICE4
	}

	if call.LessThanOrEqual(PRICESTEP0){
		realcall = REALPRICE0
	} else if call.LessThanOrEqual(PRICESTEP1){
		realcall = REALPRICE1
	} else if call.LessThanOrEqual(PRICESTEP2){
		realcall = REALPRICE2
	} else if call.LessThanOrEqual(PRICESTEP3){
		realcall = REALPRICE3
	} else if call.LessThanOrEqual(PRICESTEP4){
		realcall = REALPRICE4
	}

	smsprice0 = realsms.Mul(PLAN1TIMES).String()
	smsprice1 = realsms.Mul(PLAN2TIMES).String()
	callprice0 = realcall.Mul(PLAN1TIMES).String()
	callprice1 = realcall.Mul(PLAN2TIMES).String()
	return smsprice0, smsprice1, callprice0, callprice1
}
