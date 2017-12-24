// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import "github.com/s-gv/femtowiki/models/db"

const (
	Version         = "version"
	CRUDGroup       = "crud_group"
	FileMasterGroup = "file_master_group"
	ConfigJSON      = "config_json"
	HeaderLinks     = "header_links"
	FooterLinks     = "footer_links"
	NavSections     = "nav_sections"
	IllegalNames    = "illegal_names"
)

const (
	DefaultConfigJSON = `{
	"WikiName": "Femtowiki",
	"SignupDisabled": true,
	"DataDir": "",
	"SMTPHost": "",
	"SMTPPort": "",
	"SMTPUser": "",
	"SMTPPasswd": "",
	"FromEmail": ""
}`
	DefaultCRUDGroup = EverybodyGroup
	DefaultFileMasterGroup = EverybodyGroup
	DefaultHeaderLinks = `[
	{"Title": "Home", "URL": "/"},
	{"Title": "Download", "URL": "http://www.goodoldweb.com/"}
]`
	DefaultFooterLinks = `[
	{"Title": "Femtowiki", "URL": "http://www.goodoldweb.com/"},
	{"Title": "Privacy Policy", "URL": "/pages/Privacy_Policy"},
	{"Title": "Terms Of Service", "URL": "/pages/Terms_Of_Service"}
]`
	DefaultNavSections = `[
	{"Title": "", "Links": [{"Title": "Home", "URL": "/"}, {"Title": "Help", "URL": "/pages/Help"}]},
	{"Title": "Cities", "Links": [{"Title": "Bangalore", "URL": "/pages/Bangalore"}, {"Title": "London", "URL": "/pages/London"}, {"Title": "New York", "URL": "/pages/New_York"}]}
]`
	DefaultIllegalNames = `["shit", "crap"]`
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