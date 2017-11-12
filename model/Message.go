package model

type Message struct {
	Id string
	Owner User
	Payload string
	Time string
	ConversationId string
}