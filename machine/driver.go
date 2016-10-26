// Copyright © 2016 Constantin Bejenaru <boby@frozenminds.com>c
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
	"time"

	"github.com/hybridgroup/gobot/platforms/gpio"
)

// Interface Driver which exposes common actions.
// Extend desired drivers to implement methods exposed in Driver.
type Driver interface {
	// Turn on
	On() (err error)

	// Turn off
	Off() (err error)

	// Flash
	Flash() (err error)
}

var _ Driver = (*DirectPinDriver)(nil)
var _ Driver = (*LedDriver)(nil)
var _ Driver = (*RgbLedDriver)(nil)

type DirectPinDriver struct {
	*gpio.DirectPinDriver
}
type LedDriver struct {
	*gpio.LedDriver
}

type RgbLedDriver struct {
	*gpio.RgbLedDriver
}

// Flash set's the LED's pins to their various states
func (d *DirectPinDriver) Flash() (err error) {

	if err = d.On(); err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)

	if err = d.Off(); err != nil {
		return err
	}

	return nil
}

// Flash set's the LED's pins to their various states
func (d *LedDriver) Flash() (err error) {

	if err = d.On(); err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)

	if err = d.Off(); err != nil {
		return err
	}

	return
}

// Flash set's the LED's pins to their various states
func (l *RgbLedDriver) Flash() (err error) {

	if err = l.On(); err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)

	if err = l.Off(); err != nil {
		return err
	}

	return
}
