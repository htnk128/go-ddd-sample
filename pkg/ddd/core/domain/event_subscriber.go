package domain

// EventSubscriber Event の購読インターフェース
type EventSubscriber interface {
	Handle(event Event)
}
