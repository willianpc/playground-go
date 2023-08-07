package main

import (
	"fmt"
	"log"
	"time"
)

type span struct {
	spanType string
	time     int64
	ts       int64
	tags     map[string]string
}

func (sp *span) finish() {
	log.Println("Span Finished: ", sp)
}

func (sp span) String() string {
	return fmt.Sprintf("Span type: %s; Time: %dms; Tags: %v", sp.spanType, sp.time, sp.tags)
}

func startSpan(spanType string) *span {
	return &span{
		spanType: spanType,
		ts:       time.Now().UnixMilli(),
		tags:     make(map[string]string),
	}
}
