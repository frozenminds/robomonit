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

package raspi

import (
	"github.com/frozenminds/robomonit/machine"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

// Make sure RaspiDirectLed implements Machine
var _ machine.Machine = (*RaspiDirectLed)(nil)

// Raspberry Direct LED
type RaspiDirectLed struct {
	name  string
	Robot *gobot.Gobot
}

// NewRaspiDirectLed returns a new Machine.
func NewRaspiDirectLed(pins map[string]string) *RaspiDirectLed {

	name := "cyborg"

	gbot := gobot.NewGobot()
	adapter := raspi.NewRaspiAdaptor("raspi")

	// Robot devices
	var devicesRobot []gobot.Device

	for name, pin := range pins {
		devicesRobot = append(devicesRobot, gpio.NewDirectPinDriver(adapter, name, pin))
	}

	// Setup robot
	robot := gobot.NewRobot(name,
		[]gobot.Connection{adapter},
		devicesRobot,
	)

	gbot.AddRobot(robot)

	machine := &RaspiDirectLed{
		name:  name,
		Robot: gbot,
	}

	return machine
}

// Get name
func (m *RaspiDirectLed) Name() string {
	return m.name
}

// Start
func (m *RaspiDirectLed) Start() []error {
	return m.Robot.Start()
}

// Stop
func (m *RaspiDirectLed) Stop() []error {
	return m.Robot.Stop()
}

// Reset everything
func (m *RaspiDirectLed) Reset() {
	robot := m.Robot.Robot(m.Name())

	robot.Devices().Each(func(device gobot.Device) {
		robot.Device(device.Name()).(*gpio.DirectPinDriver).Off()
	})
}

// Add work
func (m *RaspiDirectLed) Work(work func()) {
	m.Robot.Robot(m.Name()).Work = work
}

// Get device
func (m *RaspiDirectLed) Device(name string) gobot.Device {
	return m.Robot.Robot(m.Name()).Device(name)
}
