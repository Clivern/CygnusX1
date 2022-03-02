// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// TestUnitKafka test cases
func TestUnitKafka(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestKafka", func() {
		g.It("It should satisfy all provided test cases", func() {
			kafka := GetKafkaConfig("latest", "latest", KafkaPort)
			result, err := kafka.ToString()

			g.Assert(strings.Contains(result, "restart: unless-stopped")).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
