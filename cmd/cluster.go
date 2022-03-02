// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/clivern/peacock/core/definition"
	"github.com/clivern/peacock/core/driver"
	"github.com/clivern/peacock/core/module"
	"github.com/clivern/peacock/core/runtime"

	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage kafka clusters",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var destroyCmd = &cobra.Command{
	Use:   "destroy [name]",
	Short: "Destroy local kafka cluster",
	Run: func(cmd *cobra.Command, args []string) {

		db := driver.NewDatabase(fmt.Sprintf(
			"%s/.peacock/peacock.db",
			HOME,
		))

		if len(args) == 0 {
			fmt.Println("Error! cluster name is required!")
			os.Exit(1)
		}

		clusterName := strings.TrimSpace(args[0])

		if clusterName == "" {
			fmt.Println("Error! cluster name is required!")
			os.Exit(1)
		}

		val, _ := db.FindByKey(clusterName)

		if val == "" {
			fmt.Println("Error! cluster name is missing!")
			os.Exit(1)
		}
	},
}

var runCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Run a kafka cluster locally for testing",
	Run: func(cmd *cobra.Command, args []string) {

		db := driver.NewDatabase(fmt.Sprintf(
			"%s/.peacock/peacock.db",
			HOME,
		))

		if len(args) == 0 {
			fmt.Println("Error! cluster name is required!")
			os.Exit(1)
		}

		clusterName := strings.TrimSpace(args[0])

		if clusterName == "" {
			fmt.Println("Error! cluster name is required!")
			os.Exit(1)
		}

		val, _ := db.FindByKey(clusterName)

		if val != "" {
			fmt.Println("Error! cluster name is already used!")
			os.Exit(1)
		}

		id := definition.GetServiceID()

		dc := runtime.NewDockerCompose(fmt.Sprintf(
			"%s/.peacock/cache",
			HOME,
		))

		config := definition.GetKafkaConfig(
			definition.ZookeeperDockerVersion,
			definition.KafkaDockerVersion,
		)

		spinner := module.NewCharmSpinner("Setting up a new kafka cluster!")

		go func() {
			err := dc.Deploy(id, config)

			if err != nil {
				fmt.Println(fmt.Sprintf(
					"Error while setting up kafka cluster: %s",
					err.Error(),
				))
			}

			spinner.Quit()
		}()

		if err := spinner.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		port, err := dc.FetchServicePort(id, definition.KafkaService, definition.KafkaPort, config)

		if err != nil {
			fmt.Println(fmt.Sprintf("Error! unable to fetch the port: %s", err.Error()))
			os.Exit(1)
		}

		err = db.Insert(clusterName, fmt.Sprintf(
			`{"id":"%s", "type": "local", "port": "%s"}`,
			id,
			port,
		))

		if err != nil {
			fmt.Printf("Error raised: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println("")
	},
}

func init() {
	clusterCmd.AddCommand(runCmd)
	rootCmd.AddCommand(clusterCmd)
}
