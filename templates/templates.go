// Copyright (c) 2018 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"io"
	"log"
	"html/template"
)

var Main = template.Must(template.New("main").Parse(mainSrc))

func Render(wr io.Writer, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(wr, data)
	if err != nil {
		log.Panicf("[ERROR] Error rendering %s: %s\n", tmpl.Name(), err)
	}
}