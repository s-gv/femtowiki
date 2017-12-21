// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"strings"
	"github.com/s-gv/femtowiki/models"
	"github.com/s-gv/femtowiki/models/db"
	"database/sql"
	"time"
)

var PagesHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	cTitle := r.URL.Path[7:] // r.URL will be /pages/<page_title>
	title := strings.Replace(cTitle, "_", " ", -1)
	if title == "" {
		// List all pages
		if !ctx.IsAdmin && models.IsUserInCRUDGroup(ctx.UserName) != nil {
			templates.Render(w, "accessdenied.html", map[string]interface{}{
				"ctx": ctx,
			})
			return
		}

		type Page struct {
			Title string
			URL   string
		}
		pages := []Page{}
		rows := db.Query(`SELECT title FROM pages ORDER BY title;`)
		for rows.Next() {
			page := Page{}
			rows.Scan(&page.Title)
			page.URL = strings.Replace(page.Title, " ", "_", -1)
			pages = append(pages, page)
		}
		templates.Render(w, "pagelist.html", map[string]interface{}{
			"ctx": ctx,
			"pages": pages,
		})
		return
	}
	// Render the relevant wiki page
	row := db.QueryRow(`SELECT content FROM pages WHERE title=?;`, title)
	var content string
	if row.Scan(&content) != nil {
		ErrNotFoundHandler(w, r)
		return
	}
	templates.Render(w, "index.html", map[string]interface{}{
		"ctx": ctx,
		"title": title,
		"URL": "/pages/"+cTitle,
		"EditURL": "/editpage?t="+cTitle,
		"content": content,
	})
})

var PageCreateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin && models.IsUserInCRUDGroup(ctx.UserName) != nil {
		templates.Render(w, "accessdenied.html", map[string]interface{}{
			"ctx": ctx,
		})
		return
	}
	title := r.PostFormValue("title")
	cTitle := strings.Replace(title, " ", "_", -1)
	var tmp string
	if db.QueryRow(`SELECT id FROM pages WHERE title=?;`, title).Scan(&tmp) != sql.ErrNoRows {
		ctx.SetFlashMsg("Page already exists")
		http.Redirect(w, r, "/pages/#flash", http.StatusSeeOther)
		return
	}
	for _, ch := range title {
		if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') && (ch < '0' && ch > '9') &&(ch != ' ') {
			ctx.SetFlashMsg("Only alphanumeric characters are supported in the title.")
			http.Redirect(w, r, "/pages/#flash", http.StatusSeeOther)
			return
		}
	}
	content := "#"+title
	tNow := time.Now().Unix()
	db.Exec(`INSERT INTO pages(title, content, created_date, updated_date) VALUES(?, ?, ?, ?);`, title, content, tNow, tNow)
	CRUDGroup := models.ReadCRUDGroup()
	if CRUDGroup != models.DefaultCRUDGroup {
		var gID string
		if db.QueryRow(`SELECT id FROM groups WHERE name=?;`, CRUDGroup).Scan(&gID) == nil {
			db.Exec(`UPDATE pages SET editgroupid=? WHERE title=?;`, gID, title)
		}
	}
	http.Redirect(w, r, "/editpage?t="+cTitle, http.StatusSeeOther)
})

var PageEditHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin && models.IsUserInCRUDGroup(ctx.UserName) != nil {
		templates.Render(w, "accessdenied.html", map[string]interface{}{
			"ctx": ctx,
		})
		return
	}
	cTitle := r.FormValue("t")
	title := strings.Replace(cTitle, "_", " ", -1)
	if r.Method == "POST" {
		http.Redirect(w, r, "/pages/"+cTitle, http.StatusSeeOther)
		return
	}
	row := db.QueryRow(`SELECT content FROM pages WHERE title=?;`, title)
	var content string
	if row.Scan(&content) != nil {
		ErrNotFoundHandler(w, r)
		return
	}
	templates.Render(w, "index.html", map[string]interface{}{
		"ctx": ctx,
		"IsEditMode": true,
		"URL": "/pages/"+cTitle,
		"EditURL": "/editpage?t="+cTitle,
		"title": title,
		"content": content,
	})
})

var PageDeleteHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if !ctx.IsAdmin && models.IsUserInCRUDGroup(ctx.UserName) != nil {
		templates.Render(w, "accessdenied.html", map[string]interface{}{
			"ctx": ctx,
		})
		return
	}
	if r.Method != "POST" {
		ErrForbiddenHandler(w, r)
	}
	title := r.FormValue("t")
	db.Exec(`DELETE FROM pages WHERE title=?;`, title)
	http.Redirect(w, r, "/pages/", http.StatusSeeOther)
})