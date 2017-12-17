// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"github.com/s-gv/femtowiki/models"
)

var AdminHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin {
		ErrForbiddenHandler(w, r)
		return
	}
	templates.Render(w, "admin.html", map[string]interface{}{
		"ctx": ctx,
		"config": models.ReadConfig(models.ConfigJSON),
		"DefaultConfig": models.DefaultConfigJSON,
	})
})

var AdminConfigUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin {
		ErrForbiddenHandler(w, r)
		return
	}
	models.WriteConfig(models.ConfigJSON, r.PostFormValue("config"))
	ctx.SetFlashMsg("Config updated successfully")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
})
