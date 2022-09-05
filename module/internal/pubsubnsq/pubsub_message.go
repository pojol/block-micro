package pubsubnsq

import (
	"sync"

	"github.com/pojol/braid-go/module/pubsub"
)

// Unbounded 一个支持任意长度的channel实现
type UnboundedMsg struct {
	c       chan *pubsub.Message
	backlog []*pubsub.Message

	sync.Mutex
}

// NewUnbounded 构建Unbounded
func NewUnbounded() *UnboundedMsg {
	return &UnboundedMsg{c: make(chan *pubsub.Message, 1)}
}

// Put 输入一个新的信息
func (b *UnboundedMsg) Put(msg *pubsub.Message) {
	b.Lock()
	if len(b.backlog) == 0 {
		select {
		case b.c <- msg:
			b.Unlock()
			return
		default:
		}
	}
	b.backlog = append(b.backlog, msg)
	b.Unlock()
}

// Load 将积压队列中的头部数据提取到channel，并将队列整体前移一位。
func (b *UnboundedMsg) Load() {
	b.Lock()
	if len(b.backlog) > 0 {
		select {
		case b.c <- b.backlog[0]:
			b.backlog[0] = nil
			b.backlog = b.backlog[1:]
		default:
		}
	}
	b.Unlock()
}

// Get 获取unbounded的读channel
func (b *UnboundedMsg) Get() <-chan *pubsub.Message {
	return b.c
}
