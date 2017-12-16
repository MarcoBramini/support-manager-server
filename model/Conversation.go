package model

type Conversation struct {
	Id string `bson:"_id"`
	Messages []Message `bson:"messages"`
	Users UserList `bson:"users"`
}
