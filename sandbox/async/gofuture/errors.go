package gofuture

import "errors"

var (
	ErrFututeNotCompleted     = errors.New("Future not completed")
	ErrFutureAlreadyCompleted = errors.New("Future already completed")
	ErrFutureNotInitialized   = errors.New("Future not initialized")
)
