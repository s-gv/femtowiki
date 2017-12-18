// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const adminSrc = `
{{ define "content" }}




<div class="form-container">
	<h2>Admin section</h2>

	{{ if .ctx.FlashMsg }}
	<div class="form-group">
		<span class="flash">{{ .ctx.FlashMsg }}</span>
	</div>
	{{ end }}

	<form action="/admin/config" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>Config</h3>
		<div class="form-group">
			Default Config: <pre>{{ .DefaultConfig }}</pre>
		</div>
		<div class="form-group"><textarea class="form-control" rows="15" name="config">{{ .config }}</textarea></div>
		<input type="submit" class="btn btn-default" value="Update Config">
	</form>
</div>

<div class="form-container">
	<form action="/admin/nav" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>Nav section</h3>
		<div class="form-group">
			Default Nav: <pre>{{ .DefaultNav }}</pre>
		</div>
		<div class="form-group"><textarea class="form-control" rows="15" name="nav">{{ .nav }}</textarea></div>
		<input type="submit" class="btn btn-default" value="Update Nav">
	</form>
</div>

<div class="form-container">
	<form action="/admin/header" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>Header links</h3>
		<div class="form-group">
			Default Header: <pre>{{ .DefaultHeader }}</pre>
		</div>
		<div class="form-group"><textarea class="form-control" rows="15" name="header">{{ .header }}</textarea></div>
		<input type="submit" class="btn btn-default" value="Update Header">
	</form>
</div>

<div class="form-container">
	<form action="/admin/footer" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>Footer links</h3>
		<div class="form-group">
			Default Footer: <pre>{{ .DefaultFooter }}</pre>
		</div>
		<div class="form-group"><textarea class="form-control" rows="15" name="footer">{{ .footer }}</textarea></div>
		<input type="submit" class="btn btn-default" value="Update Footer">
	</form>
</div>

<div class="form-container">
	<form action="/admin/illegalusernames" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>Illegal usernames</h3>
		<div class="form-group">
			Sample illegal username list: <pre>{{ .DefaultIllegalUsernames }}</pre>
		</div>
		<div class="form-group"><textarea class="form-control" rows="15" name="illegal_usernames">{{ .illegal_usernames }}</textarea></div>
		<input type="submit" class="btn btn-default" value="Update Illegal usernames">
	</form>
</div>

{{ end }}
`

