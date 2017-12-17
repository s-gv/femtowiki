// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"github.com/s-gv/femtowiki/models/db"
)

var ProfileHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	username := r.FormValue("u")

	if !ctx.IsAdmin && (username != ctx.UserName) {
		ErrForbiddenHandler(w, r)
		return
	}

	row := db.QueryRow(`SELECT email, is_banned FROM users WHERE username=?;`, username)
	var email string
	var isBanned bool
	if err := row.Scan(&email, &isBanned); err != nil {
		ErrNotFoundHandler(w, r)
		return
	}
	if isBanned && !ctx.IsAdmin {
		ErrNotFoundHandler(w, r)
		return
	}

	templates.Render(w, "profile.html", map[string]interface{}{
		"ctx": ctx,
		"username": username,
		"email": email,
	})
})

var ProfileUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	username := r.PostFormValue("username")
	if !ctx.IsAdmin && (username != ctx.UserName) {
		ErrForbiddenHandler(w, r)
		return
	}
	email := r.PostFormValue("email")
	db.Exec(`UPDATE users SET email=? WHERE username=?;`, email, username)
	ctx.SetFlashMsg("Email updated")
	http.Redirect(w, r, "/profile?u="+username, http.StatusSeeOther)
})