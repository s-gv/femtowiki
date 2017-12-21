// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

var pageListSrc = `
{{ define "content" }}
<div class="form-container">
	<h1>Pages</h1>
	<ul>
		{{ range .pages }}
		<li><a href="{{ .URL }}">{{ .Title }}</a></li>
		{{ end }}
	</ul>
	<h3>New page</h3>
	<form action="/newpage" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<div class="form-group">
			<input type="text" class="form-control" name="title" placeholder="Page title">
		</div>
		{{ if .ctx.FlashMsg }}
		<div class="form-group" id="flash">
			<span class="flash">{{ .ctx.FlashMsg }}</span>
		</div>
		{{ end }}
		<input type="submit" class="btn btn-default" value="Create page">
	</form>
</div>
{{ end }}
`