// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitCorrelationService
func TestUnitCorrelationService(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Correlation", func() {
		g.It("It should satisfy test cases", func() {
			g.Assert(NewCorrelation().UUIDv4() != "").Equal(true)
		})
	})
}
