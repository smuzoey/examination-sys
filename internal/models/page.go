package models

type Page struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`   // 总条数
	Size    int         `json:"size"`    // 当前 size
	Current int         `json:"current"` // 当前 num
	//Pages   int         `json:"pages"`   //
}
