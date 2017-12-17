// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"io"
	"log"
	"html/template"
)

var tmpls = make(map[string]*template.Template)

func init() {
	tmpls["index.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["index.html"].New("index").Parse(indexSrc))

	tmpls["login.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["login.html"].New("login").Parse(loginSrc))

	tmpls["signup.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["signup.html"].New("signup").Parse(signupSrc))

	tmpls["changepass.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["changepass.html"].New("changepass").Parse(changepassSrc))

	tmpls["forgotpass.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["forgotpass.html"].New("forgotpass").Parse(forgotpassSrc))

	tmpls["resetpass.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["resetpass.html"].New("resetpass").Parse(resetpassSrc))

	tmpls["profile.html"] = template.Must(template.New("base").Parse(baseSrc))
	template.Must(tmpls["profile.html"].New("profile").Parse(profileSrc))
}

func Render(wr io.Writer, template string, data interface{}) {
	err := tmpls[template].Execute(wr, data)
	if err != nil {
		log.Panicf("[ERROR] Error rendering %s: %s\n", template, err)
	}
}