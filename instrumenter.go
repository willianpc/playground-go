package main

import (
	"time"
)

type instrumenter struct {
	sp *span
	fn func(span)
	ch chan bool
}

func (in *instrumenter) begin() {
	go func() {
		if in.sp == nil {
			panic("no span to begin")
		}

		in.sp.ts = time.Now().UnixMilli()
		in.fn(*in.sp)
		in.ch <- true
	}()
}

func (in instrumenter) end() {
	now := time.Now().UnixMilli()

	<-in.ch

	if in.sp == nil {
		panic("no span to finish")
	}

	in.sp.time = now - in.sp.ts

	in.sp.finish()
}

func newInstrumenter(spanType string, fn func(sp span)) *instrumenter {
	sp := startSpan(spanType)
	return &instrumenter{fn: fn, ch: make(chan bool), sp: sp}
}
