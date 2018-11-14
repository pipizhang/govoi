package govoi

const (
	StateReady       = "ready"
	StateBatteryLow  = "bettery_low"
	StateRiding      = "riding"
	StateCollected   = "collected"
	StateDropped     = "dropped"
	StateUnknown     = "unknown"
	StateTerminated  = "terminated"
	StateServiceMode = "service_mode"
)

// hookFn state machine callback function
type hookFn func(value Stater, args ...interface{}) error

// State contains State information, including enter, exit hooks
type State struct {
	Name   string
	enters []hookFn
	exits  []hookFn
}

// OnEnter register an enter hook of State
func (state *State) OnEnter(fn hookFn) *State {
	state.enters = append(state.enters, fn)
	return state
}

// OnExit register an exit hook of State
func (state *State) OnExit(fn hookFn) *State {
	state.exits = append(state.exits, fn)
	return state
}
