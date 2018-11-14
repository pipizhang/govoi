package govoi

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getLocalTime(s string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	str := fmt.Sprintf("2018-11-14T%s.371Z", s)
	t, _ := time.Parse(layout, str)
	return t
}

func TestEndUser(t *testing.T) {
	var err error

	vsm := New()

	order_1 := &Order{
		User:    &User{Id: 1, Role: RoleEndUser},
		Vehicle: &Vehicle{Id: 1, State: StateReady, LocalTime: getLocalTime("10:01:01")},
	}
	err = vsm.Fire(EventUnlock, order_1)
	assert.Nil(t, err)

	order_2 := &Order{
		User:    &User{Id: 1, Role: RoleEndUser},
		Vehicle: &Vehicle{Id: 1, State: StateReady, LocalTime: getLocalTime("10:01:01")},
	}
	err = vsm.Fire(EventCollect, order_2)
	assert.EqualError(t, err, ErrPermissionDenied.Error())

	order_3 := &Order{
		User:    &User{Id: 1, Role: RoleEndUser},
		Vehicle: &Vehicle{Id: 1, State: StateReady, LocalTime: getLocalTime("22:30:00")},
	}
	err = vsm.Fire(EventUnlock, order_3)
	assert.EqualError(t, err, ErrServiceOff.Error())
}

func TestHunter(t *testing.T) {
	var err error

	vsm := New()

	order_1 := &Order{
		User:    &User{Id: 1, Role: RoleHunter},
		Vehicle: &Vehicle{Id: 1, State: StateBatteryLow, LocalTime: getLocalTime("10:01:01")},
	}
	err = vsm.Fire(EventCollect, order_1)
	assert.Nil(t, err)

	order_2 := &Order{
		User:    &User{Id: 1, Role: RoleHunter},
		Vehicle: &Vehicle{Id: 1, State: StateRiding, LocalTime: getLocalTime("10:01:01")},
	}
	err = vsm.Fire(EventCollect, order_2)
	assert.EqualError(t, err, "Failed to perform event 'collect' from state 'riding'")
}

func TestAdmin(t *testing.T) {
	var err error

	vsm := New()

	order_1 := &Order{
		User:    &User{Id: 1, Role: RoleAdmin},
		Vehicle: &Vehicle{Id: 1, State: StateBatteryLow, LocalTime: getLocalTime("10:01:01")},
	}
	err = vsm.Fire(EventUnlock, order_1)
	assert.Nil(t, err)

	order_2 := &Order{
		User:    &User{Id: 1, Role: RoleAdmin},
		Vehicle: &Vehicle{Id: 1, State: StateReady, LocalTime: getLocalTime("23:01:01")},
	}
	err = vsm.Fire(EventUnlock, order_2)
	assert.Nil(t, err)
}

func TestBatteryCheck(t *testing.T) {
	var err error

	vsm := New()

	order_1 := &Order{
		User:    &User{Id: 1, Role: RoleAdmin},
		Vehicle: &Vehicle{Id: 1, State: StateBatteryLow, LocalTime: getLocalTime("10:01:01"), Battery: 10},
	}
	err = vsm.Fire(EventBatteryCheck, order_1)
	assert.Nil(t, err)
	assert.Equal(t, order_1.GetState(), StateBatteryLow)

	order_2 := &Order{
		User:    &User{Id: 1, Role: RoleAdmin},
		Vehicle: &Vehicle{Id: 1, State: StateReady, LocalTime: getLocalTime("10:01:01"), Battery: 85},
	}
	err = vsm.Fire(EventBatteryCheck, order_2)
	assert.Nil(t, err)
}
