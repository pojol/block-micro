package balancer

/*
func TestParm(t *testing.T) {
	serviceName := "TestParm"

	blog.New(blog.NewWithDefault())
	mb := module.GetBuilder(pubsub.Name).Build(serviceName).(pubsub.IPubsub)

	bgb := module.GetBuilder(Name)
	b := bgb.Build(serviceName,
		moduleparm.WithPubsub(mb))
	bg := b.(balancer.IBalancer)

	bg.Init()
	bg.Run()
	defer bg.Close()

	mb.GetTopic(discover.ServiceUpdate).Pub(discover.EncodeUpdateMsg(
		discover.EventAddService,
		discover.Node{
			ID:      "A",
			Address: "A",
			Weight:  4,
			Name:    serviceName,
		},
	))
	mb.GetTopic(discover.EventAddService).Pub(discover.EncodeUpdateMsg(
		discover.EventAddService,
		discover.Node{
			ID:      "B",
			Address: "B",
			Weight:  2,
			Name:    serviceName,
		},
	))

	time.Sleep(time.Millisecond * 100)
	mb.GetTopic(discover.ServiceUpdate).Pub(discover.EncodeUpdateMsg(
		discover.EventUpdateService,
		discover.Node{
			ID:      "A",
			Address: "A",
			Weight:  3,
			Name:    serviceName,
		},
	))
	mb.GetTopic(discover.EventRemoveService).Pub(discover.EncodeUpdateMsg(
		discover.EventRemoveService,
		discover.Node{
			ID:      "B",
			Address: "B",
			Weight:  2,
			Name:    serviceName,
		},
	))

	time.Sleep(time.Millisecond * 500)
	for i := 0; i < 10; i++ {
		nod, err := bg.Pick(StrategyRandom, serviceName)
		if err != nil {
			t.FailNow()
		}
		assert.Equal(t, nod.ID, "A")
	}
}
*/
