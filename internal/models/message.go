package models

type Message struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

//type Replay struct {
//	ReplayId int    `json:"replayId" gorm:"column:replayId"`
//	Replay   string `json:"replay"`
//	//ReplayTime string `json:"replayTime"`
//	MessageId int `json:"messageId" gorm:"column:messageId"`
//}
