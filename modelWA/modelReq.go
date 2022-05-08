package modelWA

import "time"

type ReqSendWA struct {
	Nomer     string `json:"nomer"`
	Text      string `json:"text"`
	OtpNumber string `json:"otp"`
}
type ReqSendDB struct {
	ClientID  string `json:"clientId"`
	MessageID string `json:"messageId"`
}
type ResSendDB struct {
	Count int    `json:"count"`
	Data  []Data `json:"data"`
}
type Data struct {
	Nomer int `json:"nomer"`
	Ke    int `json:"ke"`
}
type ReqSendWAWithImage struct {
	Nomer string `json:"nomer"`
	Text  string `json:"text"`
	Image string `json:"image"`
}
type ResGlobal struct {
	Status         string      `json:"status"`
	StatusDateTime time.Time   `json:"statusDateTime"`
	StatusDesc     string      `json:"statusDesc"`
	Result         interface{} `json:"result"`
}
type ResSendWA struct {
	ID    string `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	Text  string `json:"text"`
	Image string `json:"image"`
}
type User struct {
	No    string `json:"no"`
	Text  string `json:"text"`
	Index string `json:"index"`
}
