// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"strings"
	"github.com/s-gv/femtowiki/models"
	"github.com/s-gv/femtowiki/templates"
	"database/sql"
	"github.com/s-gv/femtowiki/models/db"
	"os"
	"io"
	"log"
	"time"
)

var FilesHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	isPageMaster := ctx.IsAdmin || models.IsUserInFileMasterGroup(ctx.UserName)
	cTitle := r.URL.Path[7:] // r.URL will be /files/<page_title>
	title := strings.Replace(cTitle, "_", " ", -1)
	if title == "" {
		// List all files
		if !isPageMaster {
			templates.Render(w, "accessdenied.html", map[string]interface{}{
				"ctx": ctx,
			})
			return
		}
		type File struct {
			Title     string
			CTitle    string
			ReadGroup string
			EditGroup string
		}
		var files []File
		rows := db.Query(`SELECT title, readgroupid, editgroupid FROM pages WHERE is_file=? ORDER BY title;`, true)
		for rows.Next() {
			file := File{}
			var readGroupID, editGroupID sql.NullString
			rows.Scan(&file.Title, &readGroupID, &editGroupID)
			file.CTitle = strings.Replace(file.Title, " ", "_", -1)
			if readGroupID.Valid {
				db.QueryRow(`SELECT name FROM groups WHERE id=?;`, readGroupID).Scan(&file.ReadGroup)
			}
			if editGroupID.Valid {
				db.QueryRow(`SELECT name FROM groups WHERE id=?;`, editGroupID).Scan(&file.EditGroup)
			}
			files = append(files, file)
		}
		templates.Render(w, "filelist.html", map[string]interface{}{
			"ctx": ctx,
			"files": files,
		})
		return
	}
	dataDir := ctx.Config.DataDir
	if dataDir  == "" {
		ErrNotFoundHandler(w, r)
		return
	}
	if dataDir[len(dataDir)-1] != '/' {
		dataDir = dataDir + "/"
	}
	row := db.QueryRow(`SELECT readgroupid FROM pages WHERE title=?;`, title)
	var readGroupID sql.NullString
	if row.Scan(&readGroupID) != nil {
		ErrNotFoundHandler(w, r)
		return
	}
	if !isPageMaster && readGroupID.Valid {
		row := db.QueryRow(`SELECT groupmembers.id FROM groupmembers INNER JOIN users ON users.id=groupmembers.userid AND users.username=? WHERE groupmembers.groupid=?;`, ctx.UserName, readGroupID)
		var tmp string
		if row.Scan(&tmp) != nil {
			templates.Render(w, "accessdenied.html", map[string]interface{}{
				"ctx": ctx,
			})
			return
		}
	}
	http.ServeFile(w, r, ctx.Config.DataDir + title)
})

var FileCreateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	isPageMaster := ctx.IsAdmin || models.IsUserInFileMasterGroup(ctx.UserName)
	if !isPageMaster {
		templates.Render(w, "accessdenied.html", map[string]interface{}{
			"ctx": ctx,
		})
		return
	}
	dataDir := ctx.Config.DataDir
	if dataDir  == "" {
		ctx.SetFlashMsg("DataDir not configured. Contact admin.")
		http.Redirect(w, r, "/files/#flash", http.StatusSeeOther)
		return
	}
	if dataDir[len(dataDir)-1] != '/' {
		dataDir = dataDir + "/"
	}
	r.ParseMultipartForm(32*1024*1024)
	file, handler, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		if handler.Filename != "" {
			title := handler.Filename
			if err := models.IsPageTitleValid(title); err == nil {
				var tmp string
				if db.QueryRow(`SELECT id FROM pages WHERE title=?;`, title).Scan(&tmp) == sql.ErrNoRows {
					f, err := os.OpenFile(dataDir+title, os.O_WRONLY|os.O_CREATE, 0666)
					if err == nil {
						defer f.Close()
						io.Copy(f, file)
						now := time.Now().Unix()
						db.Exec(`INSERT INTO pages(title, is_file, created_date, updated_date) VALUES(?, ?, ?, ?);`, title, true, now, now)
					} else {
						log.Panicf("[ERROR] Error writing file: %s\n", err)
					}
				} else {
					ctx.SetFlashMsg("File already exists")
				}
			} else {
				ctx.SetFlashMsg(err.Error())
			}
		} else {
			ctx.SetFlashMsg("Choose file to upload")
		}
	}
	http.Redirect(w, r, "/files/", http.StatusSeeOther)
})

var FileUpdateHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}
	isPageMaster := ctx.IsAdmin || models.IsUserInFileMasterGroup(ctx.UserName)
	if !isPageMaster {
		templates.Render(w, "accessdenied.html", map[string]interface{}{
			"ctx": ctx,
		})
		return
	}
	cTitle := r.FormValue("t")
	title := strings.Replace(cTitle, "_", " ", -1)
	action := r.PostFormValue("action")
	if action == "Update" {
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
		ctx.SetFlashMsg("Updated file " + title)
	}
	if action == "Delete" {
		dataDir := ctx.Config.DataDir
		if dataDir != "" {
			if dataDir[len(dataDir)-1] != '/' {
				dataDir = dataDir + "/"
			}
			db.Exec(`DELETE FROM pages WHERE title=?;`, title)
			if err := os.Remove(dataDir + title); err != nil {
				log.Printf("[ERROR] Error deleting file: %s\n", err.Error())
			}
			ctx.SetFlashMsg("Deleted file " + title)
		} else {
			ctx.SetFlashMsg("DataDir not configured")
		}
	}
	http.Redirect(w, r, "/files/", http.StatusSeeOther)
})