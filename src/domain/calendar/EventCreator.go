package calendar

type EventCreator struct {
	eventRepository EventRepository
}

func NewEventCreator(eventRepository EventRepository) *EventCreator {
	return &EventCreator{
		eventRepository: eventRepository,
	}
}

func (eventCreator EventCreator) Create ()  {

	event := new(Event)
	event.Name = "Learn Ruby"
	event.From = "2019-07-08 00:08:00"
	event.To = "2019-07-08 00:09:00"
	event.Type = EventType(Education)

	eventCreator.eventRepository.Save(*event)
}
