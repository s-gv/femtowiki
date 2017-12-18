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
	"log"
	"strings"
	"encoding/json"
)

func CreateSuperUser(username string, passwd string) error {
	return CreateUser(username, passwd, "", true)
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

func ValidateUsername(username string) error {
	if len(username) < 2 || len(username) > 32 {
		return errors.New("Username should have 2-32 characters.")
	}
	for _, ch := range username {
		if (ch < 'A' || ch > 'Z') && (ch < 'a' || ch > 'z') && (ch != '_') && (ch != '-') && (ch < '0' || ch > '9') {
			return errors.New("Username may contain only characters, numbers, underscore, and hyphen.")
		}
	}
	illegalUsernameJSON := ReadConfig(IllegalUsernames)
	var illegalUsernames []string
	if err := json.Unmarshal([]byte(illegalUsernameJSON), &illegalUsernames); err != nil {
		json.Unmarshal([]byte(DefaultIllegalUsernames), &illegalUsernames)
		log.Printf("[ERROR] Invalid illegal usernames: %s\n", illegalUsernameJSON)
	}
	for _, illegalName := range illegalUsernames {
		if strings.Contains(username, illegalName) {
			return errors.New("Illegal username")
		}
	}
	return nil
}

func ValidatePasswd(passwd string) error {
	if len(passwd) < 8 || len(passwd) > 64 {
		return errors.New("Password must have 8-64 characters")
	}
	return nil
}

func VerifyPasswd(username string, passwd string) error {
	r := db.QueryRow(`SELECT passwdhash FROM users WHERE username=?;`, username)
	var passwdHashStr string
	if err := r.Scan(&passwdHashStr); err != nil {
		return errors.New("Incorrect username or password")
	}
	passwdHash, err := hex.DecodeString(passwdHashStr)
	if err != nil {
		log.Panicf("[ERROR] Error in converting password hash from hex to byte slice: %s\n", err)
	}
	if err := bcrypt.CompareHashAndPassword(passwdHash, []byte(passwd)); err != nil {
		return errors.New("Incorrect username or password")
	}
	return nil
}

func UpdateUserPasswd(username string, passwd string) error {
	if passwdHash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost); err == nil {
		r := db.QueryRow(`SELECT username FROM users WHERE username=?;`, username)
		var tmp string
		if err := r.Scan(&tmp); err == nil {
			db.Exec(`UPDATE users SET passwdhash=?, updated_date=? WHERE username=?;`, hex.EncodeToString(passwdHash), time.Now().Unix(), username)
		} else {
			return errors.New("User not found.")
		}
	} else {
		return err
	}
	return nil
}
