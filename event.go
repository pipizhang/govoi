package govoi

const (
	EventLock    = "lock"
	EventUnlock  = "unlock"
	EventDrop    = "drop"
	EventCollect = "collect"

	EventBatteryCheck = "battery_check"
	EventUnknowCheck  = "unknow_check"
	EventServiceMode  = "service_mode"
)

// Event contains Event information, including transition hooks
type Event struct {
	Name        string
	transitions []*EventTransition
}

// To define EventTransition of go to a state
func (event *Event) To(name string) *EventTransition {
	transition := &EventTransition{to: name}
	event.transitions = append(event.transitions, transition)
	return transition
}
