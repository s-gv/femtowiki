// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/s-gv/femtowiki/templates"
	"strings"
	"github.com/s-gv/femtowiki/models"
	"github.com/s-gv/femtowiki/models/db"
	"database/sql"
	"time"
	"html/template"
	"regexp"
	"encoding/base64"
)

var headerRe *regexp.Regexp

func init() {
	headerRe = regexp.MustCompile("<h[2-9]>.+?</h[2-9]>")
}

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
	unsafe := blackfriday.Run([]byte(strings.Replace(content, "\r\n", "\n", -1)))
	html := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	type Link struct {
		name string
		URL string
	}
	type TOCItem struct {
		title Link
		subtitles []Link
	}
	var toc []TOCItem
	html = headerRe.ReplaceAllStringFunc(html, func(h string) string {
		n := string(h[2])
		if n == "2" || n == "3" {
			if len(h) >= 10 {
				t := h[4:len(h)-5]
				id := "m-"+n+"-"+base64.StdEncoding.EncodeToString([]byte(t))
				link := Link{t, "#"+id}
				if n == "2" {
					var tocItem TOCItem
					tocItem.title = link
					toc = append(toc, tocItem)
				}
				if n == "3" {
					if len(toc) > 0 {
						toc[len(toc)-1].subtitles = append(toc[len(toc)-1].subtitles, link)
					}
				}
				return "<h"+n+" id=\""+id+"\">" + t + "</h"+n+">"
			}
		}
		return h
	})
	tocHTML := "<div class=\"toc\"><ol>\n"
	for _, t := range toc {
		tocHTML += "<li><a href=\""+t.title.URL+"\">"+t.title.name+"</a>"
		if len(t.subtitles) > 0 {
			tocHTML += "<ul>"
			for _, s:= range t.subtitles {
				tocHTML += "<li><a href=\""+s.URL+"\">"+s.name+"</a>"
			}
			tocHTML += "</ul>"
		}
		tocHTML += "</li>"
	}
	tocHTML += "\n</ol></div>"
	html = strings.Replace(html, "<p><strong>TOC</strong></p>", tocHTML, 1)
	templates.Render(w, "index.html", map[string]interface{}{
		"ctx": ctx,
		"Title": title,
		"cTitle": cTitle,
		"URL": "/pages/"+cTitle,
		"EditURL": "/editpage?t="+cTitle,
		"Content": template.HTML(html),
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
	if len(title) < 2 || len(title) > 200 {
		ctx.SetFlashMsg("Title should have 2-200 characters")
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
	content := "# "+title
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
	isCRUDGroupMember := (models.IsUserInCRUDGroup(ctx.UserName) == nil)
	if !isCRUDGroupMember {
		templates.Render(w, "accessdenied.html", map[string]interface{}{
			"ctx": ctx,
		})
		return
	}
	cTitle := r.FormValue("t")
	title := strings.Replace(cTitle, "_", " ", -1)
	if r.Method == "POST" {
		action := r.PostFormValue("action")
		if action == "Update" {
			content := r.PostFormValue("content")
			db.Exec(`UPDATE pages SET content=? WHERE title=?;`, content, title)
		}
		if action == "Delete" {
			db.Exec(`DELETE FROM pages WHERE title=?;`, title)
			http.Redirect(w, r, "/pages/", http.StatusSeeOther)
			return
		}
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
		"IsCRUDGroupMember": isCRUDGroupMember,
		"URL": "/pages/"+cTitle,
		"EditURL": "/editpage?t="+cTitle,
		"Title": title,
		"cTitle": cTitle,
		"Content": content,
	})
})