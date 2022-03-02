// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// KafkaService const
	KafkaService = "kafka"

	// KafkaPort const
	KafkaPort = "9092"

	// KafkaDockerImage const
	KafkaDockerImage = "confluentinc/cp-kafka"

	// KafkaDockerVersion const
	KafkaDockerVersion = "latest"

	// KafkaRestartPolicy const
	KafkaRestartPolicy = "unless-stopped"

	// ZookeeperDockerImage const
	ZookeeperDockerImage = "confluentinc/cp-zookeeper"

	// ZookeeperDockerVersion const
	ZookeeperDockerVersion = "latest"

	// ZookeeperRestartPolicy const
	ZookeeperRestartPolicy = "unless-stopped"
)

// GetKafkaConfig gets yaml definition object
func GetKafkaConfig(zookeeperVersion, kafkaVersion, port string) DockerComposeConfig {
	services := make(map[string]Service)

	if zookeeperVersion == "" {
		zookeeperVersion = ZookeeperDockerVersion
	}

	if kafkaVersion == "" {
		kafkaVersion = KafkaDockerVersion
	}

	services["zookeeper"] = Service{
		Image:   fmt.Sprintf("%s:%s", ZookeeperDockerImage, zookeeperVersion),
		Restart: ZookeeperRestartPolicy,
		Environment: []string{
			"ZOOKEEPER_CLIENT_PORT=2181",
			"ZOOKEEPER_TICK_TIME=2000",
		},
	}

	services["kafka"] = Service{
		Image:   fmt.Sprintf("%s:%s", KafkaDockerImage, kafkaVersion),
		Restart: KafkaRestartPolicy,
		Ports:   []string{fmt.Sprintf("%s:%s", port, port)},
		Environment: []string{
			"KAFKA_BROKER_ID=1",
			"KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181",
			fmt.Sprintf("KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:%s", port),
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT",
			"KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
		},
		DependsOn: []string{
			"zookeeper",
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
