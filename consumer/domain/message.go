package domain

type Message struct {
	MsgId  int    `bson:"Msg_id"`
	Sender string `bson:"Sender"`
	Msg    string `bson:"Msg"`
}