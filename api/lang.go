package api

const (
	// Alert Module
	ALERTMESSAGE_EN = "[Leafer] Your Pando Leaf vault 🏦#%s's current ratio is %s%%, Please add collateral in time!"
	TWILIOMESSAGE = "[Leafer] Your Pando Leaf vault is close to liquidation! (%s)"

	// Message Module
	HI_EN = "Welcome! Click the 🤖 button below to protect your Pando Leaf vaults."
	ADDSUCCESS_EN = "✔️ Add vault success!\n- Valid until: %s"
	ADDFREESUCCESS_EN = "✔️ Add vault success! This is your first vault, so it's free for a year."
	ADDFAILED_EN = "✖️ Add vault failed. Please retry or contact 28865 for help."
	ADDEXISTED_EN = "✖️ Add vault failed. This vault has been added."
	RENEWSUCCESS_EN = "✔️ Renew vault success!\n - Valid until: %s"
	RENEWFAILED_EN = "✖️ Renew vault failed."
	BUYSUCCESS_EN = "✔️ Purchase success!\n\n -💬 SMS: %d\n -📞 Call: %d"
	BUYFAILED_EN = "✖️ Purchase failed."
	SMSSENT_EN = "✔️ Test SMS sent."
	CALLSENT_EN = "✔️ Test Call made."
	TESTFAILED_EN = "✖️ Test failed."
	NOACTION_EN = "✖️ Action Missing."
	PAYADDING_EN = "Click Here To Pay"
	PAYADDINGMSG_EN = "Adding this vault requires paying $12 per year. Please click the button below to pay."
	PAYRENEW_EN = "Click Here To Renew"
	PAYRENEWMSG_EN = "❗️ Your vault #%s is about to expire in %s, Please renew as soon as possible."
	TOMETHODPAGE_EN = "Click to Buy"
	TOMETHODPAGEMSG_EN = "❗️ The balance is about to be exhausted, please buy more to maintain operation!\n\n -💬 SMS:%d\n -📞 Call:%d "
)

const (
	// Alert Module
	ALERTMESSAGE_ZH = "【Leafer】您在Pando Leaf的金庫🏦#%s正接近平倉線，當前抵押率:%s%%，請及時處理！"
	SMSBAOMESSAGE = "【Leafer】您在Pando Leaf的訂單已經觸發提醒，請及時處理！(%s)"
	DINGXINAUDIO = "尊敬的潘度用戶：您選中的商品已經降至預期價格，請及時查看！"

	// Message Module
	HI_ZH = "歡迎使用Leafer，點擊下方機器人🤖按鈕來保護您的Pando Leaf金庫。"
	ADDSUCCESS_ZH = "✔️ 添加金庫成功!\n- 有效期至: %s"
	ADDFREESUCCESS_ZH = "✔️ 添加金庫成功! 該金庫爲您的首個金庫，享受一年免費使用。"
	ADDFAILED_ZH = "✖️ 添加金庫失敗，請重試或聯繫 28865 尋求幫助。"
	ADDEXISTED_ZH = "✖️ 添加失敗，該金庫已存在。"
	RENEWSUCCESS_ZH = "✔️ 續費金庫成功!\n- 有效期至: %s"
	RENEWFAILED_ZH = "✖️ 續費金庫失敗。"
	BUYSUCCESS_ZH = "✔️ 購買成功!\n\n -💬 短信: %d\n -📞 電話: %d"
	BUYFAILED_ZH = "✖️ 購買失敗。"
	SMSSENT_ZH = "✔️ 測試短信已發送。"
	CALLSENT_ZH = "✔️ 測試電話已撥出。"
	TESTFAILED_ZH = "✖️ 測試失敗。"
	NOACTION_ZH = "✖️ 請求有誤。"
	PAYADDING_ZH = "輕按此處付款"
	PAYADDINGMSG_ZH = "添加此金庫需要支付12美元每年，請點擊下方按鈕支付。"
	PAYRENEW_ZH = "輕按此處續費"
	PAYRENEWMSG_ZH = "❗️ 您的金庫#%s將於%s到期，請儘快續費。"
	TOMETHODPAGE_ZH = "輕按此處購買"
	TOMETHODPAGEMSG_ZH = "❗️ 通知餘額即將耗盡，請及時購買以維持通知的正常運作!\n\n -💬 短信:%d\n -📞 電話:%d"
)


type LANG struct{
	Hi		string
	AlertMsg	string
	CensoredMsg	string
	AddSuccess	string
	AddFreeSuccess	string
	AddFailed	string
	AddExisted	string
	RenewSuccess	string
	RenewFailed	string
	BuySuccess	string
	BuyFailed	string
	SmsSent		string
	CallSent	string
	TestFailed	string
	NoAction	string
	PayAdding	string
	PayAddingMsg	string
	PayRenew	string
	PayRenewMsg	string
	ToMethodPage	string
	ToMethodPageMsg	string
}

var (
	EN = LANG {
		Hi: HI_EN,
		AlertMsg: ALERTMESSAGE_EN,
		CensoredMsg: TWILIOMESSAGE,
		AddSuccess: ADDSUCCESS_EN,
		AddFreeSuccess: ADDFREESUCCESS_EN,
		AddFailed: ADDFAILED_EN,
		AddExisted: ADDEXISTED_EN,
		RenewSuccess: RENEWSUCCESS_EN,
		RenewFailed: RENEWFAILED_EN,
		BuySuccess: BUYSUCCESS_EN,
		BuyFailed: BUYFAILED_EN,
		SmsSent: SMSSENT_EN,
		CallSent: CALLSENT_EN,
		TestFailed: TESTFAILED_EN,
		NoAction: NOACTION_EN,
		PayAdding: PAYADDING_EN,
		PayAddingMsg: PAYADDINGMSG_EN,
		PayRenew: PAYRENEW_EN,
		PayRenewMsg: PAYRENEWMSG_EN,
		ToMethodPage: TOMETHODPAGE_EN,
		ToMethodPageMsg: TOMETHODPAGEMSG_EN,
	}

	ZH = LANG {
		Hi: HI_ZH,
		AlertMsg: ALERTMESSAGE_ZH,
		CensoredMsg: SMSBAOMESSAGE,
		AddSuccess: ADDSUCCESS_ZH,
		AddFreeSuccess: ADDFREESUCCESS_ZH,
		AddFailed: ADDFAILED_ZH,
		AddExisted: ADDEXISTED_ZH,
		RenewSuccess: RENEWSUCCESS_ZH,
		RenewFailed: RENEWFAILED_ZH,
		BuySuccess: BUYSUCCESS_ZH,
		BuyFailed: BUYFAILED_ZH,
		SmsSent: SMSSENT_ZH,
		CallSent: CALLSENT_ZH,
		TestFailed: TESTFAILED_ZH,
		NoAction: NOACTION_ZH,
		PayAdding: PAYADDING_ZH,
		PayAddingMsg: PAYADDINGMSG_ZH,
		PayRenew: PAYRENEW_ZH,
		PayRenewMsg: PAYRENEWMSG_ZH,
		ToMethodPage: TOMETHODPAGE_ZH,
		ToMethodPageMsg: TOMETHODPAGEMSG_ZH,
	}
)

func GetContentByLang(lang string) *LANG{
	var l *LANG
	switch (lang){
	case "en":
		l = &EN
	case "zh":
		l = &ZH
	default:
		l = &EN
	}
	return l
}
