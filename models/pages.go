// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package models

import (
	"errors"
)

var (
	IndexPage = "Home Page"
)

func IsPageTitleValid(title string) error {
	if len(title) < 2 || len(title) > 200 {
		return errors.New("Should have 2-200 characters")
	}
	for _, ch := range title {
		if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') && (ch < '0' || ch > '9') && (ch != '(') && (ch != ')') && (ch != ' ') && (ch != '-') {
			return errors.New("Only alphabets, numbers, parenthesis, and hyphens are supported.")
		}
	}
	return nil
}