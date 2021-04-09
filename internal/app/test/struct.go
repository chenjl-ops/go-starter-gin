package test

type RobotTicket struct {
	Id               int    `xorm:"INT(11) 'id'"`
	Status           string `xorm:"INT(11) 'status'"`
	MessageId        string `xorm:"VARCHAR(128) 'message_id'"`
	ProcessingPerson string `xorm:"VARCHAR(32) 'processing_person'"`
	Desc             string `xorm:"VARCHAR(1024)" 'desc'`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
