// Copyright 2019 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"os"

	"github.com/enseadaio/enseada/cmd/enseada/root"
)

func main() {
	err := root.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
