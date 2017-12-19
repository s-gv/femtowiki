// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"database/sql"
	"github.com/s-gv/femtowiki/models/db"
)

var (
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