// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"strings"

	"github.com/clivern/peacock/cmd"

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

	log.SetOutput(os.Stdout)

	if os.Getenv("PC_LOG_LEVEL") == "INFO" {
		log.SetLevel(log.InfoLevel)
	} else if os.Getenv("PC_LOG_LEVEL") == "WARN" {
		log.SetLevel(log.WarnLevel)
	} else if os.Getenv("PC_LOG_LEVEL") == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	} else if os.Getenv("PC_LOG_LEVEL") == "TRACE" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	log.SetFormatter(&log.JSONFormatter{})

	cmd.Execute()
}
