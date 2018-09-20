// Copyright (c) 2018 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"gopkg.in/russross/blackfriday.v2"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"flag"
	"path/filepath"
	"os"
	"io/ioutil"
	"strings"
	"github.com/s-gv/femtowiki/templates"
	"html/template"
	"bytes"
	"regexp"
)

var titleRegex *regexp.Regexp

func renderMd(markdown string) string {
	unsafe := blackfriday.Run([]byte(strings.Replace(markdown, "\r\n", "\n", -1)))
	html := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	matches := titleRegex.FindStringSubmatch(html)
	title := "Femtowiki"
	if len(matches) >= 2 && matches[1] != "" {
		title = matches[1]
	}

	buf := new(bytes.Buffer)
	templates.Render(buf, templates.Main, map[string]interface{}{
		"Title": title,
		"Content": template.HTML(html),
	})

	return buf.String()
}



func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	wikiRoot := flag.String("wikiroot", "", "Root of the wiki")
	//templateRoot := flag.String("templateroot", "", "Root of template files")

	flag.Parse()

	if *wikiRoot == "" {
		log.Fatalf("[ERROR] Specify wiki root\n")
	}

	titleRegex = regexp.MustCompile("<h1>([^<>/]+)</h1>")

	err := filepath.Walk(*wikiRoot,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path[len(path)-3:] == ".md" {
				buf, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				md := string(buf)
				html := renderMd(md)

				ioutil.WriteFile(path[:len(path)-3]+".html", []byte(html), 0644)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

}