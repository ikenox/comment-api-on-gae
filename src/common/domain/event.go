package domain

type DomainEvent struct {
	EventType string
	Data map[string]string
}
