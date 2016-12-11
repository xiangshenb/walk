// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package walk

type MoveEventHandler func()

type MoveEvent struct {
	handlers []MoveEventHandler
}

func (e *MoveEvent) Attach(handler MoveEventHandler) int {
	for i, h := range e.handlers {
		if h == nil {
			e.handlers[i] = handler
			return i
		}
	}

	e.handlers = append(e.handlers, handler)
	return len(e.handlers) - 1
}

func (e *MoveEvent) Detach(handle int) {
	e.handlers[handle] = nil
}

type MoveEventPublisher struct {
	event MoveEvent
}

func (p *MoveEventPublisher) Event() *MoveEvent {
	return &p.event
}

func (p *MoveEventPublisher) Publish() {
	for _, handler := range p.event.handlers {
		if handler != nil {
			handler()
		}
	}
}
