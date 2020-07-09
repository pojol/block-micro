package redislinker

import (
	"testing"
	"time"

	"github.com/pojol/braid/3rd/redis"
	"github.com/pojol/braid/mock"
	"github.com/pojol/braid/module/linker"
	"github.com/stretchr/testify/assert"
)

func TestTarget(t *testing.T) {

	mock.Init()

	r := redis.New()
	r.Init(redis.Config{
		Address:        mock.RedisAddr,
		ReadTimeOut:    time.Millisecond * time.Duration(5000),
		WriteTimeOut:   time.Millisecond * time.Duration(5000),
		ConnectTimeOut: time.Millisecond * time.Duration(2000),
		IdleTimeout:    time.Millisecond * time.Duration(0),
		MaxIdle:        16,
		MaxActive:      128,
	})

	b := linker.GetBuilder(LinkerName)
	lk := b.Build(nil)

	num, err := lk.Num("testnodid")
	assert.Equal(t, num, 0)
	assert.Equal(t, err, nil)

	err = lk.Link("testtoken1", "testnodid", "192.168.0.1:8000")
	assert.Equal(t, err, nil)

	target, err := lk.Target("testtoken1")
	assert.Equal(t, target, "192.168.0.1:8000")
	assert.Equal(t, err, nil)

	lk.Offline("testnodid")
	num, err = lk.Num("testnodid")
	assert.Equal(t, num, 0)
	assert.Equal(t, err, nil)
}
