package govoi

// Stater is a interface including methods `GetState`, `SetState`
type Stater interface {
	SetState(name string)
	GetState() string
}

// Transition is a struct, embed it in your struct to enable state machine for the struct
type Transition struct {
	State string
}

// SetState set state to Stater
func (transition *Transition) SetState(name string) {
	transition.State = name
}

// GetState get current state from
func (transition Transition) GetState() string {
	return transition.State
}
