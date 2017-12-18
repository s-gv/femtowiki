// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const adminGroupMembersSrc = `
{{ define "content" }}
<div class="form-container">
	<h2>Group: {{ .groupname }}</h2>
	{{ if .members }}
	<ul>
		{{ range .members }}
		<li>
			<a href="/profile?u={{ . }}">{{ . }}</a>
			<form class="form-inline" action="/admin/groupmembers/delete" method="POST">
				<input type="hidden" name="csrf" value="{{ $.ctx.CSRFToken }}">
				<input type="hidden" name="groupname" value="{{ $.groupname }}">
				<input type="hidden" name="username" value="{{ . }}">
				<input type="submit" class="btn btn-danger btn-inline" value="Delete">
			</form>
		</li>
		{{ end }}
	</ul>
	{{ else }}
	<p>No members in this group</p>
	{{ end }}
	<h3>New member</h3>
	<form action="/admin/groupmembers/new" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<input type="hidden" name="groupname" value="{{ .groupname }}">
		<div class="form-group">
			<input type="text" class="form-control" name="username" placeholder="Username">
		</div>
		{{ if .ctx.FlashMsg }}
		<div class="form-group">
			<span class="flash">{{ .ctx.FlashMsg }}</span>
		</div>
		{{ end }}
		<input type="submit" class="btn btn-default" value="Add member">
	</form>
	<h3>Delete group</h3>
	<form action="/admin/groups/delete" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<input type="hidden" name="groupname" value="{{ .groupname }}">
		<input type="submit" class="btn btn-danger" value="Delete group">
	</form>
</div>
{{ end }}
`