// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package views

import (
	"net/http"
	"github.com/s-gv/femtowiki/templates"
)

var LoginHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx Context) {
	templates.Render(w, "login.html", map[string]interface{}{
	})
})

var SignupHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx Context) {
	templates.Render(w, "signup.html", map[string]interface{}{
	})
})

var ForgotpassHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx Context) {
	templates.Render(w, "forgotpass.html", map[string]interface{}{
	})
})