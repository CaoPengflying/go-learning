package snapshot

import "testing"

func TestInputText_CreateSnapshot(t *testing.T) {
	it := InputText{}
	sh := GetSnapshotHolder()
	sh.pushSnapshot(it.CreateSnapshot())
	it.Append("hello")
	sh.pushSnapshot(it.CreateSnapshot())
	it.Append("world")
	t.Log(it.GetText())
	restoreSnapshot1 := sh.popSnapshot()
	it.RestoreSnapshot(restoreSnapshot1)
	t.Log(it.GetText())
}
