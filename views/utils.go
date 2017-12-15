// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"log"
	"runtime/debug"
	"net/url"
	"crypto/rand"
	"encoding/base64"
)

func ErrServerHandler(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		log.Printf("[INFO] Recovered from panic: %s\n[INFO] Debug stack: %s\n", r, debug.Stack())
		http.Error(w, "Internal server error. This event has been logged.", http.StatusInternalServerError)
	}
}

func ErrNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func ErrForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "403 Forbidden", http.StatusForbidden)
}

func UA(handler func(w http.ResponseWriter, r *http.Request, ctx *Context)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer ErrServerHandler(w, r)
		sessionID := ""
		if cookie, err := r.Cookie("sessionid"); err == nil {
			sessionID = cookie.Value
		}
		ctx := ReadContext(sessionID)
		//log.Printf("[INFO] Request: %s\n", r.URL)
		handler(w, r, &ctx)
	}
}

func A(handler func(w http.ResponseWriter, r *http.Request, ctx *Context)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer ErrServerHandler(w, r)
		sessionID := ""
		if cookie, err := r.Cookie("sessionid"); err == nil {
			sessionID = cookie.Value
		}
		ctx := ReadContext(sessionID)
		if !ctx.IsUserValid {
			redirectURL := r.URL.Path
			if r.URL.RawQuery != "" {
				redirectURL += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, "/login?next="+url.QueryEscape(redirectURL), http.StatusSeeOther)
			return
		}
		if r.Method == "POST" && ctx.ValidateCSRFToken(r.PostFormValue("csrf")) != nil {
			ErrForbiddenHandler(w, r)
			return
		}
		//log.Printf("[INFO] Request: %s\n", r.URL)
		handler(w, r, &ctx)
	}
}

func randSeq(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		log.Panicf("[ERROR] Unable to generate random number: %s\n", err.Error())
	}
	return base64.URLEncoding.EncodeToString(b)
}
