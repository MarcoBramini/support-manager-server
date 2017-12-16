package datasource

import (
	"support-manager-server/model"
)

func CreateMessage(conversation model.Conversation, message model.Message) error {
	conversation.Messages = append(conversation.Messages, message)
	err := UpdateConversation(conversation)
	if err != nil {
		return err
	}
	return nil
}