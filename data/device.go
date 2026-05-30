package data

import (
	"fmt"
)

type Device interface {
	GetTimeDifference() int64
	GetType() string
}

type device struct {
	Time int64
	Type string
}

func NewDevice(timeDifference int64, deviceType string) Device {
	d := &device{
		Time: timeDifference,
		Type: deviceType,
	}
	fmt.Println("Neues Gerät erstellt: Time =", d.Time, ", Type =", d.Type)
	return d
}

func (d *device) GetTimeDifference() int64 {
	return d.Time
}

func (d *device) GetType() string {
	return d.Type
}
