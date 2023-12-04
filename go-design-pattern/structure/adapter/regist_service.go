package adapter

import "log"

type IRegister interface {
	RegisterService(serviceName string, addr string)
}

type ZkRegister struct {
}

func (zk ZkRegister) ZkRegister(serviceName string, addr string) {
	log.Print("zookeeper register service", serviceName, addr)
}

type ConsulRegister struct {
}

func (zk ConsulRegister) ConsulRegister(serviceName string, addr string) {
	log.Print("zookeeper register service", serviceName, addr)
}

type ZkRegisterAdapter struct {
	zkRegister ZkRegister
}

func (z ZkRegisterAdapter) RegisterService(serviceName string, addr string) {
	z.zkRegister.ZkRegister(serviceName, addr)
}

type ConsulRegisterAdapter struct {
	consulRegister ConsulRegister
}

func (c ConsulRegisterAdapter) RegisterService(serviceName string, addr string) {
	c.consulRegister.ConsulRegister(serviceName, addr)
}
