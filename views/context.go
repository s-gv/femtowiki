// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"time"
	"encoding/hex"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/s-gv/femtowiki/models/db"
	"errors"
	"github.com/s-gv/femtowiki/models"
	"encoding/json"
)

type Context struct {
	SessionID        string
	UserName         string
	IsUserValid      bool
	IsAdmin          bool
	CSRFToken        string
	FlashMsg         string
	IsPageCRUDMember bool
	Config		     WikiConfig
	HeaderLinks      []NavLink
	FooterLinks      []NavLink
	NavSections      []NavSection
}

type WikiConfig struct {
	WikiName		string
	SignupDisabled	bool
	DataDir         string
	SMTPHost		string
	SMTPPort		string
	SMTPUser		string
	SMTPPasswd		string
	FromEmail		string
}

type NavLink struct {
	Title string
	URL   string
}

type NavSection struct {
	Title string
	Links []NavLink
}

var configCache WikiConfig
var headerLinksCache []NavLink
var footerLinksCache []NavLink
var navSectionsCache []NavSection
var ctxCacheDate time.Time


const maxConfigCacheLife = 5*time.Minute
const maxSessionLife = 200*time.Hour
const maxSessionLifeBeforeUpdate = 100*time.Hour

func ReadContext(sessionID string) Context {
	ctx := Context{}
	if sessionID != "" {
		var username, csrftoken, flashmsg string
		var isAdmin bool
		var uDate int64
		r := db.QueryRow(`SELECT users.username, users.is_superuser, sessions.csrftoken, sessions.msg, sessions.updated_date FROM sessions INNER JOIN users ON sessions.userid=users.id WHERE sessions.sessionid=?;`, sessionID)
		if err := r.Scan(&username, &isAdmin, &csrftoken, &flashmsg, &uDate); err == nil {
			updatedDate := time.Unix(uDate, 0)
			if updatedDate.After(time.Now().Add(-maxSessionLife)) {
				if updatedDate.Before(time.Now().Add(-maxSessionLifeBeforeUpdate)) {
					tNow := int64(time.Now().Unix())
					db.Exec(`UPDATE sessions SET updated_date=? WHERE sessionid=?;`, tNow, sessionID)
				}
				db.Exec(`UPDATE sessions SET msg=? WHERE sessionid=?;`, "", sessionID)

				ctx.SessionID = sessionID
				ctx.FlashMsg = flashmsg
				ctx.UserName = username
				ctx.IsAdmin = isAdmin
				ctx.IsUserValid = true
				ctx.CSRFToken = csrftoken

				CRUDGroup := models.ReadCRUDGroup()
				row := db.QueryRow(`SELECT groupmembers.id FROM groupmembers INNER JOIN groups ON groups.id=groupmembers.groupid AND groups.name=? INNER JOIN users ON users.id=groupmembers.userid AND users.username=?;`, CRUDGroup, ctx.UserName)
				var tmp string
				ctx.IsPageCRUDMember = ctx.IsAdmin || (CRUDGroup == models.EverybodyGroup) || (row.Scan(&tmp) == nil)
			} else {
				// Session expired
				//log.Printf("[INFO] Attempted to use expired session (id: %s)\n", sessionID)
			}
		} else {
			// Invalid sessionid
			//log.Printf("[INFO] Attempted to use invalid session (id: %s). Error msg: %s\n", sessionID, err)
		}

		db.Exec(`DELETE FROM sessions WHERE updated_date < ?;`, int64(time.Now().Add(-maxSessionLife).Unix()))
	}

	if ctxCacheDate.Before(time.Now().Add(-maxConfigCacheLife)) {
		configJSON := models.ReadConfig(models.ConfigJSON)
		var config WikiConfig
		if err := json.Unmarshal([]byte(configJSON), &config); err == nil {
			configCache = config
		} else {
			json.Unmarshal([]byte(models.DefaultConfigJSON), &configCache)
			log.Printf("[ERROR] Invalid config: %s\n", configJSON)
		}

		headerJSON := models.ReadConfig(models.HeaderLinks)
		var headerLinks []NavLink
		if err := json.Unmarshal([]byte(headerJSON), &headerLinks); err == nil {
			headerLinksCache = headerLinks
		} else {
			json.Unmarshal([]byte(models.DefaultHeaderLinks), &headerLinksCache)
			log.Printf("[ERROR] Invalid header links: %s\n", headerJSON)
		}

		footerJSON := models.ReadConfig(models.FooterLinks)
		var footerLinks []NavLink
		if err := json.Unmarshal([]byte(footerJSON), &footerLinks); err == nil {
			footerLinksCache = footerLinks
		} else {
			json.Unmarshal([]byte(models.DefaultFooterLinks), &footerLinksCache)
			log.Printf("[ERROR] Invalid footer links: %s\n", footerJSON)
		}

		navJSON := models.ReadConfig(models.NavSections)
		var navSections []NavSection
		if err := json.Unmarshal([]byte(navJSON), &navSections); err == nil {
			navSectionsCache = navSections
		} else {
			json.Unmarshal([]byte(models.DefaultNavSections), &navSectionsCache)
			log.Printf("[ERROR] Invalid nav sections: %s\n", navJSON)
		}

		ctxCacheDate = time.Now()
	}

	ctx.Config = configCache
	ctx.HeaderLinks = headerLinksCache
	ctx.FooterLinks = footerLinksCache
	ctx.NavSections = navSectionsCache

	return ctx
}

func (ctx *Context) ValidateCSRFToken(token string) error {
	if ctx.SessionID != "" && ctx.CSRFToken != "" {
		if ctx.CSRFToken == token {
			return nil
		}
	}
	return errors.New("Invalid CSRF token")
}

func (ctx *Context) Authenticate(username string, passwd string) error {
	r := db.QueryRow(`SELECT id, passwdhash, is_banned, is_superuser FROM users WHERE username=?;`, username)
	var passwdHashStr string
	var userID int
	var isBanned bool
	var isAdmin bool
	if err := r.Scan(&userID, &passwdHashStr, &isBanned, &isAdmin); err != nil {
		return errors.New("Incorrect username or password")
	}
	if isBanned {
		return errors.New("User banned")
	}
	passwdHash, err := hex.DecodeString(passwdHashStr)
	if err != nil {
		log.Panicf("[ERROR] Error in converting password hash from hex to byte slice: %s\n", err)
	}
	if err := bcrypt.CompareHashAndPassword(passwdHash, []byte(passwd)); err != nil {
		return errors.New("Incorrect username or password")
	}

	sessionID := randSeq(48)
	csrfToken := randSeq(48)
	tNow := int64(time.Now().Unix())
	db.Exec(`INSERT INTO sessions(sessionid, userid, csrftoken, created_date, updated_date) VALUES(?, ?, ?, ?, ?);`, sessionID, userID, csrfToken, tNow, tNow)

	ctx.SessionID = sessionID
	ctx.UserName = username
	ctx.IsUserValid = true
	ctx.IsAdmin = isAdmin
	ctx.CSRFToken = csrfToken

	return nil
}

func (ctx *Context) SetFlashMsg(msg string) {
	db.Exec(`UPDATE sessions SET msg=? WHERE sessionid=?;`, msg, ctx.SessionID)
}