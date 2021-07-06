// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/clivern/nitro/core/definition"
	"github.com/clivern/nitro/core/driver"
	"github.com/clivern/nitro/core/model"
	"github.com/clivern/nitro/core/module"
	"github.com/clivern/nitro/core/runtime"

	"github.com/spf13/cobra"
)

// clusterCmd cluster sub-command
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

// listCmd list clusters sub-command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List kafka cluster",
	Run: func(cmd *cobra.Command, args []string) {

		db := driver.NewDatabase(fmt.Sprintf(
			"%s/.nitro/nitro.db",
			HOME,
		))

		records, err := db.FindAll()

		if err != nil {
			fmt.Printf("Error raised: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println(records)
		fmt.Println("")
	},
}

// showCmd list clusters sub-command
var showCmd = &cobra.Command{
	Use:   "show [name]",
	Short: "Show kafka cluster configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}

// addCmd list clusters sub-command
var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new kafka cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}

// destroyCmd destroy clutser sub-command
var destroyCmd = &cobra.Command{
	Use:   "destroy [name]",
	Short: "Destroy local kafka cluster",
	Run: func(cmd *cobra.Command, args []string) {

		db := driver.NewDatabase(fmt.Sprintf(
			"%s/.nitro/nitro.db",
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

		cluster, err := db.FindByKey(clusterName)

		if err != nil {
			fmt.Printf("Error raised: %s", err.Error())
			os.Exit(1)
		}

		dc := runtime.NewDockerCompose(fmt.Sprintf(
			"%s/.nitro/cache",
			HOME,
		))

		config := definition.GetKafkaConfig(
			definition.ZookeeperDockerVersion,
			definition.KafkaDockerVersion,
			cluster.Port,
		)

		spinner := module.NewCharmSpinner("Deleting kafka cluster!")

		go func() {
			err := dc.Destroy(cluster.ID, config)

			if err != nil {
				fmt.Println(fmt.Sprintf(
					"Error while deleting kafka cluster: %s",
					err.Error(),
				))
			}

			spinner.Quit()
		}()

		if err := spinner.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = db.DeleteByKey(clusterName)

		if err != nil {
			fmt.Printf("Error raised: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println("")
	},
}

// runCmd run kafka cluster locally sub-command
var runCmd = &cobra.Command{
	Use:   "run [name] [port]",
	Short: "Run a kafka cluster locally for testing",
	Run: func(cmd *cobra.Command, args []string) {

		db := driver.NewDatabase(fmt.Sprintf(
			"%s/.nitro/nitro.db",
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

		clust, _ := db.FindByKey(clusterName)

		if clust.Name != "" {
			fmt.Printf("Cluster with name %s exists", clust.Name)
			os.Exit(1)
		}

		if len(args) < 2 || strings.TrimSpace(args[1]) == "" {
			fmt.Println("Error! a free port number is required!")
			os.Exit(1)
		}

		freePort := args[1]

		id := definition.GetServiceID()

		dc := runtime.NewDockerCompose(fmt.Sprintf(
			"%s/.nitro/cache",
			HOME,
		))

		config := definition.GetKafkaConfig(
			definition.ZookeeperDockerVersion,
			definition.KafkaDockerVersion,
			freePort,
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

		port, err := dc.FetchServicePort(
			id,
			definition.KafkaService,
			freePort,
			config,
		)

		if err != nil {
			fmt.Println(fmt.Sprintf("Error! unable to fetch the port: %s", err.Error()))
			os.Exit(1)
		}

		cluster := &model.Cluster{
			ID:      id,
			Type:    "local",
			Port:    port,
			Address: fmt.Sprintf("localhost:%s", port),
			Name:    clusterName,
		}

		output, err := cluster.ConvertToJSON()

		if err != nil {
			fmt.Printf("Error raised: %s", err.Error())
			os.Exit(1)
		}

		err = db.Insert(clusterName, output)

		if err != nil {
			fmt.Printf("Error raised: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println("")
	},
}

func init() {
	clusterCmd.AddCommand(listCmd)
	clusterCmd.AddCommand(runCmd)
	clusterCmd.AddCommand(destroyCmd)
	clusterCmd.AddCommand(showCmd)
	clusterCmd.AddCommand(addCmd)
	rootCmd.AddCommand(clusterCmd)
}
