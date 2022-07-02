package mixinMessenger

import (
	"context"
	"encoding/json"
	"encoding/base64"
	"github.com/gofrs/uuid"
        "github.com/fox-one/mixin-sdk-go"
)

func SendMixinMsg(client *mixin.Client, userID, conversationID string, data []byte) error{
        payload := base64.StdEncoding.EncodeToString(data)
        return client.SendMessage(context.Background(), &mixin.MessageRequest{
                ConversationID: conversationID,
                RecipientID:    userID,
                MessageID:      uuid.Must(uuid.NewV4()).String(),
                Category:       mixin.MessageCategoryPlainText,
                Data:           payload,
        })
}

func SendMixinBtn(client *mixin.Client, userID, conversationID, label, action, color string) error{
	button := &mixin.AppButtonMessage{
		Label: label,
		Action: action,
		Color: color,
	}
	btngrp := &mixin.AppButtonGroupMessage{
		*button,
	}
	b, err := json.Marshal(btngrp)
	if err != nil {
		return err
	}
	return client.SendMessage(context.Background(), &mixin.MessageRequest{
		RecipientID:    userID,
		ConversationID: conversationID,
                MessageID:      uuid.Must(uuid.NewV4()).String(),
                Category:       mixin.MessageCategoryAppButtonGroup,
                Data:		base64.StdEncoding.EncodeToString(b),
        })
}
