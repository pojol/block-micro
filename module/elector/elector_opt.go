package elector

import (
	"time"
)

// Parm 选举器配置项
type Parm struct {
	ConsulAddr        string
	ServiceName       string
	LockTick          time.Duration
	RefushSessionTick time.Duration
}

// Option consul discover config wrapper
type Option func(*Parm)

// WithLockTick with lock tick
func WithLockTick(t time.Duration) Option {
	return func(c *Parm) {
		c.LockTick = t
	}
}

// WithSessionTick with session tick
func WithSessionTick(t time.Duration) Option {
	return func(c *Parm) {
		c.RefushSessionTick = t
	}
}
