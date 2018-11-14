package govoi

// EventTransition hold event's to/froms states, also including befores, afters hooks
type EventTransition struct {
	to      string
	froms   []string
	befores []hookFn
	afters  []hookFn
}

// From used to define from states
func (transition *EventTransition) From(states ...string) *EventTransition {
	transition.froms = states
	return transition
}

// Before register before hooks
func (transition *EventTransition) Before(fn hookFn) *EventTransition {
	transition.befores = append(transition.befores, fn)
	return transition
}

// After register after hooks
func (transition *EventTransition) After(fn hookFn) *EventTransition {
	transition.afters = append(transition.afters, fn)
	return transition
}
