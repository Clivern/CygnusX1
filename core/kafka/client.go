// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package kafka

import (
	"errors"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

// Client type
type Client struct {
}

// DefaultConsumerSaramaConfig function creates a sarama configuration with
// a client ID derived from host name and consumer name.
func DefaultConsumerSaramaConfig(name string) (*sarama.Config, error) {

	host, err := os.Hostname()

	if err != nil {
		return nil, errors.New("failed to get hostname")
	}

	config := sarama.NewConfig()
	config.ClientID = fmt.Sprintf("%s-%s", host, name)
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_0

	return config, nil
}
