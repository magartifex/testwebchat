package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"

	"testwebchat/libs"
	// не очень люблю делать вложенные пакеты, но всё же.
	"testwebchat/websocket/messages"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Можно использовать mutex и обычный map, но sync.Map удобней и в некоторых случаях быстрее
	users = sync.Map{}
)

// Handler Подключаем нового пользователя к вебсокету
func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket.Handler #1:\nError: %s\n\n", err)
		return
	}

	// Сразу после подключения пользователя, отправляем ему всю историю
	err = messages.SendHistory(conn)
	if err != nil {
		log.Printf("websocket.Handler #2:\nError: %s\n\n", err)
		return
	}

	userID := generateUserID(conn)
	defer users.Delete(userID)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			// Выводить логи кроме случая когда пользователь отключился.
			if messageType != -1 {
				log.Printf("websocket.Handler #3:\nmessageType: %d\nUserID: %s\nError: %s\n\n", messageType, userID, err)
			}
			return
		}

		broadcast([]*messages.MessageStruct{
			messages.SetHistory(userID, string(p)),
		})
	}
}

// broadcast Рассылка сообщения по всем пользователям
func broadcast(message []*messages.MessageStruct) {
	users.Range(
		func(_, value interface{}) bool {
			conn, ok := value.(*websocket.Conn)
			if !ok {
				return true
			}

			err := conn.WriteJSON(message)
			if err != nil {
				log.Printf("websocket.Handler #1:\nMessage: %v\nError: %s\n\n", message, err)
				return true
			}

			return true
		},
	)
}

// Генерируем userID с решением проблемы с коллизиями
func generateUserID(conn *websocket.Conn) string {
	userID := libs.RandStringBytes(16)
	// Если такого пользователя в системе нет, то записываем его в список пользователей.
	// Иначе заново запускаем генерацию
	_, loaded := users.LoadOrStore(userID, conn)
	if loaded {
		return generateUserID(conn)
	}

	return userID
}
