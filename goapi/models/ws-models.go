package models

import "encoding/json"

type WSMessageType string

var (
	Login  WSMessageType = "login"
	Update WSMessageType = "update"
)

type WSMessage struct {
	MessageType WSMessageType   `json:"type"`
	Data        json.RawMessage `json:"data"`
}
