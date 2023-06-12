package domain

// EventPublisher Event の出版インターフェース
type EventPublisher interface {
	Publish(event Event)
}
