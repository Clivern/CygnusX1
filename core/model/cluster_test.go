// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitClusterModel
func TestUnitClusterModel(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Cluster", func() {
		g.It("It should satisfy test cases", func() {
			cluster1 := &Cluster{
				ID:      "123",
				Name:    "abc",
				Type:    "local",
				Port:    "9000",
				Address: "localhost:9000",
			}

			out, err := cluster1.ConvertToJSON()

			g.Assert(out).Equal(`{"id":"123","name":"abc","type":"local","port":"9000","address":"localhost:9000"}`)
			g.Assert(err).Equal(nil)

			cluster2 := &Cluster{}

			err = cluster2.LoadFromJSON([]byte(out))

			g.Assert(err).Equal(nil)
			g.Assert(cluster2.ID).Equal("123")
			g.Assert(cluster2.Name).Equal("abc")
			g.Assert(cluster2.Type).Equal("local")
			g.Assert(cluster2.Port).Equal("9000")
			g.Assert(cluster2.Address).Equal("localhost:9000")
		})
	})
}
