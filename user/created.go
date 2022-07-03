package user

import (
	"cmd/client/internal/pkg/event"
	"context"
	"log"
	"time"
)

const Created event.Name = "user.created"

type CreatedEvent struct {
	Time time.Time
	ID   string
}

func (e CreatedEvent) Handle(ctx context.Context) {
	log.Printf("creating: %+v\n", e)
}
