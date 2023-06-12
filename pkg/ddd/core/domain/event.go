package domain

import (
	"time"
)

// Event DDDにおけるドメインイベントの概念
type Event interface {
	OccurredOn() time.Time
	SameEventAs(other Event) bool
}
