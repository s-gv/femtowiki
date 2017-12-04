// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const indexSrc = `
{{ define "content" }}
"Responsive Design" is the strategy of making a site that "responds" to the browser and device that it is being shown on... by looking awesome no matter what.

Media queries are the most powerful tool for doing this. Let's take our layout that uses percent widths and have it display in one column when the browser is too small to fit the menu in the sidebar:

Tada! Now our layout looks great even on mobile browsers. Here are some popular sites that use media queries this way. There is a whole lot more you can detect than min-width and max-width: check out the MDN documentation to learn more
{{ end }}`