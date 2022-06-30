package main

import (
	"cmd/client/internal/pkg/event"
	"cmd/client/user"
	"context"
	"log"
	"time"
)

func main() {
	// ----------------------------------------------------------------------------------------
	// Register events with listeners at application boot
	dispatcher := event.NewDispatcher()
	if err := dispatcher.Register(user.Listener{}, user.Created, user.Updated); err != nil {
		log.Fatalln(err)
	}

	// ----------------------------------------------------------------------------------------
	// Dispatch registered events. Valid.
	go func() {
		err := dispatcher.Dispatch(context.Background(), user.Created, user.CreatedEvent{
			Time: time.Now().UTC(),
			ID:   "111",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		err := dispatcher.Dispatch(context.Background(), user.Updated, user.UpdatedEvent{
			Time:     time.Now().UTC(),
			ID:       "222",
			Key:      "name",
			OldValue: "inanzzz",
			NewValue: "zzznani",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	// ----------------------------------------------------------------------------------------
	// Dispatch a valid event type to unregistered event name. Error.
	go func() {
		err := dispatcher.Dispatch(context.Background(), user.Deleted, user.DeletedEvent{
			Time: time.Now().UTC(),
			ID:   "333",
			Who:  "admin",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	// ----------------------------------------------------------------------------------------
	// Dispatch a wrong event type to registered event name. Error.
	go dispatcher.Dispatch(context.Background(), user.Created, nil)
	go dispatcher.Dispatch(context.Background(), user.Updated, "hi")
	go dispatcher.Dispatch(context.Background(), user.Created, 123)
	go dispatcher.Dispatch(context.Background(), user.Updated, struct{}{})
	go dispatcher.Dispatch(context.Background(), user.Created, make(chan int))

	select {}
}
