// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
)

// Cluster type
type Cluster struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Port    string `json:"port"`
	Address string `json:"address"`
}

// LoadFromJSON update object from json
func (c *Cluster) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	return nil
}

// ConvertToJSON convert object to json
func (c *Cluster) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
