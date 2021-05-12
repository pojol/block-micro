package mailboxnsq

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/pojol/braid-go/mock"
	"github.com/pojol/braid-go/module/logger"
	"github.com/pojol/braid-go/module/mailbox"
	"github.com/pojol/braid-go/modules/zaplogger"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	mock.Init()

	m.Run()
}

func TestClusterBroadcast(t *testing.T) {

	b := mailbox.GetBuilder(Name)
	log, _ := logger.GetBuilder(zaplogger.Name).Build()
	b.AddOption(WithLookupAddr([]string{mock.NSQLookupdAddr}))
	b.AddOption(WithNsqdAddr([]string{mock.NsqdAddr}))
	b.AddOption(WithNsqdHTTPAddr([]string{mock.NsqdHttpAddr}))
	b.AddOption(WithNsqLogLv(nsq.LogLevelDebug))
	mb, _ := b.Build("TestClusterBroadcast", log)

	topic := "test.clusterBroadcast"

	mb.RegistTopic(topic, mailbox.ScopeCluster)

	channel1 := mb.GetTopic(topic).Sub("Normal_1")
	channel2 := mb.GetTopic(topic).Sub("Normal_2")

	var wg sync.WaitGroup
	done := make(chan struct{})
	wg.Add(2)

	go func() {
		for {
			select {
			case <-channel1.Arrived():
				wg.Done()
			case <-channel2.Arrived():
				wg.Done()
			}
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	mb.GetTopic(topic).Pub(&mailbox.Message{Body: []byte("test msg")})

	select {
	case <-done:
		// pass
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")

		res, _ := http.Get("http://127.0.0.1:4151/info")
		byt, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(byt))
		res.Body.Close()

		t.FailNow()
	}

}

func TestClusterNotify(t *testing.T) {

	b := mailbox.GetBuilder(Name)
	log, _ := logger.GetBuilder(zaplogger.Name).Build()
	b.AddOption(WithLookupAddr([]string{mock.NSQLookupdAddr}))
	b.AddOption(WithNsqdAddr([]string{mock.NsqdAddr}))
	b.AddOption(WithNsqdHTTPAddr([]string{mock.NsqdHttpAddr}))
	b.AddOption(WithNsqLogLv(nsq.LogLevelDebug))
	mb, _ := b.Build("TestClusterNotify", log)

	var tick uint64

	topic := "test.clusterNotify"

	mb.RegistTopic(topic, mailbox.ScopeCluster)

	channel1 := mb.GetTopic(topic).Sub("Normal")
	channel2 := mb.GetTopic(topic).Sub("Normal")

	go func() {
		for {
			select {
			case <-channel1.Arrived():
				atomic.AddUint64(&tick, 1)
			case <-channel2.Arrived():
				atomic.AddUint64(&tick, 1)
			}
		}
	}()

	mb.GetTopic(topic).Pub(&mailbox.Message{Body: []byte("msg")})

	select {
	case <-time.After(time.Second * 5):
		assert.Equal(t, atomic.LoadUint64(&tick), uint64(1))
	}

}

func BenchmarkClusterBoardcast(b *testing.B) {
	log, _ := logger.GetBuilder(zaplogger.Name).Build()

	mbb := mailbox.GetBuilder(Name)
	mbb.AddOption(WithLookupAddr([]string{mock.NSQLookupdAddr}))
	mbb.AddOption(WithNsqdAddr([]string{mock.NsqdAddr}))
	mbb.AddOption(WithNsqdHTTPAddr([]string{mock.NsqdHttpAddr}))

	mb, _ := mbb.Build("BenchmarkClusterBoardcast", log)
	topic := "benchmark.ClusterBroadcast"
	body := []byte("msg")

	mb.RegistTopic(topic, mailbox.ScopeCluster)

	c1 := mb.GetTopic(topic).Sub("Normal_1")
	c2 := mb.GetTopic(topic).Sub("Normal_2")

	go func() {
		for {
			select {
			case <-c1.Arrived():
			case <-c2.Arrived():
			}
		}
	}()

	b.SetParallelism(8)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mb.GetTopic(topic).Pub(&mailbox.Message{Body: body})
		}
	})

}

/*
//BenchmarkCompetition-8   	 3238792	       335 ns/op	      79 B/op	       2 allocs/op
func BenchmarkProcCompetition(b *testing.B) {
	mbb := mailbox.GetBuilder(Name)
	log, _ := logger.GetBuilder(zaplogger.Name).Build()
	mb, _ := mbb.Build("BenchmarkProcCompetition", log)
	topic := "BenchmarkProcCompetition"
	body := []byte("msg")

	sub := mb.Sub(mailbox.Proc, topic)
	c1, _ := sub.Competition()
	c2, _ := sub.Competition()

	c1.OnArrived(func(msg mailbox.Message) error {
		return nil
	})

	c2.OnArrived(func(msg mailbox.Message) error {
		return nil
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mb.Pub(mailbox.Proc, topic, &mailbox.Message{Body: body})
	}
}

func BenchmarkProcCompetitionAsync(b *testing.B) {
	mbb := mailbox.GetBuilder(Name)
	log, _ := logger.GetBuilder(zaplogger.Name).Build()
	mb, _ := mbb.Build("BenchmarkProcCompetitionAsync", log)
	topic := "BenchmarkProcCompetitionAsync"
	body := []byte("msg")

	sub := mb.Sub(mailbox.Proc, topic)
	c1, _ := sub.Competition()
	c2, _ := sub.Competition()

	c1.OnArrived(func(msg mailbox.Message) error {
		return nil
	})

	c2.OnArrived(func(msg mailbox.Message) error {
		return nil
	})

	b.SetParallelism(8)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mb.PubAsync(mailbox.Proc, topic, &mailbox.Message{Body: body})
		}
	})
}




func TestClusterMailboxParm(t *testing.T) {
	b := mailbox.GetBuilder(Name)
	log, _ := logger.GetBuilder(zaplogger.Name).Build()
	b.AddOption(WithChannel("parm"))
	b.AddOption(WithLookupAddr([]string{mock.NSQLookupdAddr}))
	b.AddOption(WithNsqdAddr([]string{mock.NsqdAddr}))

	mb, err := b.Build("cluster", log)
	assert.Equal(t, err, nil)
	cm := mb.(*nsqMailbox)

	assert.Equal(t, cm.parm.Channel, "parm")
	assert.Equal(t, cm.parm.LookupAddress, []string{mock.NSQLookupdAddr})
	assert.Equal(t, cm.parm.Address, []string{mock.NsqdAddr})
}
*/
