// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"github.com/s-gv/femtowiki/models"
	"time"
	"github.com/s-gv/femtowiki/models/db"
	"database/sql"
)

var AdminHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin {
		ErrForbiddenHandler(w, r)
		return
	}
	templates.Render(w, "admin.html", map[string]interface{}{
		"ctx": ctx,
		"PageMasterGroup": models.ReadPageMasterGroup(),
		"FileMasterGroup": models.ReadFileMasterGroup(),
		"config": models.ReadConfig(models.ConfigJSON),
		"header": models.ReadConfig(models.HeaderLinks),
		"footer": models.ReadConfig(models.FooterLinks),
		"nav": models.ReadConfig(models.NavSections),
		"illegal_names": models.ReadConfig(models.IllegalNames),
		"DefaultConfig": models.DefaultConfigJSON,
		"DefaultHeader": models.DefaultHeaderLinks,
		"DefaultFooter": models.DefaultFooterLinks,
		"DefaultNav": models.DefaultNavSections,
		"DefaultIllegalNames": models.DefaultIllegalNames,
	})
})

var AdminPageMasterGroupHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	group := r.PostFormValue("group")
	var tmp string
	if group == models.EverybodyGroup || db.QueryRow(`SELECT id FROM groups WHERE name=?;`, group).Scan(&tmp) == nil {
		models.WriteConfig(models.PageMasterGroup, group)
	} else {
		ctx.SetFlashMsg("Group '"+ group +"' not found")
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})

var AdminFileMasterGroupHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	group := r.PostFormValue("group")
	var tmp string
	if group == models.EverybodyGroup || db.QueryRow(`SELECT id FROM groups WHERE name=?;`, group).Scan(&tmp) == nil {
		models.WriteConfig(models.FileMasterGroup, group)
	} else {
		ctx.SetFlashMsg("Group '"+ group +"' not found")
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})

var AdminConfigUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	models.WriteConfig(models.ConfigJSON, r.PostFormValue("config"))
	ctxCacheDate = time.Unix(0, 0)
	ctx.SetFlashMsg("Config updated")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})

var AdminHeaderUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	models.WriteConfig(models.HeaderLinks, r.PostFormValue("header"))
	ctxCacheDate = time.Unix(0, 0)
	ctx.SetFlashMsg("Header updated")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})

var AdminFooterUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	models.WriteConfig(models.FooterLinks, r.PostFormValue("footer"))
	ctxCacheDate = time.Unix(0, 0)
	ctx.SetFlashMsg("Footer updated")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})

var AdminNavUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	models.WriteConfig(models.NavSections, r.PostFormValue("nav"))
	ctxCacheDate = time.Unix(0, 0)
	ctx.SetFlashMsg("Nav updated")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})

var AdminIllegalNamesUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	models.WriteConfig(models.IllegalNames, r.PostFormValue("illegal_names"))
	ctxCacheDate = time.Unix(0, 0)
	ctx.SetFlashMsg("Illegal name list updated")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})

var AdminUserHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin {
		ErrForbiddenHandler(w, r)
		return
	}

	var users []string
	rows := db.Query(`SELECT username FROM users;`)
	for rows.Next() {
		var user string
		rows.Scan(&user)
		users = append(users, user)
	}
	templates.Render(w, "adminusers.html", map[string]interface{}{
		"ctx": ctx,
		"users": users,
	})
})

var AdminGroupHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin {
		ErrForbiddenHandler(w, r)
		return
	}

	var groups []string
	rows := db.Query(`SELECT name FROM groups;`)
	for rows.Next() {
		var group string
		rows.Scan(&group)
		groups = append(groups, group)
	}
	templates.Render(w, "admingroups.html", map[string]interface{}{
		"ctx": ctx,
		"groups": groups,
	})
})

var AdminGroupCreateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	groupname := r.PostFormValue("groupname")
	if groupname != models.EverybodyGroup {
		if err := models.ValidateName(groupname); err == nil {
			var tmp string
			if db.QueryRow(`SELECT id FROM groups WHERE name=?;`, groupname).Scan(&tmp) == sql.ErrNoRows {
				tNow := time.Now().Unix()
				db.Exec(`INSERT INTO groups(name, created_date, updated_date) VALUES(?, ?, ?);`, groupname, tNow, tNow)
			} else {
				ctx.SetFlashMsg("Group '" + groupname + "' already exits")
			}
		} else {
			ctx.SetFlashMsg(err.Error())
		}
	} else {
		ctx.SetFlashMsg("Group name '"+groupname+"' is reserved")
	}

	http.Redirect(w, r, "/admin/groups", http.StatusSeeOther)
})

var AdminGroupDeleteHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	groupname := r.PostFormValue("groupname")
	db.Exec(`DELETE FROM groups WHERE name=?;`, groupname)
	ctx.SetFlashMsg("Group deleted")
	http.Redirect(w, r, "/admin/groups", http.StatusSeeOther)
})

var AdminGroupMembersHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin {
		ErrForbiddenHandler(w, r)
		return
	}
	groupname := r.FormValue("g")
	var members []string
	rows := db.Query(`SELECT users.username FROM groupmembers INNER JOIN users ON groupmembers.userid=users.id INNER JOIN groups ON groupmembers.groupid=groups.id AND groups.name=?;`, groupname)
	for rows.Next() {
		var member string
		rows.Scan(&member)
		members = append(members, member)
	}
	templates.Render(w, "admingroupmembers.html", map[string]interface{}{
		"ctx": ctx,
		"groupname": groupname,
		"members": members,
	})
})

var AdminGroupMemberCreateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	groupname := r.PostFormValue("groupname")
	username := r.PostFormValue("username")

	var groupID string
	if db.QueryRow(`SELECT id FROM groups WHERE name=?;`, groupname).Scan(&groupID) == nil {
		var userID string
		if db.QueryRow(`SELECT id FROM users WHERE username=?;`, username).Scan(&userID) == nil {
			var tmp string
			if db.QueryRow(`SELECT id FROM groupmembers WHERE userid=? AND groupid=?;`, userID, groupID).Scan(&tmp) == sql.ErrNoRows {
				tNow := time.Now().Unix()
				db.Exec(`INSERT INTO groupmembers(groupid, userid, created_date) VALUES(?, ?, ?);`, groupID, userID, tNow)
				ctx.SetFlashMsg("Added user '"+username+"'")
			} else {
				ctx.SetFlashMsg("User '"+username+"' already in this group")
			}
		} else {
			ctx.SetFlashMsg("User '"+username+"' not found")
		}
	}
	http.Redirect(w, r, "/admin/groupmembers?g="+groupname, http.StatusSeeOther)
})

var AdminGroupMemberDeleteHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin || r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	groupname := r.PostFormValue("groupname")
	username := r.PostFormValue("username")

	var userID string
	if db.QueryRow(`SELECT id FROM users WHERE username=?;`, username).Scan(&userID) == nil {
		db.Exec(`DELETE FROM groupmembers WHERE userid=?;`, userID)
		ctx.SetFlashMsg("User '"+username+"' removed from this group")
	} else {
		ctx.SetFlashMsg("User '"+username+"' not in this group")
	}
	http.Redirect(w, r, "/admin/groupmembers?g="+groupname, http.StatusSeeOther)
})