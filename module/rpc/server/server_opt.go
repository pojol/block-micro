package server

import (
	"github.com/pojol/braid-go/depend/btracer"
	"google.golang.org/grpc"
)

// Parm Service 配置
type Parm struct {
	ListenAddr string

	Tracer btracer.ITracer

	Interceptors []grpc.UnaryServerInterceptor

	GracefulStop bool
}

// Option config wraps
type Option func(*Parm)

// WithListen 服务器侦听地址配置
func WithListen(address string) Option {
	return func(c *Parm) {
		c.ListenAddr = address
	}
}

func WithGracefulStop() Option {
	return func(c *Parm) {
		c.GracefulStop = true
	}
}

func AppendInterceptors(interceptor grpc.UnaryServerInterceptor) Option {
	return func(c *Parm) {
		c.Interceptors = append(c.Interceptors, interceptor)
	}
}

func WithTracer(t btracer.ITracer) Option {
	return func(c *Parm) {
		c.Tracer = t
	}
}
