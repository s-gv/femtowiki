// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"encoding/hex"
	"time"
	"github.com/s-gv/femtowiki/models/db"
	"errors"
)

func CreateSuperUser(username string, passwd string) error {
	return CreateUser(username, passwd, "", false)
}

func CreateUser(username string, passwd string, email string, isSuperUser bool) error {
	if passwdHash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost); err == nil {
		r := db.QueryRow(`SELECT username FROM users WHERE username=?;`, username)
		var tmp string
		if err := r.Scan(&tmp); err == sql.ErrNoRows {
			db.Exec(`INSERT INTO users(username, passwdhash, email, is_superuser, created_date, updated_date) VALUES(?, ?, ?, ?, ?, ?);`,
				username, hex.EncodeToString(passwdHash), email, isSuperUser, time.Now().Unix(), time.Now().Unix())
		} else {
			return errors.New("Username already exists.")
		}
	} else {
		return err
	}
	return nil
}

func UpdateUserPasswd(username string, passwd string) error {
	return nil
}