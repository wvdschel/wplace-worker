package bot

import (
	"log"
	"time"

	"github.com/avast/retry-go"
)

var defaultRetryOpts = []retry.Option{
	retry.Attempts(5),
	retry.DelayType(retry.BackOffDelay),
	retry.Delay(100 * time.Millisecond),
}

func (b *Bot) log(accountIdx int, formatString string, ival ...any) {
	accountName := b.accounts[accountIdx].userInfo.Name

	args := []any{accountName}
	args = append(args, ival...)
	formatString = "%s: " + formatString
	log.Printf(formatString, args...)
}
