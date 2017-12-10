// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const loginSrc = `
{{ define "content" }}
<div class="form-container">
	<h3>Login to Femtowiki</h3>
	<form action="/login" method="POST">
		<div class="form-group">
			<input type="text" class="form-control" name="username" placeholder="Username">
		</div>
		<div class="form-group">
			<input type="password" class="form-control" name="passwd" placeholder="Password">
		</div>
		<div class="form-group">
			<a href="/forgotpass">Forgot password?</a>
		</div>
		<input type="submit" class="btn btn-default" value="Login">
	</form>
</div>
{{ end }}
`
