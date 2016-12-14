// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package walk

import (
	"fmt"
)

type KeyDownEventHandler func(key Key)

type KeyDownEvent struct {
	handlers []KeyDownEventHandler
}

func (e *KeyDownEvent) Attach(handler KeyDownEventHandler) int {
	e.handlers = append(e.handlers, handler)
	return len(e.handlers)

	for i, h := range e.handlers {
		if h == nil {
			e.handlers[i] = handler
			return i
		}
	}
	fmt.Println("im goto attch")
	e.handlers = append(e.handlers, handler)
	return len(e.handlers) - 1
}

func (e *KeyDownEvent) Detach(handle int) {
	e.handlers[handle] = nil
}

type KeyDownEventPublisher struct {
	event KeyDownEvent
}

func (p *KeyDownEventPublisher) Event() *KeyDownEvent {
	return &p.event
}

func (p *KeyDownEventPublisher) Publish(key Key) {
	fmt.Println("my key publish", p.event.handlers)
	for _, handler := range p.event.handlers {
		if handler != nil {
			handler(key)
		}
	}
}
