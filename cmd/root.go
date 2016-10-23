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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "robomonit",
	Short: "Let robots monitor your apps and files.",
	Long: `Robot Monitor let's you monitor files, logs or apps using robots.
	
Use Raspberry, Arduino, Beaglebone or even your drone do fancy stuff depending on your monitoring needs.
The current working example is a RaspberryPi blinking LED's when certain patterns match.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is ./robomonit.yaml, falling back to $HOME/robomonit.yaml and /etc/robomonit.yaml).")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("CfgFile: ", cfgFile)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("robomonit") // name of config file (without extension)
	viper.AddConfigPath(".")         // adding current directory as first search path
	viper.AddConfigPath("$HOME")     // adding home directory as backup search path
	viper.AddConfigPath("/etc/")     // adding /etc/ directory as backup search path
	//	viper.AutomaticEnv()                  // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	err := viper.ReadInConfig()
	fmt.Println(err)
}
