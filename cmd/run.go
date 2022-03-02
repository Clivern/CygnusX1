// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a kafka clusters locally",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var runSingleNodeCmd = &cobra.Command{
	Use:   "single",
	Short: "Run a single kafka node locally",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run single node")
	},
}

func init() {
	runCmd.AddCommand(runSingleNodeCmd)
	rootCmd.AddCommand(runCmd)
}
