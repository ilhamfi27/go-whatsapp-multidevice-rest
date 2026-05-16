package whatsapp

import (
	"time"

	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

// WebhookPayload is the unified envelope sent to the webhook URL for all events.
type WebhookPayload struct {
	Event     string      `json:"event"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// --- Message Event ---

type WebhookMessageData struct {
	ID         string    `json:"id"`
	From       string    `json:"from"`
	FromName   string    `json:"from_name"`
	Chat       string    `json:"chat"`
	Timestamp  time.Time `json:"timestamp"`
	IsGroup    bool      `json:"is_group"`
	IsFromMe   bool      `json:"is_from_me"`
	IsEdit     bool      `json:"is_edit,omitempty"`
	IsViewOnce bool      `json:"is_view_once,omitempty"`
	Type       string    `json:"type"`
	Text       string    `json:"text,omitempty"`
}

func buildMessagePayload(v *events.Message) WebhookPayload {
	msgType := "unknown"
	msgText := ""

	msg := v.Message
	if msg != nil {
		switch {
		case msg.GetConversation() != "":
			msgType = "text"
			msgText = msg.GetConversation()
		case msg.GetExtendedTextMessage() != nil:
			msgType = "text"
			msgText = msg.GetExtendedTextMessage().GetText()
		case msg.GetImageMessage() != nil:
			msgType = "image"
			msgText = msg.GetImageMessage().GetCaption()
		case msg.GetVideoMessage() != nil:
			msgType = "video"
			msgText = msg.GetVideoMessage().GetCaption()
		case msg.GetAudioMessage() != nil:
			msgType = "audio"
		case msg.GetDocumentMessage() != nil:
			msgType = "document"
			msgText = msg.GetDocumentMessage().GetCaption()
		case msg.GetStickerMessage() != nil:
			msgType = "sticker"
		case msg.GetLocationMessage() != nil:
			msgType = "location"
		case msg.GetContactMessage() != nil:
			msgType = "contact"
		case msg.GetReactionMessage() != nil:
			msgType = "reaction"
			msgText = msg.GetReactionMessage().GetText()
		case msg.GetPollCreationMessage() != nil:
			msgType = "poll"
			msgText = msg.GetPollCreationMessage().GetName()
		}
	}

	return WebhookPayload{
		Event:     "message",
		Timestamp: time.Now(),
		Data: WebhookMessageData{
			ID:         v.Info.ID,
			From:       v.Info.Sender.String(),
			FromName:   v.Info.PushName,
			Chat:       v.Info.Chat.String(),
			Timestamp:  v.Info.Timestamp,
			IsGroup:    v.Info.IsGroup,
			IsFromMe:   v.Info.IsFromMe,
			IsEdit:     v.IsEdit,
			IsViewOnce: v.IsViewOnce,
			Type:       msgType,
			Text:       msgText,
		},
	}
}

// --- Receipt Event ---

type WebhookReceiptData struct {
	MessageIDs    []string          `json:"message_ids"`
	From          string            `json:"from"`
	Chat          string            `json:"chat"`
	IsGroup       bool              `json:"is_group"`
	Timestamp     time.Time         `json:"timestamp"`
	Type          types.ReceiptType `json:"type"`
	MessageSender string            `json:"message_sender,omitempty"`
}

func buildReceiptPayload(v *events.Receipt) WebhookPayload {
	msgSender := ""
	if !v.MessageSender.IsEmpty() {
		msgSender = v.MessageSender.String()
	}
	ids := make([]string, len(v.MessageIDs))
	copy(ids, v.MessageIDs)
	return WebhookPayload{
		Event:     "receipt",
		Timestamp: time.Now(),
		Data: WebhookReceiptData{
			MessageIDs:    ids,
			From:          v.Sender.String(),
			Chat:          v.Chat.String(),
			IsGroup:       v.IsGroup,
			Timestamp:     v.Timestamp,
			Type:          v.Type,
			MessageSender: msgSender,
		},
	}
}

// --- Chat Presence Event ---

type WebhookChatPresenceData struct {
	From  string `json:"from"`
	Chat  string `json:"chat"`
	State string `json:"state"`
	Media string `json:"media"`
}

func buildChatPresencePayload(v *events.ChatPresence) WebhookPayload {
	return WebhookPayload{
		Event:     "chat_presence",
		Timestamp: time.Now(),
		Data: WebhookChatPresenceData{
			From:  v.Sender.String(),
			Chat:  v.Chat.String(),
			State: string(v.State),
			Media: string(v.Media),
		},
	}
}

// --- Presence Event ---

type WebhookPresenceData struct {
	From        string    `json:"from"`
	Unavailable bool      `json:"unavailable"`
	LastSeen    time.Time `json:"last_seen,omitempty"`
}

func buildPresencePayload(v *events.Presence) WebhookPayload {
	return WebhookPayload{
		Event:     "presence",
		Timestamp: time.Now(),
		Data: WebhookPresenceData{
			From:        v.From.String(),
			Unavailable: v.Unavailable,
			LastSeen:    v.LastSeen,
		},
	}
}

// --- Group Info Event ---

type WebhookGroupInfoData struct {
	JID       string    `json:"jid"`
	Sender    string    `json:"sender,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name,omitempty"`
	Topic     string    `json:"topic,omitempty"`
	Join      []string  `json:"join,omitempty"`
	Leave     []string  `json:"leave,omitempty"`
	Promote   []string  `json:"promote,omitempty"`
	Demote    []string  `json:"demote,omitempty"`
}

func jidSliceToStrings(jids []types.JID) []string {
	if len(jids) == 0 {
		return nil
	}
	result := make([]string, len(jids))
	for i, j := range jids {
		result[i] = j.String()
	}
	return result
}

func buildGroupInfoPayload(v *events.GroupInfo) WebhookPayload {
	sender := ""
	if v.Sender != nil {
		sender = v.Sender.String()
	}
	name := ""
	if v.Name != nil {
		name = v.Name.Name
	}
	topic := ""
	if v.Topic != nil {
		topic = v.Topic.Topic
	}
	return WebhookPayload{
		Event:     "group_info",
		Timestamp: time.Now(),
		Data: WebhookGroupInfoData{
			JID:       v.JID.String(),
			Sender:    sender,
			Timestamp: v.Timestamp,
			Name:      name,
			Topic:     topic,
			Join:      jidSliceToStrings(v.Join),
			Leave:     jidSliceToStrings(v.Leave),
			Promote:   jidSliceToStrings(v.Promote),
			Demote:    jidSliceToStrings(v.Demote),
		},
	}
}

// --- Joined Group Event ---

type WebhookJoinedGroupData struct {
	JID    string `json:"jid"`
	Name   string `json:"name"`
	Reason string `json:"reason,omitempty"`
	Type   string `json:"type,omitempty"`
}

func buildJoinedGroupPayload(v *events.JoinedGroup) WebhookPayload {
	return WebhookPayload{
		Event:     "joined_group",
		Timestamp: time.Now(),
		Data: WebhookJoinedGroupData{
			JID:    v.JID.String(),
			Name:   v.GroupInfo.Name,
			Reason: v.Reason,
			Type:   v.Type,
		},
	}
}

// --- Picture Event ---

type WebhookPictureData struct {
	JID       string    `json:"jid"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
	Removed   bool      `json:"removed"`
	PictureID string    `json:"picture_id,omitempty"`
}

func buildPicturePayload(v *events.Picture) WebhookPayload {
	return WebhookPayload{
		Event:     "picture",
		Timestamp: time.Now(),
		Data: WebhookPictureData{
			JID:       v.JID.String(),
			Author:    v.Author.String(),
			Timestamp: v.Timestamp,
			Removed:   v.Remove,
			PictureID: v.PictureID,
		},
	}
}

// --- Connection Events ---

type WebhookConnectionData struct {
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

func buildConnectedPayload() WebhookPayload {
	return WebhookPayload{
		Event:     "connection",
		Timestamp: time.Now(),
		Data:      WebhookConnectionData{Status: "connected"},
	}
}

func buildDisconnectedPayload() WebhookPayload {
	return WebhookPayload{
		Event:     "connection",
		Timestamp: time.Now(),
		Data:      WebhookConnectionData{Status: "disconnected"},
	}
}

func buildLoggedOutPayload(v *events.LoggedOut) WebhookPayload {
	return WebhookPayload{
		Event:     "connection",
		Timestamp: time.Now(),
		Data:      WebhookConnectionData{Status: "logged_out", Reason: v.Reason.String()},
	}
}

// --- Call Events ---

type WebhookCallData struct {
	CallID    string    `json:"call_id"`
	From      string    `json:"from"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
	Reason    string    `json:"reason,omitempty"`
}

func buildCallOfferPayload(v *events.CallOffer) WebhookPayload {
	return WebhookPayload{
		Event:     "call",
		Timestamp: time.Now(),
		Data: WebhookCallData{
			CallID:    v.CallID,
			From:      v.From.String(),
			Timestamp: v.Timestamp,
			Status:    "offer",
		},
	}
}

func buildCallTerminatePayload(v *events.CallTerminate) WebhookPayload {
	return WebhookPayload{
		Event:     "call",
		Timestamp: time.Now(),
		Data: WebhookCallData{
			CallID:    v.CallID,
			From:      v.From.String(),
			Timestamp: v.Timestamp,
			Status:    "terminate",
			Reason:    v.Reason,
		},
	}
}

func buildCallRejectPayload(v *events.CallReject) WebhookPayload {
	return WebhookPayload{
		Event:     "call",
		Timestamp: time.Now(),
		Data: WebhookCallData{
			CallID:    v.CallID,
			From:      v.From.String(),
			Timestamp: v.Timestamp,
			Status:    "reject",
		},
	}
}
