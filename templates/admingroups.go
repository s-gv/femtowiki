// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const adminGroupsSrc = `
{{ define "content" }}
<div class="form-container">
	<h2>Groups</h2>
	{{ if .groups }}
	<ul>
		{{ range .groups }}
		<li><a href="/admin/groupmembers?g={{ . }}">{{ . }}</a></li>
		{{ end }}
	</ul>
	{{ else }}
	<p>No groups found</p>
	{{ end }}
	<h3>New group</h3>
	<form action="/admin/groups/new" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<div class="form-group">
			<input type="text" class="form-control" name="groupname" placeholder="Group name">
		</div>
		{{ if .ctx.FlashMsg }}
		<div class="form-group">
			<span class="flash">{{ .ctx.FlashMsg }}</span>
		</div>
		{{ end }}
		<input type="submit" class="btn btn-default" value="Create group">
	</form>
</div>
{{ end }}
`