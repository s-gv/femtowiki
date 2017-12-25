// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"io"
	"github.com/s-gv/femtowiki/static"
	"github.com/s-gv/femtowiki/models"
)

var SearchHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	searchTerms := r.FormValue("q")
	results := models.PageSearch(searchTerms)
	templates.Render(w, "search.html", map[string]interface{}{
		"ctx": ctx,
		"Results": results,
		"SearchTerms": searchTerms,
	})
})

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	defer ErrServerHandler(w, r)
	ErrNotFoundHandler(w, r)
}

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	defer ErrServerHandler(w, r)
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "max-age=31536000, public")
	io.WriteString(w, static.StyleSrc)
}

func ScriptHandler(w http.ResponseWriter, r *http.Request) {
	defer ErrServerHandler(w, r)
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=31536000, public")
	io.WriteString(w, static.ScriptSrc)
}