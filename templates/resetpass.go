// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const resetpassSrc = `
{{ define "content" }}
<div class="form-container">
	<h3>Reset password</h3>
	<form action="/resetpass" method="POST">
		<input type="hidden" name="token" value="{{ .ResetToken }}">
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
		<input type="submit" class="btn btn-default" value="Reset Password">
	</form>
</div>
{{ end }}
`