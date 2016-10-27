// Copyright © 2016 Constantin Bejenaru <boby@frozenminds.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE

package machine

import (
	"io"
	"log"
	"strings"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/spf13/cast"

	"github.com/hybridgroup/gobot/platforms/gpio"

	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot/platforms/chip"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

// MachineFactory struct, implements Machine
type MachineFactory struct {
	Gobot *gobot.Gobot
}

// Make sure MachineFactory implements Machine
var _ Machine = (*MachineFactory)(nil)

// Make sure MachineFactory implements io.Closer
var _ io.Closer = (*MachineFactory)(nil)

// NewMachineFromConfig will return an instance of Machine from the usual configurations.
// For more control, safety, drivers and so on, please create your own implementation.
func NewMachineFromConfig(config map[string]interface{}) *MachineFactory {

	gbot := gobot.NewGobot()
	machine := &MachineFactory{
		Gobot: gbot,
	}

	if len(config)-1 <= 0 {
		log.Println("No platform configured!")
	}

	for name, drivers := range config {

		var adaptor gobot.Adaptor

		switch name {
		case "arduino":
			adaptor = firmata.NewFirmataAdaptor(name, "/dev/ttyACM0")
			break

		case "beaglebone":
			adaptor = beaglebone.NewBeagleboneAdaptor(name)
			break

		case "chip":
			adaptor = chip.NewChipAdaptor(name)
			break

		case "edison":
			adaptor = edison.NewEdisonAdaptor(name)
			break

		case "raspi":
			adaptor = raspi.NewRaspiAdaptor(name)
			break
		}

		robot := gobot.NewRobot(
			name,
			[]gobot.Connection{adaptor},
		)

		machine.Gobot.AddRobot(robot)

		machine.AddDevicesFromConfig(cast.ToStringMap(drivers))
	}

	return machine
}

// Start
func (m *MachineFactory) Start() []error {
	err := m.Gobot.Start()

	m.Reset()

	return err
}

// Stop
func (m *MachineFactory) Stop() []error {
	m.Reset()

	return m.Gobot.Stop()
}

func (m *MachineFactory) Close() error {
	m.Stop()
	return nil
}

// Reset
func (m *MachineFactory) Reset() {

	m.Gobot.Robots().Each(func(robot *gobot.Robot) {
		robot.Devices().Each(func(device gobot.Device) {

			switch device.(type) {
			case *gpio.DirectPinDriver:
				device.(*gpio.DirectPinDriver).Off()
				break

			case *gpio.LedDriver:
				device.(*gpio.LedDriver).Off()
				break
			}
		})
	})
}

// Add work
func (m *MachineFactory) Work(work func()) {
	m.Gobot.Robots().Each(func(robot *gobot.Robot) {
		robot.Work = work
	})
}

// Execute default action.
// This will turn a matching device On then Off.
func (m *MachineFactory) DefaultAction(identifier string) {

	ids := strings.Split(identifier, ".")

	if !keyInSlice(ids, 0) {
		return
	}
	robot := m.Gobot.Robot(ids[0])

	if !keyInSlice(ids, 2) {
		return
	}
	device := robot.Device(ids[2])

	// Excuse the ugly implementation.
	// Gobot's drivers do not implement a specific interface with On()/Off() or alike.
	// Will probably switch to local driver implementation (machine/driver.go).
	switch device.(type) {
	case *gpio.DirectPinDriver:

		action := func(d *gpio.DirectPinDriver) {
			d.On()
			sleep()
			d.Off()
		}
		go action(device.(*gpio.DirectPinDriver))
		break

	case *gpio.LedDriver:

		action := func(d *gpio.LedDriver) {
			d.On()
			sleep()
			d.Off()
		}

		go action(device.(*gpio.LedDriver))
		break
	}
}

// Add devices from configuration
func (m *MachineFactory) AddDevicesFromConfig(devices map[string]interface{}) {

	for driver, config := range devices {
		m.AddDeviceFromConfig(driver, cast.ToStringMapString(config))
	}
}

// Add device from configuration
func (m *MachineFactory) AddDeviceFromConfig(name string, config map[string]string) {

	m.Gobot.Robots().Each(func(robot *gobot.Robot) {

		robot.Connections().Each(func(connection gobot.Connection) {

			for drivername, number := range config {

				switch name {
				case "direct-pin":
					robot.AddDevice(gpio.NewDirectPinDriver(connection, drivername, number))
					break

				case "led-driver":
					robot.AddDevice(gpio.NewLedDriver(connection.(gpio.DigitalWriter), drivername, number))
					break
				}
			}
		})
	})
}

// Check if key is in slice
func keyInSlice(search []string, idx int) bool {
	if len(search)-1 >= idx {
		return true
	}
	return false
}

// Sleep
func sleep() {
	time.Sleep(300 * time.Millisecond)
}
