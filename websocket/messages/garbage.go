package messages

import (
	"time"
)

// Сколько времени будет храниться история. Я бы вынес в конфиг
const timeOut int64 = 60 * 5

func garbageСollector() {
	history.Lock()
	defer history.Unlock()

	// Всё что раньше этого времени, должно быть удалено
	t := time.Now().Unix() - timeOut

	for i := len(history.messages); i > 0; i-- {
		if history.messages[i-1].Time > t {
			continue
		}

		history.messages = history.messages[i:]
		return
	}
}

// GarbageСollectorWorker worker
func GarbageСollectorWorker() {
	for {
		// C какой частотой мы запускаем чистку истории
		time.Sleep(5 * time.Second)
		garbageСollector()
	}
}
