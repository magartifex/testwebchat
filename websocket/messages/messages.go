package messages

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// MessageStruct структура сообщения
type MessageStruct struct {
	UserID  string
	Message string
	Time    int64
}

type historyStruct struct {
	messages []*MessageStruct
	sync.Mutex
}

var history = &historyStruct{}

// SetHistory генерация сообщения и запись в историю
func SetHistory(userID string, text string) *MessageStruct {
	message := &MessageStruct{
		UserID:  userID,
		Message: text,
		Time:    time.Now().Unix(),
	}

	// Добавляем в историю
	history.Lock()
	history.messages = append(history.messages, message)
	// Можно поставить на строку выше и в defer, но defer немного замедляет работу, а здесь его можно избежать.
	history.Unlock()

	return message
}

// SendHistory Отправить историю новому пользователю
func SendHistory(conn *websocket.Conn) error {
	history.Lock()
	defer history.Unlock()

	if len(history.messages) != 0 {
		return conn.WriteJSON(history.messages)
	}

	return nil
}
