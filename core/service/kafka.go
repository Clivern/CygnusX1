// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"github.com/Shopify/sarama"
)

// Kafka type
type Kafka struct {
	Config *sarama.Config
}

// NewKafkaClient creates a kafka client
func NewKafkaClient() *Kafka {
	conf := sarama.NewConfig()

	return &Kafka{
		Config: conf,
	}
}

// Connect connects into kafka cluster
func (k *Kafka) Connect() {}
