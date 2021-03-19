/*
 *
 * Copyright 2021 The Vitess Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * /
 */

package infra

import (
	"github.com/spf13/cobra"
	"github.com/vitessio/arewefastyet/go/infra"
)

const (
	flagInfraPath = "infra-path"
)

func InfraCmd() *cobra.Command {
	var cfg infra.Config

	cmd := &cobra.Command{
		Use:     "infra <command>",
		Short:   "Manage infrastructure",
		Aliases: []string{"i"},
	}

	cmd.PersistentFlags().StringVar(&cfg.Path, flagInfraPath, "", "Path to the infra directory")
	cmd.AddCommand(create(&cfg))
	return cmd
}