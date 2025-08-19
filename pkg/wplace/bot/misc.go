package bot

import (
	"time"

	"github.com/avast/retry-go"
)

var defaultRetryOpts = []retry.Option{
	retry.Attempts(5),
	retry.DelayType(retry.BackOffDelay),
	retry.Delay(100 * time.Millisecond),
}
