package calendar

type EventRepository interface {
	Save(event Event)
}
