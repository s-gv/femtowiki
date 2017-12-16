// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const changepassSrc = `
{{ define "content" }}
<div class="form-container">
	<h3>Change password</h3>
	<form action="/changepass" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<div class="form-group">
			<input type="text" class="form-control" name="username" value="{{ .ctx.UserName }}" disabled>
		</div>
		<div class="form-group">
			<input type="password" class="form-control" name="oldpasswd" placeholder="Current password">
		</div>
		<div class="form-group">
			<input type="password" class="form-control" name="passwd" placeholder="New password">
		</div>
		<div class="form-group">
			<input type="password" class="form-control" name="passwd2" placeholder="Confirm new password">
		</div>
		{{ if .ctx.FlashMsg }}
		<div class="form-group">
			<span class="flash">{{ .ctx.FlashMsg }}</span>
		</div>
		{{ end }}
		<input type="submit" class="btn btn-default" value="Change Password">
	</form>
</div>
{{ end }}
`