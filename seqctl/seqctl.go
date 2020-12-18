// Copyright (c) 2020, devgo.club
// All rights reserved.

package seqctl

import (
	"errors"
	"fmt"
)

// error
var (
	ErrParam = errors.New("invalid parameters")
)

// SeqCtl : sequence control to handle resource concurrently and safely
type SeqCtl interface {
	Req(func())
	String() string
	IsEmpty() bool
}

// New : create a new SeqCtl
func New(info string, threads int) SeqCtl {
	return initSeqCtl(info, threads)
}

type seqctl struct {
	info string
	ch   chan func()
}

func (ctrl *seqctl) run() {
	go func() {
		for f := range ctrl.ch {
			f()
		}
	}()
}

func (ctrl *seqctl) Req(f func()) {
	go func() {
		ctrl.ch <- f
	}()
}

func (ctrl *seqctl) String() string {
	return fmt.Sprintf("Controller: {info = %v, threads = %v, waiting = %v}",
		ctrl.info, cap(ctrl.ch), len(ctrl.ch))
}

func (ctrl *seqctl) IsEmpty() bool {
	return len(ctrl.ch) == 0
}

func initSeqCtl(info string, threads int) *seqctl {
	if len(info) == 0 || threads <= 0 {
		panic(ErrParam)
	}
	var ctrl = seqctl{
		info: info,
		ch:   make(chan func(), threads),
	}
	ctrl.run()
	return &ctrl
}
