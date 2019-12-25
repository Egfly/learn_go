package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalance("random", &RandomBalance{})
}

/**
随机调度
*/
type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (ins *Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("instance can not be empty")
		return
	}
	index := rand.Intn(lens)
	ins = insts[index]

	return
}
