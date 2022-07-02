package api

import (
	"fmt"
	"testing"
	"gorm.io/gorm"
        "gorm.io/driver/postgres"
        "github.com/gin-gonic/gin"
	"github.com/zed-wong/leafer/api"
)
type Vault struct{
	ID string
	Ratio string
}
type User struct{
	ID string
	Conv string
}
type AllData struct{
	Vaultd []Vault
	Userd User
}

func TestGinReturnEmptyArray(t *testing.T){
	r := gin.Default()
	r.GET("/test", func(c *gin.Context){
		v := make([]Vault, 0)
		t.Log("v:",v)
		u := User{
			ID: "123321",
			Conv: "456784",
		}
		t.Log("u:",u)
		a := AllData{
			Userd: u,
			Vaultd: v,
		}
		c.JSON(200,a)
	})
	r.Run("127.0.0.1:8080")
}

func TestPriceHandler(t *testing.T){
	token := ""
	r := gin.Default()
	r.GET("/price", api.PriceHandler(token))
}
func TestOauthHandler(t *testing.T){
	db := initDB()
	r := gin.Default()
	clientID := "51186d7e-d488-417d-a031-b4e34f4fdf86"
	appSecret := "//"
	r.GET("/mixinoauth", api.MixinOauthHandler(db, clientID, appSecret))
	r.Run("127.0.0.1:8080")
}

func initDB() *gorm.DB{
        const (
                host = "localhost"
                user = ""
                password = ""
                dbname = ""
                port = ""
                sslmode = ""
        )
        dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
                panic(err)
        }
        return db
}
