// Copyright © 2018 Humio Ltd.
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
	"github.com/humio/cli/cmd/humioctl/internal/helpers"
	"github.com/spf13/cobra"
)

func newReposShowCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "show <repo>",
		Short: "Show details about a repository.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoName := args[0]
			client := NewApiClient(cmd)

			repo, err := client.Repositories().Get(repoName)
			helpers.ExitOnError(cmd, err, "Error fetching repository")

			printRepoDetailsTable(cmd, repo)
		},
	}

	return &cmd
}
