package model

type Message struct {
	Id string `bson:"_id"`
	Owner User `bson:"owner"`
	Payload string `bson:"payload"`
	Time string `bson:"time"`
	ConversationId string `bson:"conversationId"`
}