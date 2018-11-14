package govoi

import (
	"fmt"
)

// StateMachine a struct that hold states, events definitions
type StateMachine struct {
	initialState string
	states       map[string]*State
	events       map[string]*Event
}

// New initialize a new StateMachine that hold states, events definitions
func NewStateMachine() *StateMachine {
	return &StateMachine{
		states: map[string]*State{},
		events: map[string]*Event{},
	}
}

// Initial define the initial state
func (sm *StateMachine) Initial(name string) *StateMachine {
	sm.initialState = name
	return sm
}

// State define a state
func (sm *StateMachine) State(name string) *State {
	state := &State{Name: name}
	sm.states[name] = state
	return state
}

// Event define an event
func (sm *StateMachine) Event(name string) *Event {
	event := &Event{Name: name}
	sm.events[name] = event
	return event
}

// Fire fire an event
func (sm *StateMachine) Fire(name string, value Stater, notes ...string) error {
	var stateWas = value.GetState()

	if stateWas == "" {
		stateWas = sm.initialState
		value.SetState(sm.initialState)
	}

	if event := sm.events[name]; event != nil {
		var matchedTransitions []*EventTransition
		for _, transition := range event.transitions {
			var validFrom = len(transition.froms) == 0
			if len(transition.froms) > 0 {
				for _, from := range transition.froms {
					if from == stateWas {
						validFrom = true
					}
				}
			}

			if validFrom {
				matchedTransitions = append(matchedTransitions, transition)
			}
		}

		if len(matchedTransitions) == 1 {
			transition := matchedTransitions[0]

			// State: exit
			if state, ok := sm.states[stateWas]; ok {
				for _, exit := range state.exits {
					if err := exit(value); err != nil {
						return err
					}
				}
			}

			// Transition: before
			for _, before := range transition.befores {
				if err := before(value); err != nil {
					return err
				}
			}

			value.SetState(transition.to)

			// State: enter
			if state, ok := sm.states[transition.to]; ok {
				for _, enter := range state.enters {
					if err := enter(value); err != nil {
						value.SetState(stateWas)
						return err
					}
				}
			}

			// Transition: after
			for _, after := range transition.afters {
				if err := after(value); err != nil {
					value.SetState(stateWas)
					return err
				}
			}

			return nil
		}
	}
	return fmt.Errorf("Failed to perform event '%s' from state '%s'", name, stateWas)
}
