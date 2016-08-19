// +build !go1.7

package context

import (
	"golang.org/x/net/context"
	"time"
)

var Canceled = context.Canceled

var DeadlineExceeded = context.DeadlineExceeded

func Background() Context {
	return context.Background()
}

func TODO() Context {
	return context.TODO()
}

type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	ctx, f := context.WithCancel(parent)
	return ctx, CancelFunc(f)
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	ctx, f := context.WithDeadline(parent, deadline)
	return ctx, CancelFunc(f)
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	ctx, f := context.WithTimeout(parent, timeout)
	return ctx, CancelFunc(f)
}

func WithValue(parent Context, key, val interface{}) Context {
	return context.WithValue(parent, key, val)
}
