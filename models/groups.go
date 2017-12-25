// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"database/sql"
	"github.com/s-gv/femtowiki/models/db"
)

const (
	EverybodyGroup = "everybody"
)

func ReadPageMasterGroup() string {
	pageMasterGroup := ReadConfig(PageMasterGroup)
	var tmp string
	if db.QueryRow(`SELECT id FROM groups WHERE name=?;`, pageMasterGroup).Scan(&tmp) == sql.ErrNoRows {
		pageMasterGroup = DefaultPageMasterGroup
	}
	return pageMasterGroup
}

func ReadFileMasterGroup() string {
	fileMasterGroup := ReadConfig(FileMasterGroup)
	var tmp string
	if db.QueryRow(`SELECT id FROM groups WHERE name=?;`, fileMasterGroup).Scan(&tmp) == sql.ErrNoRows {
		fileMasterGroup = DefaultFileMasterGroup
	}
	return fileMasterGroup
}

func IsUserInPageMasterGroup(username string) bool {
	if username != "" {
		pageMasterGroup := ReadPageMasterGroup()
		if pageMasterGroup == EverybodyGroup {
			return true
		}
		row := db.QueryRow(`SELECT groupmembers.id FROM groupmembers INNER JOIN groups ON groups.id=groupmembers.groupid AND groups.name=? INNER JOIN users ON users.id=groupmembers.userid AND users.username=?;`, pageMasterGroup, username)
		var tmp string
		if row.Scan(&tmp) == nil {
			return true
		}
	}
	return false
}

func IsUserInFileMasterGroup(username string) bool {
	if username != "" {
		fileMasterGroup := ReadFileMasterGroup()
		if fileMasterGroup == EverybodyGroup {
			return true
		}
		row := db.QueryRow(`SELECT groupmembers.id FROM groupmembers INNER JOIN groups ON groups.id=groupmembers.groupid AND groups.name=? INNER JOIN users ON users.id=groupmembers.userid AND users.username=?;`, fileMasterGroup, username)
		var tmp string
		if row.Scan(&tmp) == nil {
			return true
		}
	}
	return false
}
