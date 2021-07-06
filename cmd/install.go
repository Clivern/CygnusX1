// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/clivern/nitro/core/module"
	"github.com/clivern/nitro/core/service"

	"github.com/spf13/cobra"
)

// installCmd install kafka cluster command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs nitro command line locally.",
	Run: func(cmd *cobra.Command, args []string) {

		if HOME == "" {
			fmt.Println("Error! `HOME` environment variable is not set")
			os.Exit(1)
		}

		fs := service.NewFileSystem()

		spinner := module.NewCharmSpinner("Installing nitro!")

		go func() {
			// Create $HOME/.nitro/cache
			err := fs.EnsureDir(fmt.Sprintf("%s/.nitro/cache", HOME), 0755)

			if err != nil {
				fmt.Println(fmt.Sprintf(
					"Error while creating ~/.nitro/cache: %s",
					err.Error(),
				))
			}

			time.Sleep(1 * time.Second)

			spinner.Quit()
		}()

		if err := spinner.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
