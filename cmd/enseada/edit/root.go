// Copyright 2019 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package edit

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "edit [resource]",
	Short: "Edit a resource",
}

func init() {
	RootCmd.AddCommand(client)
}
