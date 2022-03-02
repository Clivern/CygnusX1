// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitShell
func TestUnitShell(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Shell", func() {
		g.It("It should satisfy test cases", func() {
			out1, out2, err := Exec("echo hello")

			g.Assert(out1).Equal("hello\n")
			g.Assert(out2).Equal("")
			g.Assert(err).Equal(nil)
		})
	})
}
