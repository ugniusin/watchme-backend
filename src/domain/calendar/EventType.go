package calendar

type EventType string

const (
	Work      EventType = "work"
	Education EventType = "education"
	Health    EventType = "health"
	FreeTime  EventType = "free-time"
)
