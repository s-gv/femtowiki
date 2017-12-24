// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"database/sql"
	"github.com/s-gv/femtowiki/models/db"
	"errors"
)

const (
	EverybodyGroup = "everybody"
)

func ReadCRUDGroup() string {
	CRUDGroup := ReadConfig(CRUDGroup)
	var tmp string
	if db.QueryRow(`SELECT id FROM groups WHERE name=?;`, CRUDGroup).Scan(&tmp) == sql.ErrNoRows {
		CRUDGroup = DefaultCRUDGroup
	}
	return CRUDGroup
}

func IsUserInCRUDGroup(username string) error {
	if username != "" {
		CRUDGroup := ReadCRUDGroup()
		if CRUDGroup == EverybodyGroup {
			return nil
		}
		row := db.QueryRow(`SELECT groupmembers.id FROM groupmembers INNER JOIN groups ON groups.id=groupmembers.groupid AND groups.name=? INNER JOIN users ON users.id=groupmembers.userid AND users.username=?;`, CRUDGroup, username)
		var tmp string
		if row.Scan(&tmp) == nil {
			return nil
		}
	}
	return errors.New("Access denied.")
}