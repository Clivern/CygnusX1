// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/clivern/peacock/core/module"
	"github.com/clivern/peacock/core/service"

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

		fs := service.NewFileSystem()

		spinner := module.NewCharmSpinner("Installing peacock!")

		go func() {
			// Create $HOME/.peacock/cache
			err := fs.EnsureDir(fmt.Sprintf("%s/.peacock/cache", HOME), 0755)

			if err != nil {
				fmt.Println(fmt.Sprintf(
					"Error while creating ~/.peacock/cache: %s",
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

		fmt.Println("\n")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
