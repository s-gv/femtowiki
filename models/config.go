// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"strconv"
	"github.com/s-gv/femtowiki/models/db"
)

func Version() int {
	row := db.QueryRow(`SELECT val FROM configs WHERE name=?;`, "version")
	sval := "0"
	if err := row.Scan(&sval); err == nil {
		if ival, err := strconv.Atoi(sval); err == nil {
			return ival
		}
	}
	return 0
}