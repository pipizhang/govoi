package govoi

import (
	"fmt"
	"log"
)

// VehicleStateMachine a struct that hold StateMachine and RoleManager
type VehicleStateMachine struct {
	FSM         *StateMachine
	RoleManager *RoleManager
}

// New return a new VehicleStateMachine
func New() *VehicleStateMachine {
	vsm := VehicleStateMachine{
		FSM:         NewStateMachine(),
		RoleManager: NewRoleManager(),
	}
	vsm.Init()
	return &vsm
}

func sendBetteryLowMessage(vehicleId int) {
	msg := fmt.Sprintf("battery low, vehicle '%d' needs to charge", vehicleId)
	go func(msg string) {
		log.Println(msg)
	}(msg)
}

func unlockTimeCheck() hookFn {
	return func(value Stater, args ...interface{}) error {
		if value.GetState() == StateServiceMode {
			return nil
		}

		order := value.(*Order)
		if order.Vehicle.LocalTime.Hour() >= 21 || order.Vehicle.LocalTime.Hour() <= 5 {
			return fmt.Errorf("Service off during 9pm ~ 5am")
		}
		return nil
	}
}

func betteryCheck() hookFn {
	return func(value Stater, args ...interface{}) error {
		if value.GetState() == StateServiceMode {
			return nil
		}

		order := value.(*Order)
		if order.Vehicle.Bettery < 20 {
			sendBetteryLowMessage(order.Vehicle.Id)
			return nil
		}
		return fmt.Errorf("Battery >= 20%")
	}
}

func saveVehicleState(order *Order) {
	//TODO
	log.Println("save vehicle state")
}

// Init initialize VehicleStateMachine
func (v *VehicleStateMachine) Init() {
	// Register states
	v.FSM.State(StateReady)
	v.FSM.State(StateBatteryLow)
	v.FSM.State(StateRiding)
	v.FSM.State(StateCollected)
	v.FSM.State(StateDropped)
	v.FSM.State(StateUnknown)
	v.FSM.State(StateTerminated)

	// Reigster events
	v.FSM.Event(EventUnlock).To(StateRiding).From(StateReady, StateServiceMode).
		Before(unlockTimeCheck())
	v.FSM.Event(EventLock).To(StateReady).From(StateRiding, StateServiceMode)
	v.FSM.Event(EventCollect).To(StateCollected).From(StateReady, StateBatteryLow, StateServiceMode)
	v.FSM.Event(EventDrop).To(StateDropped).From(StateCollected, StateServiceMode)
	v.FSM.Event(EventServiceMode).To(StateServiceMode).From(
		StateReady, StateBatteryLow, StateRiding, StateCollected, StateDropped, StateUnknown, StateTerminated,
	)

	v.FSM.Event(EventUnknowCheck).To(StateUnknown).From(StateReady, StateServiceMode)
	v.FSM.Event(EventBetteryCheck).To(StateBatteryLow).From(StateRiding, StateServiceMode).
		Before(betteryCheck())
}

// Fire fire an event
func (v *VehicleStateMachine) Fire(event string, order *Order) (err error) {
	_, err = v.RoleManager.IsAllow(order.GetRole(), event)
	if err != nil {
		return err
	}

	if order.GetRole() == RoleAdmin {
		order.SetState(StateServiceMode)
	}

	err = v.FSM.Fire(event, order)
	if err != nil {
		saveVehicleState(order)
	}

	return err
}
