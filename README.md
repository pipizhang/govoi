# govoi
A small golang library that validates and handles state-transitions for an abstract vehicle.

[![Go Report Card](https://goreportcard.com/badge/github.com/pipizhang/govoi)](https://goreportcard.com/report/github.com/pipizhang/govoi)
[![GoDoc](https://godoc.org/github.com/pipizhang/govoi?status.svg)](https://godoc.org/github.com/pipizhang/govoi)

## Usage
### Basic example
```go
vsm := govoi.New()

order := &govoi.Order{
	User:    &govoi.User{Id: 1, Role: govoi.RoleEndUser},
	Vehicle: &govoi.Vehicle{Id: 1, State: govoi.StateReady, LocalTime: time.Now()},
}

err := vsm.Fire(govoi.EventUnlock, order)
if err != nil {
	fmt.Println(err)
}
```
