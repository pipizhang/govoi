package govoi

import (
	"time"
)

type StateOrder interface {
	GetRole() string
}

type (
	Order struct {
		Transition
		Id      int
		User    *User
		Vehicle *Vehicle
	}

	User struct {
		Id   int
		Role string
		// TODO add more fields
	}

	Vehicle struct {
		Id        int
		Bettery   int // (0~100)
		State     string
		LocalTime time.Time
		// TODO add more fields
	}
)

// GetRole get user's role
func (o *Order) GetRole() string {
	return o.User.Role
}

// GetState get vehicle state
func (o *Order) GetState() string {
	return o.Vehicle.State
}

// SetState set vehicle state
func (o *Order) SetState(state string) {
	o.Vehicle.State = state
}
