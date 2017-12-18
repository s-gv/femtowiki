// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"github.com/s-gv/femtowiki/models"
	"time"
)

var AdminHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin {
		ErrForbiddenHandler(w, r)
		return
	}
	templates.Render(w, "admin.html", map[string]interface{}{
		"ctx": ctx,
		"config": models.ReadConfig(models.ConfigJSON),
		"header": models.ReadConfig(models.HeaderLinks),
		"footer": models.ReadConfig(models.FooterLinks),
		"nav": models.ReadConfig(models.NavSections),
		"DefaultConfig": models.DefaultConfigJSON,
		"DefaultHeader": models.DefaultHeaderLinks,
		"DefaultFooter": models.DefaultFooterLinks,
		"DefaultNav": models.DefaultNavSections,
	})
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