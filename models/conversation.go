package models

type Conversation struct {
	Base
	Sender   int64  `json:"sender"`
	Receiver int64  `json:"receiver"`
	Message  string `json:"message"`
	Status   string `json:"status"`
}

func (c Conversation) TableName() string {
	return "conversations"
}

func (c Conversation) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}
