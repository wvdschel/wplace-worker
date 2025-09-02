package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go"
)

var defaultRetryOpts = []retry.Option{
	retry.Attempts(5),
	retry.DelayType(retry.BackOffDelay),
	retry.Delay(100 * time.Millisecond),
}

type LogEntry struct {
	Time    string `json:"time"`
	Message string `json:"message"`
}

func (b *Bot) log(accountIdx int, formatString string, ival ...any) {
	accountName := b.accounts[accountIdx].userInfo.Name

	args := []any{accountName}
	args = append(args, ival...)
	formatString = "%s: " + formatString
	message := fmt.Sprintf(formatString, args...)

	// Store in log buffer
	b.lock.Lock()
	b.logBuffer = append(b.logBuffer, LogEntry{
		Time:    time.Now().Format("15:04:05"),
		Message: message,
	})
	// Keep only last 100 entries
	if len(b.logBuffer) > 100 {
		b.logBuffer = b.logBuffer[len(b.logBuffer)-100:]
	}
	b.lock.Unlock()

	log.Printf(formatString, args...)
}
