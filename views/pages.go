// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"strings"
	"github.com/s-gv/femtowiki/models"
)

var PagesHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	pageTitle := strings.Replace(r.URL.Path[7:], "_", " ", -1) // r.URL will be /pages/<page_title>
	if pageTitle == "" {
		// List all pages
		if !ctx.IsAdmin && models.IsUserInCRUDGroup(ctx.UserName) != nil {
			templates.Render(w, "accessdenied.html", map[string]interface{}{
				"ctx": ctx,
			})
			return
		}
		templates.Render(w, "pagelist.html", map[string]interface{}{
			"ctx": ctx,
		})
		return
	}
	// Render the relevant wiki page
	templates.Render(w, "index.html", map[string]interface{}{
		"ctx": ctx,
		"title": pageTitle,
		"content": "",
	})
})