package balance

import (
	"errors"
)

func init() {
	RegisterBalance("roundRobinBalance", &RoundRobinBalance{})
}

/**
轮询调度
*/
type RoundRobinBalance struct {
	currentIndex int
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (ins *Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("instance can not be empty")
		return
	}

	ins = insts[p.currentIndex]
	p.currentIndex = (p.currentIndex + 1) % lens
	return
}
