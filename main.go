// Copyright (c) 2018 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/s-gv/femtowiki/templates"
	"gopkg.in/russross/blackfriday.v2"
)

var titleRegex *regexp.Regexp = regexp.MustCompile("<h1>([^<>/]+)</h1>")

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
		"Title":   title,
		"Content": template.HTML(html),
	})

	return buf.String()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	wikiRoot := flag.String("wikiroot", "", "Root of the wiki")
	htmlRoot := flag.String("htmlroot", "", "Root to the folder of generated html files")
	templateRoot := flag.String("templateroot", "", "Root of template files")

	flag.Parse()

	if *wikiRoot == "" {
		log.Fatalf("[ERROR] Specify wiki root\n")
	}
	var err error
	*wikiRoot, err = filepath.Abs(*wikiRoot)
	if err != nil {
		log.Fatalf("[ERROR] processing wiki root\n")
	}
	if *htmlRoot == "" {
		*htmlRoot = *wikiRoot
	}

	if *templateRoot != "" {
		templates.OverwriteTemplates(*templateRoot)
	}

	err = filepath.Walk(*wikiRoot,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			path, err = filepath.Rel(*wikiRoot, path)
			if err != nil {
				return err
			}
			if regexp.MustCompile(`(?i).md$`).MatchString(path) {
				buf, err := ioutil.ReadFile(*wikiRoot + "/" + path)
				if err != nil {
					return err
				}
				md := string(buf)
				html := renderMd(md)

				htmlfn := *htmlRoot + "/" + path[:len(path)-3] + ".html"
				htmldn := filepath.Dir(htmlfn)
				if !isDir(htmldn) {
					err = os.Mkdir(htmldn, 0750)
					if err != nil && !os.IsExist(err) {
						log.Fatal(err)
					}
				}
				ioutil.WriteFile(htmlfn, []byte(html), 0644)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
