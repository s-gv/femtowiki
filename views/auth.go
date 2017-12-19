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
	"github.com/s-gv/femtowiki/models/db"
	"github.com/s-gv/femtowiki/models"
	"strings"
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
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		passwd := r.PostFormValue("passwd")
		passwd2 := r.PostFormValue("passwd2")

		if ctx.IsAdmin || !ctx.Config.SignupDisabled {
			if err := models.ValidateName(username); err == nil {
				if passwd == passwd2 {
					if err := models.ValidatePasswd(passwd); err == nil {
						if err := models.CreateUser(username, passwd, "", false); err == nil {
							if ctx.IsAdmin {
								http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
								return
							} else {
								if err := ctx.Authenticate(username, passwd); err == nil {
									http.SetCookie(w, &http.Cookie{Name: "sessionid", Path: "/", Value: ctx.SessionID, HttpOnly: true})
									http.Redirect(w, r, "/", http.StatusSeeOther)
									return
								} else {
									ctx.FlashMsg = err.Error()
								}
							}
						} else {
							ctx.FlashMsg = err.Error()
						}
					} else {
						ctx.FlashMsg = err.Error()
					}
				} else {
					ctx.FlashMsg = "Passwords don't match"
				}
			} else {
				ctx.FlashMsg = err.Error()
			}
		} else {
			ctx.FlashMsg = "Signup disabled. Contact admin."
		}
	}
	templates.Render(w, "signup.html", map[string]interface{}{
		"ctx": ctx,
	})
})

var ChangepassHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if r.Method != "POST" {
		ErrForbiddenHandler(w, r)
		return
	}

	username := r.PostFormValue("username")
	oldPasswd := r.PostFormValue("oldpasswd")
	newPasswd := r.PostFormValue("passwd")
	newPasswd2 := r.PostFormValue("passwd2")

	if !ctx.IsAdmin {
		if err := models.VerifyPasswd(username, oldPasswd); err != nil {
			ctx.SetFlashMsg(err.Error())
			http.Redirect(w, r, "/profile?u="+username, http.StatusSeeOther)
			return
		}
	}

	if err := models.ValidatePasswd(newPasswd); err != nil {
		ctx.SetFlashMsg(err.Error())
		http.Redirect(w, r, "/profile?u="+username, http.StatusSeeOther)
		return
	}
	if newPasswd != newPasswd2 {
		ctx.SetFlashMsg("New passwords do not match")
		http.Redirect(w, r, "/profile?u="+username, http.StatusSeeOther)
		return
	}

	if err := models.UpdateUserPasswd(username, newPasswd); err != nil {
		ctx.SetFlashMsg(err.Error())
		http.Redirect(w, r, "/profile?u="+username, http.StatusSeeOther)
		return
	}

	ctx.SetFlashMsg("Password changed successfully")
	http.Redirect(w, r, "/profile?u="+username, http.StatusSeeOther)
	return
})

var ForgotpassHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		row := db.QueryRow(`SELECT email FROM users WHERE username=?;`, username)
		var email string
		if err := row.Scan(&email); err == nil {
			if strings.Contains(email, "@") {
				resetToken := randSeq(40)
				db.Exec(`UPDATE users SET reset_token=?, reset_token_date=? WHERE username=?;`, resetToken, int64(time.Now().Unix()), username)
				resetLink := "https://" + r.Host + "/resetpass?token=" + resetToken
				sub := ctx.Config.WikiName + " Password Recovery"
				msg := "Someone (hopefully you) requested we reset your password at " + ctx.Config.WikiName + ".\r\n" +
					"If you want to change it, visit "+resetLink+"\r\n\r\nIf not, just ignore this message."

				SendMail(email, sub, msg, ctx.Config)
				ctx.FlashMsg = "Password reset link has been sent to your email"
			} else {
				ctx.FlashMsg = "We don't have your email. Please contact the admin to reset your password"
			}
		} else {
			ctx.FlashMsg = "User not found"
		}
	}
	templates.Render(w, "forgotpass.html", map[string]interface{}{
		"ctx": ctx,
	})
})

var ResetpassHandler = UA(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	resetToken := r.FormValue("token")

	row := db.QueryRow(`SELECT username, reset_token_date FROM users WHERE reset_token=?;`, resetToken)
	var username string
	var rDate int64
	if err := row.Scan(&username, &rDate); err != nil {
		ErrNotFoundHandler(w, r)
		return
	}
	resetTokenDate := time.Unix(rDate, 0)
	if resetTokenDate.Before(time.Now().Add(-100*time.Hour)) {
		ErrNotFoundHandler(w, r)
		return
	}

	if r.Method == "POST" {
		passwd := r.PostFormValue("passwd")
		passwd2 := r.PostFormValue("passwd2")
		if passwd == passwd2 {
			if err := models.ValidatePasswd(passwd); err == nil {
				models.UpdateUserPasswd(username, passwd)
				db.Exec(`UPDATE users SET reset_token_date=0 WHERE username=?;`, username)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			} else {
				ctx.FlashMsg = err.Error()
			}
		} else {
			ctx.FlashMsg = "New passwords do not match"
		}
	}
	templates.Render(w, "resetpass.html", map[string]interface{}{
		"ctx": ctx,
		"ResetToken": resetToken,
	})
})

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	defer ErrServerHandler(w, r)
	http.SetCookie(w, &http.Cookie{Name: "sessionid", Value: "", Expires: time.Now().Add(-300*time.Hour), HttpOnly: true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

var LogoutAllHandler = A(func(w http.ResponseWriter, r *http.Request, ctx *Context) {
	username := r.FormValue("u")
	if !ctx.IsAdmin && (username != ctx.UserName) {
		ErrForbiddenHandler(w, r)
		return
	}

	var userID string
	if db.QueryRow(`SELECT id FROM users WHERE users.username=?;`, username).Scan(&userID) == nil {
		db.Exec(`DELETE FROM sessions WHERE userid=?;`, userID)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
})