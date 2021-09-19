// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"database/sql"
	"fmt"

	"github.com/clivern/peacock/core/model"
	"github.com/clivern/peacock/core/service"

	_ "modernc.org/sqlite"
)

// Database type
type Database struct {
	Datafile string
}

// NewDatabase gets a new database object
func NewDatabase(datafile string) *Database {
	return &Database{Datafile: datafile}
}

// Migrate creates the database
func (d *Database) Migrate() error {
	fs := service.NewFileSystem()

	if fs.FileExists(d.Datafile) {
		return nil
	}

	db, err := sql.Open("sqlite", d.Datafile)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(`DROP TABLE IF EXISTS cluster; CREATE TABLE IF NOT EXISTS cluster (key text unique, value text)`)

	if err != nil {
		return err
	}

	return nil
}

// Insert adds a record to database
func (d *Database) Insert(key, value string) error {
	db, err := sql.Open("sqlite", d.Datafile)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(fmt.Sprintf(`INSERT INTO cluster values('%s', '%s')`, key, value))

	if err != nil {
		return err
	}

	return nil
}

// FindByKey gets the value by key
func (d *Database) FindByKey(key string) (*model.Cluster, error) {
	db, err := sql.Open("sqlite", d.Datafile)

	cluster := &model.Cluster{}

	if err != nil {
		return cluster, err
	}

	defer db.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT value FROM cluster WHERE key = '%s'`, key))

	if err != nil {
		return cluster, err
	}

	defer rows.Close()

	for rows.Next() {
		var value string

		err = rows.Scan(&value)

		if err != nil {
			return cluster, err
		}

		err = cluster.LoadFromJSON([]byte(value))

		if err != nil {
			return cluster, err
		}

		return cluster, nil
	}

	return cluster, fmt.Errorf("Unable to find value with key %s", key)
}

// DeleteByKey deletes a record by a key
func (d *Database) DeleteByKey(key string) error {
	db, err := sql.Open("sqlite", d.Datafile)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(fmt.Sprintf(`DELETE FROM cluster WHERE key = '%s'`, key))

	if err != nil {
		return err
	}

	return nil
}

// FindAll get all records
func (d *Database) FindAll() ([]model.Cluster, error) {

	clusters := []model.Cluster{}

	db, err := sql.Open("sqlite", d.Datafile)

	if err != nil {
		return clusters, err
	}

	defer db.Close()

	rows, err := db.Query(`SELECT key, value FROM cluster`)

	if err != nil {
		return clusters, err
	}

	defer rows.Close()

	for rows.Next() {
		var key string
		var value string

		err = rows.Scan(&key, &value)

		if err != nil {
			return clusters, err
		}

		cluster := &model.Cluster{}

		err := cluster.LoadFromJSON([]byte(value))

		if err != nil {
			continue
		}

		clusters = append(clusters, *cluster)

		return clusters, nil
	}

	return clusters, nil
}
