package handlers

import (
	"context"
	"time"
)

const contextTimeoutSecond = 2

func getContext() (context.Context, context.CancelFunc) {
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, contextTimeoutSecond*time.Second)

	return ctxWithTimeout, cancel
}
