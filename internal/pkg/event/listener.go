package event

import (
	"context"
)

// All custom event listeners must satisfy this interface.
type Listener interface {
	Listen(ctx context.Context, event interface{})
}
