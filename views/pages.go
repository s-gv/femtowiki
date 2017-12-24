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
	isCRUDGroupMember := (models.IsUserInCRUDGroup(ctx.UserName) == nil)
	cTitle := models.IndexPage
	title := strings.Replace(cTitle, "_", " ", -1)
	if r.URL.Path != "/" {
		cTitle = r.URL.Path[7:] // r.URL will be /pages/<page_title>
		title = strings.Replace(cTitle, "_", " ", -1)
		if title == models.IndexPage {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	if title == "" {
		// List all pages
		if !ctx.IsAdmin && !isCRUDGroupMember {
			templates.Render(w, "accessdenied.html", map[string]interface{}{
				"ctx": ctx,
			})
			return
		}

		type Page struct {
			Title     string
			CTitle    string
			ReadGroup string
			EditGroup string
		}
		pages := []Page{}
		rows := db.Query(`SELECT title, readgroupid, editgroupid FROM pages ORDER BY title;`)
		for rows.Next() {
			page := Page{}
			var readGroupID, editGroupID sql.NullString
			rows.Scan(&page.Title, &readGroupID, &editGroupID)
			page.CTitle = strings.Replace(page.Title, " ", "_", -1)
			if readGroupID.Valid {
				db.QueryRow(`SELECT name FROM groups WHERE id=?;`, readGroupID).Scan(&page.ReadGroup)
			}
			if editGroupID.Valid {
				db.QueryRow(`SELECT name FROM groups WHERE id=?;`, editGroupID).Scan(&page.EditGroup)
			}
			pages = append(pages, page)
		}
		templates.Render(w, "pagelist.html", map[string]interface{}{
			"ctx": ctx,
			"pages": pages,
		})
		return
	}
	// Render the relevant wiki page
	row := db.QueryRow(`SELECT readgroupid, content FROM pages WHERE title=?;`, title)
	var content string
	var readGroupID sql.NullString
	if row.Scan(&readGroupID, &content) != nil {
		ErrNotFoundHandler(w, r)
		return
	}
	if !ctx.IsAdmin && !isCRUDGroupMember && readGroupID.Valid {
		row := db.QueryRow(`SELECT groupmembers.id FROM groupmembers INNER JOIN users ON users.id=groupmembers.userid AND users.username=? WHERE groupmembers.groupid=?;`, ctx.UserName, readGroupID)
		var tmp string
		if row.Scan(&tmp) != nil {
			templates.Render(w, "accessdenied.html", map[string]interface{}{
				"ctx": ctx,
			})
			return
		}
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
	if r.URL.Path != "" {
		ctx.PageTitle = title
	}
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
	if r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	isCRUDGroupMember := (models.IsUserInCRUDGroup(ctx.UserName) == nil)
	if !ctx.IsAdmin && !isCRUDGroupMember {
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
	if err := models.IsPageTitleValid(title); err != nil {
		ctx.SetFlashMsg(err.Error())
		http.Redirect(w, r, "/pages/#flash", http.StatusSeeOther)
		return
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
	cTitle := r.FormValue("t")
	title := strings.Replace(cTitle, "_", " ", -1)

	var editGroupID sql.NullString
	db.QueryRow(`SELECT editGroupID FROM pages WHERE title=?;`, title).Scan(&editGroupID)
	if !ctx.IsAdmin && !isCRUDGroupMember && editGroupID.Valid {
		var tmp string
		row := db.QueryRow(`SELECT groupmembers.id FROM groupmembers INNER JOIN users ON users.id=groupmembers.userid AND users.username=? WHERE groupmembers.groupid=?;`, ctx.UserName, editGroupID)
		if row.Scan(&tmp) != nil {
			templates.Render(w, "accessdenied.html", map[string]interface{}{
				"ctx": ctx,
			})
			return
		}
	}
	if r.Method == "POST" {
		action := r.PostFormValue("action")
		if r.PostFormValue("meta") == "" {
			if action == "Update" {
				content := r.PostFormValue("content")
				db.Exec(`UPDATE pages SET content=? WHERE title=?;`, content, title)
				http.Redirect(w, r, "/pages/"+cTitle, http.StatusSeeOther)
				return
			}
		} else {
			if action == "Update" {
				if ctx.IsAdmin {
					editGroup := r.PostFormValue("editgroup")
					readGroup := r.PostFormValue("readgroup")
					var editGroupID, readGroupID string
					db.QueryRow(`SELECT id FROM groups WHERE name=?;`, editGroup).Scan(&editGroupID)
					db.QueryRow(`SELECT id FROM groups WHERE name=?;`, readGroup).Scan(&readGroupID)
					if editGroupID != "" {
						db.Exec(`UPDATE pages SET editgroupid=? WHERE title=?;`, editGroupID, title)
					} else {
						db.Exec(`UPDATE pages SET editgroupid=NULL WHERE title=?;`, title)
					}
					if readGroupID != "" {
						db.Exec(`UPDATE pages SET readgroupid=? WHERE title=?;`, readGroupID, title)
					} else {
						db.Exec(`UPDATE pages SET readgroupid=NULL WHERE title=?;`, title)
					}
				}
			}
			if action == "Delete" {
				if !ctx.IsAdmin && !isCRUDGroupMember {
					templates.Render(w, "accessdenied.html", map[string]interface{}{
						"ctx": ctx,
					})
					return
				}
				db.Exec(`DELETE FROM pages WHERE title=?;`, title)
			}
			http.Redirect(w, r, "/pages/", http.StatusSeeOther)
			return
		}
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
		"Title": title,
		"cTitle": cTitle,
		"Content": content,
	})
})