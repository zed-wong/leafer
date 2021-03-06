package api

const (
	// Alert Module
	ALERTMESSAGE_EN = "[Leafer] Your Pando Leaf vault ð¦#%s's current ratio is %s%%, Please add collateral in time!"
	TWILIOMESSAGE = "[Leafer] Your Pando Leaf vault is close to liquidation! (%s)"

	// Message Module
	HI_EN = "Welcome! Click the ð¤ button below to protect your Pando Leaf vaults."
	ADDSUCCESS_EN = "âï¸ Add vault success!\n- Valid until: %s"
	ADDFREESUCCESS_EN = "âï¸ Add vault success! This is your first vault, so it's free for a year."
	ADDFAILED_EN = "âï¸ Add vault failed. Please retry or contact 28865 for help."
	ADDEXISTED_EN = "âï¸ Add vault failed. This vault has been added."
	RENEWSUCCESS_EN = "âï¸ Renew vault success!\n - Valid until: %s"
	RENEWFAILED_EN = "âï¸ Renew vault failed."
	BUYSUCCESS_EN = "âï¸ Purchase success!\n\n -ð¬ SMS: %d\n -ð Call: %d"
	BUYFAILED_EN = "âï¸ Purchase failed."
	SMSSENT_EN = "âï¸ Test SMS sent."
	CALLSENT_EN = "âï¸ Test Call made."
	TESTFAILED_EN = "âï¸ Test failed."
	NOACTION_EN = "âï¸ Action Missing."
	PAYADDING_EN = "Click Here To Pay"
	PAYADDINGMSG_EN = "Adding this vault requires paying $12 per year. Please click the button below to pay."
	PAYRENEW_EN = "Click Here To Renew"
	PAYRENEWMSG_EN = "âï¸ Your vault #%s is about to expire in %s, Please renew as soon as possible."
	TOMETHODPAGE_EN = "Click to Buy"
	TOMETHODPAGEMSG_EN = "âï¸ The balance is about to be exhausted, please buy more to maintain operation!\n\n -ð¬ SMS:%d\n -ð Call:%d "
)

const (
	// Alert Module
	ALERTMESSAGE_ZH = "ãLeaferãæ¨å¨Pando Leafçéåº«ð¦#%sæ­£æ¥è¿å¹³åç·ï¼ç¶åæµæ¼ç:%s%%ï¼è«åæèçï¼"
	SMSBAOMESSAGE = "ãLeaferãæ¨å¨Pando Leafçè¨å®å·²ç¶è§¸ç¼æéï¼è«åæèçï¼(%s)"
	DINGXINAUDIO = "å°æ¬çæ½åº¦ç¨æ¶ï¼æ¨é¸ä¸­çååå·²ç¶éè³é æå¹æ ¼ï¼è«åææ¥çï¼"

	// Message Module
	HI_ZH = "æ­¡è¿ä½¿ç¨Leaferï¼é»æä¸æ¹æ©å¨äººð¤æéä¾ä¿è­·æ¨çPando Leaféåº«ã"
	ADDSUCCESS_ZH = "âï¸ æ·»å éåº«æå!\n- æææè³: %s"
	ADDFREESUCCESS_ZH = "âï¸ æ·»å éåº«æå! è©²éåº«ç²æ¨çé¦åéåº«ï¼äº«åä¸å¹´åè²»ä½¿ç¨ã"
	ADDFAILED_ZH = "âï¸ æ·»å éåº«å¤±æï¼è«éè©¦æè¯ç¹« 28865 å°æ±å¹«å©ã"
	ADDEXISTED_ZH = "âï¸ æ·»å å¤±æï¼è©²éåº«å·²å­å¨ã"
	RENEWSUCCESS_ZH = "âï¸ çºè²»éåº«æå!\n- æææè³: %s"
	RENEWFAILED_ZH = "âï¸ çºè²»éåº«å¤±æã"
	BUYSUCCESS_ZH = "âï¸ è³¼è²·æå!\n\n -ð¬ ç­ä¿¡: %d\n -ð é»è©±: %d"
	BUYFAILED_ZH = "âï¸ è³¼è²·å¤±æã"
	SMSSENT_ZH = "âï¸ æ¸¬è©¦ç­ä¿¡å·²ç¼éã"
	CALLSENT_ZH = "âï¸ æ¸¬è©¦é»è©±å·²æ¥åºã"
	TESTFAILED_ZH = "âï¸ æ¸¬è©¦å¤±æã"
	NOACTION_ZH = "âï¸ è«æ±æèª¤ã"
	PAYADDING_ZH = "è¼ææ­¤èä»æ¬¾"
	PAYADDINGMSG_ZH = "æ·»å æ­¤éåº«éè¦æ¯ä»12ç¾åæ¯å¹´ï¼è«é»æä¸æ¹æéæ¯ä»ã"
	PAYRENEW_ZH = "è¼ææ­¤èçºè²»"
	PAYRENEWMSG_ZH = "âï¸ æ¨çéåº«#%så°æ¼%så°æï¼è«åå¿«çºè²»ã"
	TOMETHODPAGE_ZH = "è¼ææ­¤èè³¼è²·"
	TOMETHODPAGEMSG_ZH = "âï¸ éç¥é¤é¡å³å°èç¡ï¼è«åæè³¼è²·ä»¥ç¶­æéç¥çæ­£å¸¸éä½!\n\n -ð¬ ç­ä¿¡:%d\n -ð é»è©±:%d"
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
