// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

var pageListSrc = `
{{ define "content" }}
<h1>Pages</h1>
<div class="table">
	<div class="trow">
		<div class="tcol3"><strong>Title</strong></div>
		<div class="tcol6"><strong>Read Group</strong></div>
		<div class="tcol6"><strong>Edit Group</strong></div>
	</div>
	{{ range .pages }}
	<div class="trow">
		<form action="/editpage" method="POST">
			<input type="hidden" name="csrf" value="{{ $.ctx.CSRFToken }}">
			<input type="hidden" name="t" value="{{ .CTitle }}">
			<input type="hidden" name="meta" value="true">
			<div class="tcol3">
				<a href="/pages/{{ .CTitle }}">{{ .Title }}</a>
			</div>
			<div class="tcol6">
				<input type="text" class="form-control" name="readgroup" value="{{ .ReadGroup }}"{{ if not $.ctx.IsAdmin}} disabled{{ end }}>
			</div>
			<div class="tcol6">
				<input type="text" class="form-control" name="editgroup" value="{{ .EditGroup }}"{{ if not $.ctx.IsAdmin}} disabled{{ end }}>
			</div>
			<div class="tcol6">
				<input type="submit" class="btn btn-default" name="action" value="Update"{{ if not $.ctx.IsAdmin}} disabled{{ end }}>
			</div>
			<div class="tcol6">
				<input type="submit" class="btn btn-danger" name="action" value="Delete">
			</div>
		</form>
	</div>
	{{ end }}
	<p>Note: Set the group column to blank to allow any registered user to perform the relevant action.</p>
</div>
<h3>New page</h3>
<div style="max-width: 300px;">
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