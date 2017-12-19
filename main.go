// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"flag"
	"log"
	"github.com/s-gv/femtowiki/models/db"
	"net/http"
	"github.com/s-gv/femtowiki/models"
	"net/http/fcgi"
	"time"
	"github.com/s-gv/femtowiki/views"
)

func getCreds() (string, string) {
	var userName string
	fmt.Printf("Username: ")
	fmt.Scan(&userName)

	fmt.Printf("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n")
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return "", ""
	}
	if len(password) < 8 {
		fmt.Printf("[ERROR] Password should have at least 8 characters.\n")
		return "", ""
	}

	fmt.Printf("Password (again): ")
	password2, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n")
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}

	pass := string(password)
	pass2 := string(password2)
	if pass != pass2 {
		fmt.Printf("[ERROR] The two psasswords do not match.\n")
		return "", ""
	}

	return userName, pass
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dsn := flag.String("dsn", "femtowiki.sqlite3", "Data source name")
	dbDriver := flag.String("dbdriver", "sqlite3", "DB driver name")
	addr := flag.String("addr", ":9123", "Port to listen on")
	shouldMigrate := flag.Bool("migrate", false, "Migrate DB")
	createSuperUser := flag.Bool("createsuperuser", false, "Create superuser (interactive)")
	createUser := flag.Bool("createuser", false, "Create user. Optional arguments: <username> <password> <email>")
	changePasswd := flag.Bool("changepasswd", false, "Change password")
	deleteSessions := flag.Bool("deletesessions", false, "Delete all sessions (logout all users)")
	fcgiMode := flag.Bool("fcgi", false, "Fast CGI rather than listening on a port")

	flag.Parse()

	db.Init(*dbDriver, *dsn)

	if *shouldMigrate {
		models.Migrate()
		return
	}

	if models.IsMigrationNeeded() {
		log.Fatalf("[ERROR] DB migration needed.\n")
	}

	if *createSuperUser {
		fmt.Printf("Creating superuser...\n")
		userName, pass := getCreds()
		if userName != "" && pass != "" {
			if err := models.CreateSuperUser(userName, pass); err != nil {
				fmt.Printf("Error creating superuser: %s\n", err)
			}
		}
		return
	}

	if *createUser {
		args := flag.Args()
		var username, passwd, email string
		if len(args) >= 2 {
			username = args[0]
			passwd = args[1]
		} else {
			username, passwd = getCreds()
		}
		if len(args) >= 3 {
			email = args[2]
		}
		if username != "" && passwd != "" {
			if err := models.CreateUser(username, passwd, email, false); err != nil {
				fmt.Printf("Error creating user: %s\n", err)
			}
		} else {
			fmt.Printf("Error: Username and password cannot be blank.\n")
		}
		return
	}

	if *changePasswd {
		userName, pass := getCreds()
		if userName != "" && pass != "" {
			if err := models.UpdateUserPasswd(userName, pass); err != nil {
				fmt.Printf("Error changing password: %s\n", err)
			}
		}
		return
	}

	if *deleteSessions {
		db.Exec(`DELETE FROM sessions;`)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", views.IndexHandler)

	mux.HandleFunc("/static/css/femtowiki.css", views.StyleHandler)

	mux.HandleFunc("/static/js/femtowiki.js", views.ScriptHandler)

	mux.HandleFunc("/favicon.ico", views.FaviconHandler)

	mux.HandleFunc("/login", views.LoginHandler)
	mux.HandleFunc("/signup", views.SignupHandler)
	mux.HandleFunc("/changepass", views.ChangepassHandler)
	mux.HandleFunc("/forgotpass", views.ForgotpassHandler)
	mux.HandleFunc("/resetpass", views.ResetpassHandler)
	mux.HandleFunc("/logout", views.LogoutHandler)
	mux.HandleFunc("/logoutall", views.LogoutAllHandler)

	mux.HandleFunc("/profile", views.ProfileHandler)
	mux.HandleFunc("/profile/update", views.ProfileUpdateHandler)
	mux.HandleFunc("/profile/ban", views.ProfileBanHandler)
	mux.HandleFunc("/profile/unban", views.ProfileUnbanHandler)

	mux.HandleFunc("/admin", views.AdminHandler)
	mux.HandleFunc("/admin/config", views.AdminConfigUpdateHandler)
	mux.HandleFunc("/admin/header", views.AdminHeaderUpdateHandler)
	mux.HandleFunc("/admin/footer", views.AdminFooterUpdateHandler)
	mux.HandleFunc("/admin/nav", views.AdminNavUpdateHandler)
	mux.HandleFunc("/admin/illegalnames", views.AdminIllegalNamesUpdateHandler)
	mux.HandleFunc("/admin/users", views.AdminUserHandler)
	mux.HandleFunc("/admin/groups", views.AdminGroupHandler)
	mux.HandleFunc("/admin/groups/new", views.AdminGroupCreateHandler)
	mux.HandleFunc("/admin/groups/delete", views.AdminGroupDeleteHandler)
	mux.HandleFunc("/admin/groupmembers", views.AdminGroupMembersHandler)
	mux.HandleFunc("/admin/groupmembers/new", views.AdminGroupMemberCreateHandler)
	mux.HandleFunc("/admin/groupmembers/delete", views.AdminGroupMemberDeleteHandler)
	mux.HandleFunc("/admin/crudgroup", views.AdminCRUDGroupHandler)

	if *fcgiMode {
		fcgi.Serve(nil, mux)
		return
	}

	srv := &http.Server{
		Handler: mux,
		Addr: *addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("[INFO] Starting femtowiki at", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panicf("[ERROR] %s\n", err)
	}
}