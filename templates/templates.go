// Copyright (c) 2018 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"io"
	"log"
	"html/template"
	"io/ioutil"
)

var Main = template.Must(template.New("main").Parse(mainSrc))

func OverwriteTemplates(rootDir string) {
	mainSrc, err := ioutil.ReadFile(rootDir + "/main.html")
	if err == nil {
		Main = template.Must(template.New("main").Parse(string(mainSrc)))
	} else {
		log.Fatalf("[ERROR] Error reading template: %s", err)
	}
}

func Render(wr io.Writer, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(wr, data)
	if err != nil {
		log.Panicf("[ERROR] Error rendering %s: %s\n", tmpl.Name(), err)
	}
}