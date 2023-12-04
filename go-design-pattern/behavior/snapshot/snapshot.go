package snapshot

import (
	"container/list"
	"strings"
)

type InputText struct {
	text strings.Builder
}

func (i *InputText) Append(input string) {
	i.text.WriteString(input)
	//fmt.Fprint(&i.text, input)
}

func (i *InputText) GetText() string {
	return i.text.String()
}

func (i *InputText) CreateSnapshot() Snapshot {
	return newSnapshot(i.GetText())
}

func (i *InputText) RestoreSnapshot(snapshot *Snapshot) {
	reset := strings.Builder{}
	reset.WriteString(snapshot.text)
	i.text = reset
}

type Snapshot struct {
	text string
}

func newSnapshot(text string) Snapshot {
	return Snapshot{text: text}
}

type SnapshotHolder struct {
	snapshots *Stack
}

func (s *SnapshotHolder) popSnapshot() *Snapshot {
	s1 := s.snapshots.Pop()
	l := list.New()
	l.PushBack(s1)
	// Iterate through list and print its contents.
	e := l.Front()
	snapshot, ok := (e.Value).(Snapshot)
	if ok {
		return &snapshot
	}
	return nil
}

func (s *SnapshotHolder) pushSnapshot(snapshot Snapshot) {
	s.snapshots.Push(snapshot)
}

func GetSnapshotHolder() *SnapshotHolder {
	return &SnapshotHolder{
		snapshots: NewStack(),
	}
}
