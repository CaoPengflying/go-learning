package lru

import (
	"errors"
	"time"
)

const _maxCount = 10

var (
	ErrNotFound = errors.New("not found")
)

type DataNode struct {
	Key     string
	Value   string
	TimeOut int64

	Pre  *DataNode
	Next *DataNode
}

type LRU struct {
	head    *DataNode
	count   int
	dataMap map[string]*DataNode
}

func (lru *LRU) Set(key, value string, timeout int64) {
	node := initNode(key, value, timeout)
	defer func() {
		lru.dataMap[key] = &node
	}()

	//缓存为空
	if lru.head.Next == nil {
		lru.head.Next = &node
		node.Pre = lru.head
		lru.count++
		return
	}

	//缓存不为空，判断是否包含了这个元素
	//包含这个元素则将这个元素清除并重新插入
	if existNode, ok := lru.dataMap[key]; ok {
		pre := existNode.Pre
		pre.Next = existNode.Next
		existNode.Next.Pre = pre

		existNode.Pre = nil
		existNode.Next = nil

		node.Next = lru.head.Next
		lru.head.Next.Pre = &node
		lru.head.Next = &node
		node.Pre = lru.head
		return
	}
	/*	cursor := lru.head
		for cursor.Next != nil {
			//包含这个元素则将这个元素清除并重新插入
			if cursor.Next.Key == key {
				cursor.Next = cursor.Next.Next
				node.Next = lru.head.Next
				lru.head.Next = &node
				return
			}
			cursor = cursor.Next
		}*/

	//不包含判断长度是否满了
	cursor := lru.head
	if lru.count >= _maxCount {
		for cursor.Next.Next != nil {
			cursor = cursor.Next
		}
		delete(lru.dataMap,cursor.Next.Key)
		cursor.Next.Pre = nil
		cursor.Next = nil

		node.Next = lru.head.Next
		lru.head.Next.Pre = &node
		lru.head.Next = &node
		node.Pre = lru.head
		return
	}

	//正常插入
	node.Next = lru.head.Next
	lru.head.Next.Pre = &node
	lru.head.Next = &node
	node.Pre = lru.head
	lru.count++
	return

}

func (lru *LRU) Get(key string) (string, error) {
	if lru.head == nil {
		return "", ErrNotFound
	}

	if existNode, ok := lru.dataMap[key]; ok {
		if existNode.TimeOut <= time.Now().Unix() && existNode.TimeOut > 0 {
			pre := existNode.Pre
			pre.Next = existNode.Next
			if existNode.Next != nil {
				existNode.Next.Pre = pre
			}

			existNode.Next = nil
			existNode.Pre = nil

			delete(lru.dataMap, existNode.Key)

			lru.count--
			return "", ErrNotFound
		} else {
			return existNode.Value, nil
		}
	}
	/*cursor := lru.head
	for cursor.Next != nil {
		//找到时需要判断是否过期
		if cursor.Next.Key == key {
			if cursor.Next.TimeOut <= time.Now().Unix() && cursor.Next.TimeOut > 0 {
				cursor.Next = cursor.Next.Next
				lru.count--
				return "", ErrNotFound
			}

			return cursor.Next.Value, nil
		}
		cursor = cursor.Next
	}*/

	return "", ErrNotFound
}

func initNode(key string, value string, timeout int64) DataNode {
	node := DataNode{
		Key:   key,
		Value: value,
	}

	if timeout < 0 {
		node.TimeOut = timeout
	} else {
		node.TimeOut = time.Now().Unix() + timeout
	}

	return node
}

func NewLru() LRU {
	return LRU{head: &DataNode{}, count: 0, dataMap: make(map[string]*DataNode, _maxCount)}
}
