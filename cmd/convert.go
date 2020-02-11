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
	//"strings"
	"github.com/frase1sr/swarm_to_helm/pkg/app"
	"github.com/frase1sr/swarm_to_helm/pkg/kobject"
	//log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
// TODO: comment
var (
	ConvertServerURL                   string
	ConvertAuthToken				   string
	ConvertOpt                   kobject.ConvertOptions
	ConvertOut                    string

)

var convertCmd = &cobra.Command{
	Use:   "convert [outdir]",
	Short: "Create Helm Chart from Deployed Docker Swarm Services",
	PreRun: func(cmd *cobra.Command, args []string) {

		// Check that build-config wasn't passed in with --provider=kubernetes
		//if GlobalProvider == "kubernetes" && UpBuild == "build-config" {
		//	log.Fatalf("build-config is not a valid --build parameter with provider Kubernetes")
		//}

		// Create the Convert Options.
		ConvertOpt = kobject.ConvertOptions{
			ServerURL:                   ConvertServerURL,
			AuthToken:					 ConvertAuthToken,
			Out:							 ConvertOut,
		}

		// Validate before doing anything else. Use "bundle" if passed in.
		//app.ValidateFlags(GlobalBundle, args, cmd, &ConvertOpt)
		//app.ValidateComposeFile(&ConvertOpt)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("It's running")
		app.Convert(ConvertOpt)
	},
}

func init() {

	// Automatically grab environment variables
	viper.AutomaticEnv()

	convertCmd.Flags().StringVar(&ConvertServerURL, "server-url", "https://localhost:9443", "please specify the base url")
	convertCmd.Flags().StringVar(&ConvertAuthToken, "auth-token", "", "please specify the auth token")
	convertCmd.Flags().StringVarP(&ConvertOut, "out", "o", "", "Specify a file name or directory to save objects to (if path does not exist, a file will be created)")

	RootCmd.AddCommand(convertCmd)
}