package models

type SmsMessage struct {
	Address string `json:"address"`
	Body    string `json:"body"`
	Date    int64  `json:"date"`
	Type    int    `json:"type"`
}
