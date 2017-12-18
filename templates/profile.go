// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const profileSrc = `
{{ define "content" }}
<div class="form-container">
	<h3>{{ .username }}</h3>
	<form action="/profile/update" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<input type="hidden" name="username" value="{{ .username }}">
		<div class="form-group">
			<input type="text" class="form-control" name="email" value="{{ .email }}" placeholder="Email">
		</div>
		{{ if .ctx.FlashMsg }}
		<div class="form-group">
			<span class="flash">{{ .ctx.FlashMsg }}</span>
		</div>
		{{ end }}
		<input type="submit" class="btn btn-default" value="Update Email">
		<div class="form-group">
			<a href="/changepass">Change password</a>
		</div>
		<div class="form-group">
			<a href="/logoutall">Logout all sessions</a>
		</div>
	</form>
</div>
{{ end }}
`
