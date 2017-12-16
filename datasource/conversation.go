package datasource

import (
	"support-manager-server/model"
	"errors"
)

func GetConversation(id string) (model.Conversation, error){
	var conversation model.Conversation
	err := db.C("conversations").FindId(id).One(&conversation)
	if err != nil {
		return conversation, errors.New("db: conversation not found")
	}
	return conversation, nil
}

func UpdateConversation(conversation model.Conversation) error{
	err := db.C("conversations").UpdateId(conversation.Id, conversation)
	if err!= nil {
		return errors.New("db: conversation update failed")
	}
	return nil
}