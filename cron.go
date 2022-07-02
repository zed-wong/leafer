package main

import(
	"github.com/robfig/cron/v3"
	"github.com/zed-wong/leafer/workers"
)

func InitCron(uw *workers.UpdaterWorker, mw *workers.MethodWorker, nw *workers.NotifierWorker){
	c := cron.New()
	update := func() { uw.Update() }
	check := func() { uw.CheckAndAlert(mw) }
	notify := func() { 
		nw.AlertRenewVaults()
	//	nw.AlertChargeService()  Need to Rewrite logic
	}

	c.AddFunc("* * * * *", update)
	c.AddFunc("*/5 * * * *", check)

	c.AddFunc("0 0 * * *", notify)
	c.Start()
}
