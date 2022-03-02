// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"fmt"
	"strings"

	"github.com/clivern/peacock/core/definition"
	"github.com/clivern/peacock/core/service"
	"github.com/clivern/peacock/core/util"

	log "github.com/sirupsen/logrus"
)

// DockerCompose type
type DockerCompose struct {
	StoragePath string
	fs          *service.FileSystem
}

// NewDockerCompose creates a new instance
func NewDockerCompose(storagePath string) *DockerCompose {
	return &DockerCompose{
		StoragePath: storagePath,
		fs:          service.NewFileSystem(),
	}
}

// Deploy deploys services
func (d *DockerCompose) Deploy(serviceID string, config definition.DockerComposeConfig) error {
	result, err := config.ToString()

	if err != nil {
		return err
	}

	err = d.fs.StoreFile(
		fmt.Sprintf("%s/%s/docker-compose.yml", d.StoragePath, serviceID),
		result,
	)

	if err != nil {
		return err
	}

	command := fmt.Sprintf(
		"docker-compose -f %s/%s/docker-compose.yml -p %s up -d --force-recreate",
		d.StoragePath,
		serviceID,
		serviceID,
	)

	_, _, err = util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return err
	}

	return nil
}

// Destroy destroys services
func (d *DockerCompose) Destroy(serviceID string, config definition.DockerComposeConfig) error {
	result, err := config.ToString()

	if err != nil {
		return err
	}

	err = d.fs.StoreFile(
		fmt.Sprintf("%s/%s/docker-compose.yml", d.StoragePath, serviceID),
		result,
	)

	if err != nil {
		return err
	}

	command := fmt.Sprintf(
		"docker-compose -f %s/%s/docker-compose.yml -p %s down -v --remove-orphans",
		d.StoragePath,
		serviceID,
		serviceID,
	)

	_, _, err = util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return err
	}

	return nil
}

// FetchServicePort destroys services
func (d *DockerCompose) FetchServicePort(serviceID, serviceName, servicePort string, config definition.DockerComposeConfig) (string, error) {
	result, err := config.ToString()

	if err != nil {
		return "", err
	}

	err = d.fs.StoreFile(
		fmt.Sprintf("%s/%s/docker-compose.yml", d.StoragePath, serviceID),
		result,
	)

	if err != nil {
		return "", err
	}

	command := fmt.Sprintf(
		"docker-compose -f %s/%s/docker-compose.yml -p %s port %s %s",
		d.StoragePath,
		serviceID,
		serviceID,
		serviceName,
		servicePort,
	)

	stdout, _, err := util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command")

	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(strings.Replace(stdout, "0.0.0.0:", "", -1), "\n"), nil
}

// Prune prune docker system
func (d *DockerCompose) Prune() error {
	command := "docker system prune -a -f --volumes"

	_, _, err := util.Exec(command)

	log.WithFields(log.Fields{
		"command": command,
	}).Info("Run a shell command to prune docker system")

	if err != nil {
		return err
	}

	return nil
}
