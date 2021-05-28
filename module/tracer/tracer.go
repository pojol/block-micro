// 接口文件 tracer 分布式追踪服务
package tracer

import (
	"github.com/pojol/braid-go/module"
)

// SpanFactory span 工厂
type SpanFactory func(interface{}) (ISpan, error)

// ISpan span interface
type ISpan interface {
	Begin(ctx interface{})
	SetTag(key string, val interface{})
	GetID() string
	End(ctx interface{})
}

// ITracer tracer interface
type ITracer interface {
	module.IModule

	GetSpan(strategy string) (ISpan, error)

	GetTracing() interface{}
}
