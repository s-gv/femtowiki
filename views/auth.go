// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
	"net/url"
	"fmt"
	"time"
	"html/template"
)

var LoginHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	redirectURL, err := url.QueryUnescape(r.FormValue("next"))
	if redirectURL == "" || err != nil {
		redirectURL = "/"
	}
	if ctx.IsUserValid {
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		userName := r.PostFormValue("username")
		passwd := r.PostFormValue("passwd")
		if len(userName) > 200 || len(passwd) > 200 {
			fmt.Fprint(w, "username / password too long.")
			return
		}
		if err = ctx.Authenticate(userName, passwd); err == nil {
			http.SetCookie(w, &http.Cookie{Name: "sessionid", Path: "/", Value: ctx.SessionID, HttpOnly: true})
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			return
		} else {
			ctx.FlashMsg = err.Error()
		}
	}
	templates.Render(w, "login.html", map[string]interface{}{
		"ctx": ctx,
		"next": template.URL(url.QueryEscape(redirectURL)),
	})
})

var SignupHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	templates.Render(w, "signup.html", map[string]interface{}{
	})
})

var ForgotpassHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	templates.Render(w, "forgotpass.html", map[string]interface{}{
	})
})

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	defer ErrServerHandler(w, r)
	http.SetCookie(w, &http.Cookie{Name: "sessionid", Value: "", Expires: time.Now().Add(-300*time.Hour), HttpOnly: true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}