// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

var fileListSrc = `
{{ define "content" }}
<h1>Files</h1>
<div class="table">
	<div class="trow">
		<div class="tcol3"><strong>Name</strong></div>
		<div class="tcol6"><strong>Read Group</strong></div>
	</div>
	{{ if .files }}
	{{ range .files }}
	<div class="trow">
		<form action="/editfile" method="POST">
			<input type="hidden" name="csrf" value="{{ $.ctx.CSRFToken }}">
			<input type="hidden" name="t" value="{{ .CTitle }}">
			<div class="tcol3">
				<a href="/files/{{ .CTitle }}">{{ .Title }}</a>
			</div>
			<div class="tcol6">
				<input type="text" class="form-control" name="readgroup" value="{{ .ReadGroup }}"{{ if not $.ctx.IsAdmin}} disabled{{ end }}>
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
	<p>Note: Set the group column to blank to allow any user to perform the relevant action.</p>
	{{ else }}
	<p>No files found.</p>
	{{ end }}
</div>
<h3>New file</h3>
<div style="max-width: 300px;">
	<form action="/newfile" method="POST" enctype="multipart/form-data">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<div class="form-group">
			<input type="file" class="form-control" name="file">
		</div>
		{{ if .ctx.FlashMsg }}
		<div class="form-group" id="flash">
			<span class="flash">{{ .ctx.FlashMsg }}</span>
		</div>
		{{ end }}
		<input type="submit" class="btn btn-default" value="Upload">
	</form>
</div>
{{ end }}
`