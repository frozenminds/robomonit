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
	"errors"
	DEATH "github.com/vrecan/death"
	"io"
	"os"
	SYS "syscall"

	"github.com/frozenminds/robomonit/machine"
	"github.com/frozenminds/robomonit/monitor"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Piped data monitoring.",
	Long:  `Read data from a pipe and transform it into hardware commands.`,
	// Error Handling
	PreRunE: func(cmd *cobra.Command, args []string) error {

		platforms := platforms()
		if len(platforms) == 0 {
			return errors.New("No platforms defined in configuration file.")
		}

		patterns := patterns()
		if len(patterns) == 0 {
			return errors.New("No patterns defined in configuration file.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		// Signals to end app
		death := DEATH.NewDeath(SYS.SIGINT, SYS.SIGTERM)
		close := make([]io.Closer, 0)

		platforms := platforms()
		patterns := patterns()

		m := machine.NewMachineFromConfig(platforms)
		close = append(close, m)

		m.Reset()
		work := func() {
			action := func(id string) {
				m.DefaultAction(id)
			}

			monitor.Monitor(os.Stdin, patterns, action)
		}
		m.Work(work)
		m.Start()

		death.WaitForDeath(close...)
	},
}

func init() {

	RootCmd.AddCommand(monitorCmd)
}

// Get platforms
func platforms() map[string]interface{} {
	return viper.GetStringMap("platforms")
}

// Get patterns
func patterns() map[string]string {
	return viper.GetStringMapString("patterns")
}
