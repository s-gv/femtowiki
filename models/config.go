// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import "github.com/s-gv/femtowiki/models/db"

const (
	Version = "version"
	ConfigJSON = "config_json"
)

const (
	DefaultConfigJSON = `{
	"WikiName": "Femtowiki",
	"SignupDisabled": false,
	"SMTPHost": "",
	"SMTPPort": "",
	"SMTPUser": "",
	"SMTPPasswd": "",
	"FromEmail": ""
}`
)

func WriteConfig(key string, val string) {
	var oldVal string
	if db.QueryRow(`SELECT val FROM configs WHERE name=?;`, key).Scan(&oldVal) == nil {
		if oldVal != val {
			db.Exec(`UPDATE configs SET val=? WHERE name=?;`, val, key)
		}
	} else {
		db.Exec(`INSERT INTO configs(name, val) values(?, ?);`, key, val)
	}
}

func ReadConfig(key string) string {
	var val string
	if db.QueryRow(`SELECT val FROM configs WHERE name=?;`, key).Scan(&val) == nil {
		return val
	}
	return ""
}