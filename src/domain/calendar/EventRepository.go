package calendar

type EventRepository interface {
	FindOne (id int) (Event, error)
	Save (event Event) int
}
