// Copyright © 2020 Humio Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

func newActionsExportCmd() *cobra.Command {
	var outputName string

	cmd := cobra.Command{
		Use:   "export [flags] <repo-or-view> <action>",
		Short: "Export an action <action> in <repo-or-view> to a file.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			repoOrViewName := args[0]
			actionName := args[1]
			client := NewApiClient(cmd)

			if outputName == "" {
				outputName = actionName
			}

			action, err := client.Actions().Get(repoOrViewName, actionName)
			exitOnError(cmd, err, "Error fetching action")

			yamlData, err := yaml.Marshal(&action)
			exitOnError(cmd, err, "Failed to serialize the action")

			outFilePath := outputName + ".yaml"
			err = ioutil.WriteFile(outFilePath, yamlData, 0600)
			exitOnError(cmd, err, "Error saving the action file")
		},
	}

	cmd.Flags().StringVarP(&outputName, "output", "o", "", "The file path where the action should be written. Defaults to ./<action-name>.yaml")

	return &cmd
}
