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
	"github.com/clivern/peacock/core/model"
	"github.com/clivern/peacock/core/module"
	"github.com/clivern/peacock/core/runtime"

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
			"%s/.peacock/peacock.db",
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

// destroyCmd destroy clutser sub-command
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
			fmt.Println("Error! cluster not found!")
			os.Exit(1)
		}

		cluster := &model.Cluster{}

		err := cluster.LoadFromJSON([]byte(val))

		if err != nil {
			fmt.Printf("Error raised: %s", err.Error())
			os.Exit(1)
		}

		dc := runtime.NewDockerCompose(fmt.Sprintf(
			"%s/.peacock/cache",
			HOME,
		))

		config := definition.GetKafkaConfig(
			definition.ZookeeperDockerVersion,
			definition.KafkaDockerVersion,
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

		port, err := dc.FetchServicePort(
			id,
			definition.KafkaService,
			definition.KafkaPort,
			config,
		)

		if err != nil {
			fmt.Println(fmt.Sprintf("Error! unable to fetch the port: %s", err.Error()))
			os.Exit(1)
		}

		cluster := &model.Cluster{
			ID:   id,
			Type: "local",
			Port: port,
			Name: clusterName,
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
	rootCmd.AddCommand(clusterCmd)
}
