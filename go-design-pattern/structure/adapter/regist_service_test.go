package adapter

import "testing"

func Test_RegisterService(t *testing.T) {
	var a IRegister = &ZkRegisterAdapter{
		zkRegister: ZkRegister{},
	}
	a.RegisterService(" zookeeper regist service","127.0.0.1")
}

func TestZkRegisterAdapter_RegisterService(t *testing.T) {
	var a IRegister = &ConsulRegisterAdapter{
		consulRegister: ConsulRegister{},
	}
	a.RegisterService(" consul regist service","127.0.0.1")
}


