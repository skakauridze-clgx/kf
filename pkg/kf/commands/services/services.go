// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"github.com/google/kf/v2/pkg/kf/commands/config"
	"github.com/google/kf/v2/pkg/kf/internal/genericcli"
	"github.com/google/kf/v2/pkg/kf/serviceinstances"
	"github.com/spf13/cobra"
)

// NewListServicesCommand allows users to list service instances.
func NewListServicesCommand(p *config.KfParams) *cobra.Command {
	return genericcli.NewListCommand(
		serviceinstances.NewResourceInfo(),
		p,
		genericcli.WithListPluralFriendlyName("services"),
		genericcli.WithListCommandName("services"),
		genericcli.WithListAliases([]string{"s"}),
	)
}
