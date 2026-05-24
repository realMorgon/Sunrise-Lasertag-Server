package data

import (
	"fmt"
)

type Device interface {
	GetTime() int
	GetType() string
}

type device struct {
	Time int
	Type string
}

func NewDevice(time int, deviceType string) Device {
	d := &device{
		Time: time,
		Type: deviceType,
	}
	fmt.Println("Neues Gerät erstellt: Time =", d.Time, ", Type =", d.Type)
	return d
}

func (d *device) GetTime() int {
	return d.Time
}

func (d *device) GetType() string {
	return d.Type
}
