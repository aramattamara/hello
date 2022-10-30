package main

type Response struct {
	Ok     bool
	Result []Update
}

type Update struct {
	UpdateId int64 `json:"update_id"`
	Message  Message
}

type Message struct {
	MessageId int32
}
