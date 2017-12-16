// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const signupSrc = `
{{ define "content" }}
<div class="form-container">
	<h3>Signup</h3>
	<form action="/signup" method="POST">
		<div class="form-group">
			<input type="text" class="form-control" name="username" placeholder="Username">
		</div>
		<div class="form-group">
			<input type="password" class="form-control" name="passwd" placeholder="Password">
		</div>
		<div class="form-group">
			<input type="password" class="form-control" name="passwd2" placeholder="Confirm password">
		</div>
		<div class="form-group">
			<span>Already have an account?</span> <a href="/login">Signin</a>
		</div>
		{{ if .ctx.FlashMsg }}
		<div class="form-group">
			<span class="flash">{{ .ctx.FlashMsg }}</span>
		</div>
		{{ end }}
		<input type="submit" class="btn btn-default" value="Signup">
	</form>
</div>
{{ end }}
`