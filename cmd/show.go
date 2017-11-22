// Copyright 2017 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ksonnet/ksonnet/metadata"
	"github.com/ksonnet/ksonnet/pkg/kubecfg"
)

const (
	flagFormat = "format"
)

func init() {
	RootCmd.AddCommand(showCmd)
	addEnvCmdFlags(showCmd)
	bindJsonnetFlags(showCmd)
	showCmd.PersistentFlags().StringP(flagFormat, "o", "yaml", "Output format.  Supported values are: json, yaml")
}

var showCmd = &cobra.Command{
	Use:   "show [<env>|-f <file-or-dir>]",
	Short: "Show expanded resource definitions",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("'show' requires an environment name; use `env list` to see available environments")
		}
		env := args[0]

		flags := cmd.Flags()
		var err error

		componentNames, err := flags.GetStringArray(flagComponent)
		if err != nil {
			return err
		}

		c := kubecfg.ShowCmd{}

		c.Format, err = flags.GetString(flagFormat)
		if err != nil {
			return err
		}

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		wd := metadata.AbsPath(cwd)

		objs, err := expandEnvCmdObjs(cmd, env, componentNames, wd)
		if err != nil {
			return err
		}

		return c.Run(objs, cmd.OutOrStdout())
	},
}
