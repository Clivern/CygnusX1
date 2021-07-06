// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/clivern/nitro/cmd"
	"github.com/clivern/nitro/core/driver"

	log "github.com/sirupsen/logrus"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit
	cmd.Date = date
	cmd.BuiltBy = builtBy
	cmd.HOME = strings.TrimSpace(os.Getenv("HOME"))
	level := strings.ToLower(os.Getenv("PC_LOG_LEVEL"))

	log.SetOutput(os.Stdout)

	if level == "info" {
		log.SetLevel(log.InfoLevel)
	} else if level == "warn" {
		log.SetLevel(log.WarnLevel)
	} else if level == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if level == "trace" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	if cmd.HOME == "" {
		fmt.Println("Error! `HOME` environment variable is not set")
		os.Exit(1)
	}

	log.SetFormatter(&log.JSONFormatter{})

	db := driver.NewDatabase(fmt.Sprintf(
		"%s/.nitro/nitro.db",
		cmd.HOME,
	))

	err := db.Migrate()

	if err != nil {
		fmt.Printf("Error while migrating database: %s", err.Error())
		os.Exit(1)
	}

	cmd.Execute()
}
