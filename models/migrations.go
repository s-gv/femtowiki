// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"github.com/s-gv/femtowiki/models/db"
	"log"
	"time"
)

const ModelVersion = 1

func Migration1() {
	db.Exec(`CREATE TABLE configs(name VARCHAR(250), val TEXT);`)
	db.Exec(`CREATE UNIQUE INDEX configs_key_index on configs(name);`)

	db.Exec(`CREATE TABLE users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(32) NOT NULL,
		passwdhash VARCHAR(250) NOT NULL,
		email VARCHAR(250) DEFAULT '',
		reset_token VARCHAR(250) DEFAULT '',
		reset_token_date INTEGER DEFAULT 0,
		is_banned INTEGER DEFAULT 0,
		is_superuser INTEGER DEFAULT 0,
		created_date INTEGER DEFAULT 0,
		updated_date INTEGER DEFAULT 0
	);`)
	db.Exec(`CREATE UNIQUE INDEX users_username_index on users(username);`)
	db.Exec(`CREATE INDEX users_email_index on users(email);`)
	db.Exec(`CREATE INDEX users_reset_token_index on users(reset_token);`)
	db.Exec(`CREATE INDEX users_created_index on users(created_date);`)

	db.Exec(`CREATE TABLE groups(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(250) DEFAULT '',
		created_date INTEGER DEFAULT 0,
		updated_date INTEGER DEFAULT 0
	);`)
	db.Exec(`CREATE UNIQUE INDEX groups_name_index on groups(name);`)

	db.Exec(`CREATE TABLE groupmembers(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userid INTEGER REFERENCES users(id) ON DELETE CASCADE,
		groupid INTEGER REFERENCES groups(id) ON DELETE CASCADE,
		created_date INTEGER DEFAULT 0
	);`)
	db.Exec(`CREATE INDEX groupmembers_userid_index on groupmembers(userid);`)
	db.Exec(`CREATE INDEX groupmembers_groupid_index on groupmembers(groupid);`)

	db.Exec(`CREATE TABLE sessions(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sessionid VARCHAR(250) DEFAULT '',
		csrftoken VARCHAR(250) DEFAULT '',
		userid INTEGER REFERENCES users(id) ON DELETE CASCADE,
		msg VARCHAR(250) DEFAULT '',
		created_date INTEGER DEFAULT 0,
		updated_date INTEGER DEFAULT 0
	);`)
	db.Exec(`CREATE INDEX sessions_sessionid_index on sessions(sessionid);`)
	db.Exec(`CREATE INDEX sessions_userid_index on sessions(userid);`)

	db.Exec(`CREATE TABLE pages(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(250) DEFAULT '',
		content TEXT DEFAULT '',
		editgroupid INTEGER REFERENCES groups(id) ON DELETE SET NULL,
		readgroupid INTEGER REFERENCES groups(id) ON DELETE SET NULL,
		created_date INTEGER DEFAULT 0,
		updated_date INTEGER DEFAULT 0
	);`)
	db.Exec(`CREATE INDEX pages_title_index on pages(title);`)
	db.Exec(`CREATE INDEX pages_editgroupid_index on pages(editgroupid);`)
	db.Exec(`CREATE INDEX pages_readgroupid_index on pages(readgroupid);`)

	db.Exec(`CREATE TABLE uploads(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(250) DEFAULT '',
		location VARCHAR(250) DEFAULT '',
		editgroupid INTEGER REFERENCES groups(id) ON DELETE SET NULL,
		readgroupid INTEGER REFERENCES groups(id) ON DELETE SET NULL,
		created_date INTEGER DEFAULT 0,
		updated_date INTEGER DEFAULT 0
	);`)
	db.Exec(`CREATE INDEX uploads_name_index on uploads(name);`)
	db.Exec(`CREATE INDEX uploads_editgroupid_index on uploads(editgroupid);`)
	db.Exec(`CREATE INDEX uploads_readgroupid_index on uploads(readgroupid);`)
	// db.Exec(`CREATE VIRTUAL TABLE email USING fts5(sender, title, body);`)
}

func IsMigrationNeeded() bool {
	return db.Version() != ModelVersion
}

func Migrate() {
	dbver := db.Version()
	if dbver == ModelVersion {
		log.Panicf("[ERROR] DB migration not needed. DB up-to-date.\n")
	} else if dbver > ModelVersion {
		log.Panicf("[ERROR] DB version (%d) is greater than binary version (%d). Use newer binary.\n", dbver, ModelVersion)
	}
	for dbver < ModelVersion {
		if dbver == 0 {
			log.Printf("[INFO] Migrating to version 1...")
			Migration1()

			WriteConfig(Version, "1")
			WriteConfig(ConfigJSON, DefaultConfigJSON)
			WriteConfig(HeaderLinks, DefaultHeaderLinks)
			WriteConfig(FooterLinks, DefaultFooterLinks)
			WriteConfig(NavSections, DefaultNavSections)
			WriteConfig(IllegalNames, DefaultIllegalNames)
			WriteConfig(CRUDGroup, DefaultCRUDGroup)
			db.Exec(`INSERT INTO pages(title, content, created_date, updated_date) VALUES(?, ?, ?, ?);`, "Home Page", "# Home Page", time.Now().Unix(), time.Now().Unix())
		}
		dbver = db.Version()
	}
}