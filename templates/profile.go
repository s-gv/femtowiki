// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const profileSrc = `
{{ define "content" }}
<div class="form-container">
	<h2>User profile: {{ .username }}</h2>
	{{ if .ctx.FlashMsg }}
	<span class="flash">{{ .ctx.FlashMsg }}</span>
	{{ end }}
	<h3>Email</h3>
	<form action="/profile/update" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<input type="hidden" name="username" value="{{ .username }}">
		<div class="form-group">
			<input type="text" class="form-control" name="email" value="{{ .email }}" placeholder="Email">
		</div>
		<input type="submit" class="btn btn-default" value="Update Email">
	</form>
	<h3>Password</h3>
	<form action="/changepass" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<input type="hidden" name="username" value="{{ .username }}">
		{{ if not .ctx.IsAdmin }}
		<div class="form-group">
			<input type="password" class="form-control" name="oldpasswd" placeholder="Current password">
		</div>
		{{ end }}
		<div class="form-group">
			<input type="password" class="form-control" name="passwd" placeholder="New password">
		</div>
		<div class="form-group">
			<input type="password" class="form-control" name="passwd2" placeholder="Confirm new password">
		</div>
		<input type="submit" class="btn btn-default" value="Change Password">
		<div class="form-group">
			<a href="/logoutall?u={{ .username }}">Logout all sessions</a>
		</div>
	</form>
	{{ if .ctx.IsAdmin }}
	{{ if ne .ctx.UserName .username }}
	<h3>Ban/Unban user</h3>
	<p>A banned user will be prevented from logging in.</p>
	{{ if .IsBanned }}
	<form action="/profile/unban" method="POST">
	{{ else }}
	<form action="/profile/ban" method="POST">
	{{ end }}
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<input type="hidden" name="username" value="{{ .username }}">
		<input type="submit" class="btn btn-default" value="{{ if .IsBanned }}Unban{{ else }}Ban{{ end }} {{ .username }}">
	</form>
	{{ end }}
	{{ end }}
</div>
{{ end }}
`
