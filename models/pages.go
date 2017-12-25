// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"errors"
	"github.com/s-gv/femtowiki/models/db"
	"html/template"
	"regexp"
)

var (
	IndexPage = "Home Page"
)

var snippetRe *regexp.Regexp

func init() {
	snippetRe = regexp.MustCompile("__(.+)__")
}

type PagesSearchResult struct {
	Title   string
	CTitle  string
	Snippet template.HTML
}

func IsPageTitleValid(title string) error {
	if len(title) < 2 || len(title) > 200 {
		return errors.New("Should have 2-200 characters")
	}
	for _, ch := range title {
		if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') && (ch < '0' || ch > '9') && (ch != '(') && (ch != ')') && (ch != ' ') && (ch != '-') && (ch != '.') {
			return errors.New("Only alphabets, numbers, parenthesis, and hyphens are supported.")
		}
	}
	return nil
}

func PageSearch(terms string) []PagesSearchResult {
	rows := db.Query(`SELECT title, snippet(pages_search_index, 1, '__', '__', '', 20) FROM pages_search_index WHERE pages_search_index MATCH ?;`, terms)
	results := []PagesSearchResult{}
	for rows.Next() {
		res := PagesSearchResult{}
		var snippet string
		rows.Scan(&res.Title, &snippet)
		res.Snippet = template.HTML(snippetRe.ReplaceAllString(template.HTMLEscapeString(snippet), "<b>$1</b>"))
		results = append(results, res)
	}
	return results
}
