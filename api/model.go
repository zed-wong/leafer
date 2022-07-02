package api

const (
	TIMEFORMAT = "2006-01-02T15:04:05Z"
)

type VAULT struct{
        UserID  string `json:"user_id"`
        VaultID string `json:"vault_id"`
        IdentityID string `json:"identity_id";gorm:"primaryKey"`
	Avatar  string `json:"avatar"`
        Ratio   string `json:"ratio"`
        NextRatio string `json:"next_ratio"`
        AlertRatio      string `json:"alert_ratio"`
        AddAt   string `json:"add_at"`
        EndAt   string `json:"end_at"`
        Triggered       bool `json:"triggered"`
}

type USER struct{
	UserID string `json:"user_id";gorm:"primaryKey"`
	ConversationID string `json:"conversation_id"`
	IdentityID string `json:"identity_id"`
	Name string `json:"name"`
        Avatar  string `json:"avatar"`
	Lang string `json:"lang"`
	LastActive string `json:"last_active"`
	AccessToken string `json:"access_token"`
	TgID	string `json:"tg_id"`
	TgName	string `json:"tg_name"`
	SignalNumber string `json:"signal_number"`
	PhoneNumber string `json:"phone_number"`
	SMSBalance int `json:"sms_balance"`
	CallBalance int `json:"call_balance"`
}

type AlertHistory struct{
	UserID string `json:"user_id";gorm:"primaryKey"`
	Method string `json:"method"`
	AlertAt string `json:"alert_at"`
}

type PAYMENT struct {
	UserID string `json:"user_id";gorm:"primaryKey"`
	TraceID string `json:"trace_id"`
	Amount	string `json:"amount"`
	Action	string `json:"action"`
	Time	string `json:"time"`
}

type USERDATA struct{
	Vault   []VAULT `json:"vault"`
	User    USER	`json:"user"`
}

type PriceResult struct{
        Type string `json:"type"`
        Plan string `json:"plan"`
        Num string `json:"num"`
        Price string `json:"price"`
}
