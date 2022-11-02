package main

type Response struct {
	Ok     bool
	Result []Update
}

type Update struct {
	UpdateId      float64 `json:"update_id"`
	Message       *Message
	EditedMessage *Message `json:"edited_message"`
}

type Message struct {
	MessageId float64 `json:"message_id"`
	Date      float64
	Text      string
	From      any
	Chat      any
}
