package user

import (
	"context"
	"log"
)

type Listener struct{}

func (u Listener) Listen(ctx context.Context, event interface{}) {
	switch event := event.(type) {
	case CreatedEvent:
		event.Handle(ctx)
	case UpdatedEvent:
		event.Handle(ctx)
	case DeletedEvent:
		event.Handle(ctx)
	default:
		log.Printf("registered an invalid user event: %T\n", event)
	}
}
