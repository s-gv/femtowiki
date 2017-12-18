// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const adminUsersSrc = `
{{ define "content" }}
<div class="form-container">
	<h2>Users</h2>
	<ul>
		{{ range .users }}
		<li><a href="/profile?u={{ . }}">{{ . }}</a></li>
		{{ end }}
	</ul>
	<a href="/signup">Signup new user</a>
</div>
{{ end }}
`