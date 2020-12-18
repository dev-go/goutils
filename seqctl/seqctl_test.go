// Copyright (c) 2020, devgo.club
// All rights reserved.

package seqctl

import (
	"testing"
)

type SeqMap struct {
	m  map[string]int
	sc SeqCtl
}

func NewSeqMap() *SeqMap {
	return &SeqMap{
		m:  make(map[string]int, 1<<10),
		sc: New("seq map", 10),
	}
}

func (sm *SeqMap) Get(key string) int {
	var keyCh = make(chan string, 1)
	var valueCh = make(chan int, 1)
	var req = func() {
		valueCh <- sm.m[<-keyCh]
	}
	sm.sc.Req(req)
	keyCh <- key
	return <-valueCh
}

func (sm *SeqMap) Put(key string, value int) {
	var keyCh = make(chan string, 1)
	var valueCh = make(chan int, 1)
	var okCh = make(chan struct{}, 1)
	var req = func() {
		sm.m[<-keyCh] = <-valueCh
		okCh <- struct{}{}
	}
	sm.sc.Req(req)
	keyCh <- key
	valueCh <- value
	<-okCh
	return
}

func TestSeqMap(t *testing.T) {
	var m = NewSeqMap()
	for i := 0; i < 100; i++ {
		var k = "T"
		var v = 128 * i
		go func() { t.Logf("get: key = %v, v = %v\n", k, m.Get(k)) }()
		go func() { m.Put(k, v) }()
	}
}
