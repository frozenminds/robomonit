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
// THE SOFTWARE.

package cmd

import (
	"os"
	"time"

	"github.com/frozenminds/robomonit/machine/raspi"
	"github.com/frozenminds/robomonit/monitor"

	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pipeCmd represents the monitor command
var pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Piped data monitoring.",
	Long:  `Read data from a pipe and transform it into hardware commands.`,
	Run: func(cmd *cobra.Command, args []string) {

		pins := pins()
		patterns := patterns()

		machine := raspi.NewRaspiDirectLed(pins)

		work := func() {

			machine.Reset()

			action := func(pin string) {
				device := machine.Device(pin).(*gpio.DirectPinDriver)

				device.On()

				time.Sleep(300 * time.Millisecond)
				device.Off()
			}

			monitor.Monitor(os.Stdin, patterns, action)
		}

		machine.Work(work)
		defer machine.Stop()
		machine.Start()
	},
}

func init() {
	RootCmd.AddCommand(pipeCmd)
}

// Get pins
func pins() map[string]string {
	return viper.GetStringMapString("pins")
}

// Get patterns
func patterns() map[string]string {
	return viper.GetStringMapString("patterns")
}
