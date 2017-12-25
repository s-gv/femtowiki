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

	<h3><a href="/admin/users">Users</a></h3>
	<h3><a href="/admin/groups">Groups</a></h3>
	<form action="/admin/pagemaster" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>PageMaster group</h3>
		<p>
			Members of this group will be permitted to create new pages and delete existing pages in the wiki.
			Note that the group 'everybody' is a special group that includes all registered users.
		</p>
		<div class="form-group">
			<input type="text" class="form-control" name="group" value="{{ .PageMasterGroup }}" placeholder="Group name">
		</div>
		<input type="submit" class="btn btn-default" value="Update">
	</form>
</div>

<div class="form-container">
	<form action="/admin/filemaster" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>FileMaster group</h3>
		<p>
			Members of this group will be permitted to upload and delete files in the wiki.
			Note that the group 'everybody' is a special group that includes all registered users.
		</p>
		<div class="form-group">
			<input type="text" class="form-control" name="group" value="{{ .FileMasterGroup }}" placeholder="Group name">
		</div>
		<input type="submit" class="btn btn-default" value="Update">
	</form>
</div>

<div class="form-container">
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
	<form action="/admin/illegalnames" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<h3>Illegal names</h3>
		<div class="form-group">
			Sample illegal name list: <pre>{{ .DefaultIllegalNames }}</pre>
		</div>
		<div class="form-group"><textarea class="form-control" rows="15" name="illegal_names">{{ .illegal_names }}</textarea></div>
		<input type="submit" class="btn btn-default" value="Update Illegal names">
	</form>
</div>

{{ end }}
`

