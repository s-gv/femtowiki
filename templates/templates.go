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
	template.Must(tmpls["index.html"].New("adminindex").Parse(indexSrc))
}

func Render(wr io.Writer, template string, data interface{}) {
	err := tmpls[template].Execute(wr, data)
	if err != nil {
		log.Panicf("[ERROR] Error rendering %s: %s\n", template, err)
	}
}