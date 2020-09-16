package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage(t *testing.T) {
	m := cm.CreateConcurrentMap(100)
	m.Set(cm.StrKey("cpf"),"zzc")
	t.Log(m.Get(cm.StrKey("cpf")))
}
