package event

import (
	"context"
	"fmt"
)

type Dispatcher struct {
	events map[Name]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		events: make(map[Name]Listener),
	}
}

func (d *Dispatcher) Register(listener Listener, names ...Name) error {
	for _, name := range names {
		if _, ok := d.events[name]; ok {
			return fmt.Errorf("the '%s' event is already registered", name)
		}

		d.events[name] = listener
	}

	return nil
}

func (d *Dispatcher) Dispatch(ctx context.Context, name Name, event interface{}) error {
	if _, ok := d.events[name]; !ok {
		return fmt.Errorf("the '%s' event is not registered", name)
	}

	d.events[name].Listen(ctx, event)

	return nil
}
