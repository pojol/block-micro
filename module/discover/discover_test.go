package discover

import (
	"testing"

	"github.com/pojol/braid-go/mock"
)

func TestMain(m *testing.M) {
	mock.Init()

	m.Run()
}

/*
func TestDiscover(t *testing.T) {

	b := module.GetBuilder(Name)

	blog.New(blog.NewWithDefault())
	mb := module.GetBuilder(pubsub.Name).Build("TestDiscover").(pubsub.IPubsub)

	b.AddModuleOption(WithConsulAddr(mock.ConsulAddr))
	b.AddModuleOption(WithSyncServiceInterval(time.Millisecond * 100))
	b.AddModuleOption(WithSyncServiceWeightInterval(time.Millisecond * 100))
	b.AddModuleOption(WithBlacklist([]string{"gate"}))

	dc := b.Build("test",
		moduleparm.WithPubsub(mb)).(*consulDiscover)
	assert.Equal(t, dc.InBlacklist("gate"), true)
	assert.Equal(t, dc.InBlacklist("login"), false)

	dc.Init()
	dc.Run()

	time.Sleep(time.Second)
	dc.Close()
}
*/

/*
func TestParm(t *testing.T) {
	b := module.GetBuilder(Name)

	blog.New(blog.NewWithDefault())
	mb := module.GetBuilder(pubsub.Name).Build("TestParm").(pubsub.IPubsub)

	b.AddModuleOption(WithConsulAddr(mock.ConsulAddr))
	b.AddModuleOption(WithTag("TestParm"))
	b.AddModuleOption(WithBlacklist([]string{"gate"}))
	b.AddModuleOption(WithSyncServiceInterval(time.Second))
	b.AddModuleOption(WithSyncServiceWeightInterval(time.Second))

	discv := b.Build("test",
		moduleparm.WithPubsub(mb)).(*consulDiscover)

	assert.Equal(t, discv.parm.Address, mock.ConsulAddr)
	assert.Equal(t, discv.parm.Tag, "TestParm")
	assert.Equal(t, discv.parm.Blacklist, []string{"gate"})
	assert.Equal(t, discv.parm.SyncServicesInterval, time.Second)
	assert.Equal(t, discv.parm.SyncServiceWeightInterval, time.Second)
}
*/
