/*
Copyright 2017 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	//"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Logrus hooks

// Hook for error and exit out on warning
type errorOnWarningHook struct{}

func (errorOnWarningHook) Levels() []log.Level {
	return []log.Level{log.WarnLevel}
}

func (errorOnWarningHook) Fire(entry *log.Entry) error {
	log.Fatalln(entry.Message)
	return nil
}

// TODO: comment
var (
	GlobalBundle           string
	GlobalProvider         string
	GlobalVerbose          bool
	GlobalSuppressWarnings bool
	GlobalErrorOnWarning   bool
	GlobalFiles            []string
)

// RootCmd root level flags and commands
var RootCmd = &cobra.Command{
	Use:   "swarm_to_helm",
	Short: "A tool helping Docker Swarm conversion to helm",
	Long:  `swarm_to_helm is a tool to help convert full Docker Swarm deplyoments to helm`,
	// PersistentPreRun will be "inherited" by all children and ran before *every* command unless
	// the child has overridden the functionality. This functionality was implemented to check / modify
	// all global flag calls regardless of app call.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Add extra logging when verbosity is passed
		if GlobalVerbose {
			log.SetLevel(log.DebugLevel)
		}

		// Disable the timestamp (Kompose is too fast!)
		formatter := new(log.TextFormatter)
		formatter.DisableTimestamp = true
		formatter.ForceColors = true
		log.SetFormatter(formatter)

		// Set the appropriate suppress warnings and error on warning flags
		if GlobalSuppressWarnings {
			log.SetLevel(log.ErrorLevel)
		} else if GlobalErrorOnWarning {
			hook := errorOnWarningHook{}
			log.AddHook(hook)
		}


	},
}

// Execute TODO: comment
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&GlobalVerbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVar(&GlobalSuppressWarnings, "suppress-warnings", false, "Suppress all warnings")
	RootCmd.PersistentFlags().BoolVar(&GlobalErrorOnWarning, "error-on-warning", false, "Treat any warning as an error")
	RootCmd.PersistentFlags().StringArrayVarP(&GlobalFiles, "file", "f", []string{}, "Specify an alternative compose file")
}
