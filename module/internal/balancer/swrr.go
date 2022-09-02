// 实现文件 balancerswrr 平滑加权负载均衡算法实现
package balancer

import (
	"errors"
	"sync"

	"github.com/pojol/braid-go/depend/blog"
	"github.com/pojol/braid-go/service"
)

type weightedNod struct {
	orgNod    service.Node
	curWeight int
}

// swrrBalancer 平滑加权轮询
type swrrBalancer struct {
	totalWeight int
	nods        []weightedNod
	sync.Mutex
}

func (wr *swrrBalancer) calcTotalWeight() {
	wr.totalWeight = 0

	for _, v := range wr.nods {
		wr.totalWeight += v.orgNod.Weight
	}
}

func (wr *swrrBalancer) isExist(id string) (int, bool) {
	for k, v := range wr.nods {
		if v.orgNod.ID == id {
			return k, true
		}
	}

	return -1, false
}

// Pick 执行算法，选取节点
func (wr *swrrBalancer) Get() (service.Node, error) {
	var tmpWeight int
	var idx int
	wr.Lock()
	defer wr.Unlock()

	if len(wr.nods) <= 0 {
		return service.Node{}, errors.New("empty")
	}

	for k, v := range wr.nods {
		if tmpWeight < v.curWeight+wr.totalWeight {
			tmpWeight = v.curWeight + wr.totalWeight
			idx = k
		}
	}

	for k := range wr.nods {
		if k == idx {
			wr.nods[idx].curWeight = wr.nods[idx].curWeight - wr.totalWeight + wr.nods[idx].orgNod.Weight
		} else {
			wr.nods[k].curWeight += wr.nods[k].orgNod.Weight
		}
	}

	return wr.nods[idx].orgNod, nil
}

func (wr *swrrBalancer) Add(nod service.Node) {

	if _, ok := wr.isExist(nod.ID); ok {
		return
	}

	wr.nods = append(wr.nods, weightedNod{
		orgNod:    nod,
		curWeight: int(nod.Weight),
	})

	wr.calcTotalWeight()

	blog.Debugf("add weighted nod id : %s name : %s weight : %d", nod.ID, nod.Name, nod.Weight)
}

func (wr *swrrBalancer) Rmv(nod service.Node) {

	var ok bool
	var idx int

	idx, ok = wr.isExist(nod.ID)
	if !ok {
		// log
		return
	}

	wr.nods = append(wr.nods[:idx], wr.nods[idx+1:]...)

	wr.calcTotalWeight()
	blog.Debugf("rmv weighted nod id : %s name : %s", nod.ID, nod.Name)
}

func (wr *swrrBalancer) Update(nod service.Node) {

	var ok bool
	var idx int

	idx, ok = wr.isExist(nod.ID)
	if ok {
		wr.nods[idx].orgNod.Weight = nod.Weight
		wr.calcTotalWeight()
	}

	blog.Debugf("update weighted nod id : %s name : %s weight : %d", nod.ID, nod.Name, nod.Weight)
}
