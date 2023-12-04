package prototype

import (
	"encoding/json"
	"time"
)

type KeyWord struct {
	word     string
	visit    int
	UpdateAt *time.Time
}

func (k *KeyWord) Clone() *KeyWord {
	var newKeyWord KeyWord
	b, _ := json.Marshal(k)
	json.Unmarshal(b, newKeyWord)
	return &newKeyWord

}

type KeyWords map[string]*KeyWord

func (words KeyWords) Clone(updateWords []*KeyWord) KeyWords {
	newKeyWords := KeyWords{}
	for s, word := range words {
		newKeyWords[s] = word
	}

	for _, word := range updateWords {
		newKeyWords[word.word] = word.Clone()
	}
	return newKeyWords
}
