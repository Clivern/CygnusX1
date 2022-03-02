// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs peacock command line locally.",
	Run: func(cmd *cobra.Command, args []string) {

		if HOME == "" {
			fmt.Println("Error! `HOME` environment variable is not set")
			os.Exit(1)
		}

		fmt.Println("Peacock installed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
