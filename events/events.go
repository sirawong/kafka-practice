package events

type Event interface{}

type Messages struct {
	MsgId  int    `json:"Msg_id"`
	Sender string `json:"Sender"`
	Msg    string `json:"Msg"`
}