package worker

import (
	"context"

	"github.com/crxfoz/outbox"
)

type Options struct {
	logger      outbox.Logger
	ctxProvider func(context.Context) context.Context
}

type Option func(*Options)

// WithLogger allows you to define custom logger
//
// Be default no logged is used
func WithLogger(logger outbox.Logger) Option {
	return func(options *Options) {
		options.logger = logger
	}
}

// WithCtxProvider will be used each time ctx is required for any call in Repository
//
// Allows you to define timeouts for requests to DB made by library
func WithCtxProvider(fn func(context.Context) context.Context) Option {
	return func(options *Options) {
		options.ctxProvider = fn
	}
}
