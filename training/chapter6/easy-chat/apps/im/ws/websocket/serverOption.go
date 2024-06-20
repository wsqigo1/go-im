package websocket

import "time"

type ServerOptionFunc func(opt *serverOption)

type serverOption struct {
	Authentication

	ack        AckType
	ackTimeout time.Duration

	pattern string

	maxConnectionIdle time.Duration
}

func newServerOption(opts ...ServerOptionFunc) serverOption {
	o := serverOption{
		Authentication:    new(authentication),
		maxConnectionIdle: defaultMaxConnectionIdle,
		ackTimeout:        defaultAckTimeout,
		pattern:           "/ws",
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithServerAuthentication(auth Authentication) ServerOptionFunc {
	return func(opt *serverOption) {
		opt.Authentication = auth
	}
}

func WithServerPattern(pattern string) ServerOptionFunc {
	return func(opt *serverOption) {
		opt.pattern = pattern
	}
}

func WithServerAck(ack AckType) ServerOptionFunc {
	return func(opt *serverOption) {
		opt.ack = ack
	}
}

func WithServerMaxConnectionIdle(maxConnectionIdle time.Duration) ServerOptionFunc {
	return func(opt *serverOption) {
		if maxConnectionIdle > 0 {
			opt.maxConnectionIdle = maxConnectionIdle
		}
	}
}
