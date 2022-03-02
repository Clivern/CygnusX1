// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitBase test cases
func TestUnitBase(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestBase", func() {
		g.It("It should satisfy all provided test cases", func() {
			g.Assert(GetServiceID() != "").Equal(true)
		})
	})
}
