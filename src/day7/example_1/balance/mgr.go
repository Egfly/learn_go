package balance

import "fmt"

type BalanceMgr struct {
	allBalance map[string]Balance
}

var (
	mgr = BalanceMgr{
		allBalance: make(map[string]Balance),
	}
)

/**
自定义注册负载均衡方法
*/
func (p *BalanceMgr) registerBalance(name string, b Balance) {
	p.allBalance[name] = b
}

func RegisterBalance(name string, b Balance) {
	mgr.allBalance[name] = b
}

func DoBalance(name string, intsts []*Instance) (inst *Instance, err error) {
	balance, ok := mgr.allBalance[name]
	if !ok {
		err = fmt.Errorf("not found %s balance", name)
	}
	fmt.Printf("use %s balance\n", name)
	return balance.DoBalance(intsts)
}
