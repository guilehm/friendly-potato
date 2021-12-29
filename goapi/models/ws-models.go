package models

type WSMessageType string

var (
	Login  WSMessageType = "login"
	Update               = "update"
)

type WSMessage struct {
	messageType WSMessageType
	data        interface{}
}
