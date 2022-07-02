package api

import (
        "gorm.io/gorm"
        "gorm.io/driver/postgres"
        "github.com/gin-gonic/gin"
        "github.com/gin-contrib/static"
)

var (
	DEBUG = false
	Path = "web/dist"
	TLSpem = "cert.pem"
	TLSkey = "cert.key"
)

type Config struct{
	ClientID string
	AppSecret string
	TgToken string
	TwilioSid string
	TwilioToken string
}

type APIWorker struct{
	r	*gin.Engine
	db	*gorm.DB
	c	*Config
}

func NewAPIWorker(dsn string, c *Config) *APIWorker{
	if !DEBUG{
		gin.SetMode(gin.ReleaseMode)
	}
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
                panic(err)
        }

	return &APIWorker{
		r: gin.Default(),
		db: db,
		c: c,
	}
}

func (aw *APIWorker) Loop(host string){
	// PAGES
        aw.r.Use(static.Serve("/",static.LocalFile(Path,false)))
	aw.r.Use(static.Serve("/oauth",static.LocalFile(Path,false)))
	aw.r.Use(static.Serve("/methods",static.LocalFile(Path,false)))
	aw.r.Use(static.Serve("/login",static.LocalFile(Path,false)))
	aw.r.Use(static.Serve("/help",static.LocalFile(Path,false)))
	aw.r.Use(static.Serve("/questions",static.LocalFile(Path,false)))
	aw.r.Use(static.Serve("/zh",static.LocalFile(Path,false)))

	aw.r.Use(static.Serve("/__webpack_hmr/client",static.LocalFile(Path,false)))
	aw.r.Use(static.Serve("/_nuxt/*",static.LocalFile(Path,false)))
	
	// API
	aw.r.GET("/mixinoauth", MixinOauthHandler(aw.db, aw.c.ClientID, aw.c.AppSecret))
	aw.r.GET("/poll", PollHandler(aw.db))
	aw.r.PUT("/update/tg",TelegramHandler(aw.db))
	aw.r.PUT("/update/ratio",UpdateRatioHandler(aw.db))
	aw.r.PUT("/update/number", NumberHandler(aw.db))
	aw.r.DELETE("/delete/vault",DeleteVaultHandler(aw.db))
	aw.r.GET("/price",PriceHandler(aw.c.TwilioSid, aw.c.TwilioToken))
	// aw.r.Run(host)
	aw.r.RunTLS(host, TLSpem, TLSkey)
}
