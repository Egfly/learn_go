package balance

/**
负载均衡接口
*/
type Balance interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}
