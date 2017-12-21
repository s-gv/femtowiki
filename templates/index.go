// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const indexSrc = `
{{ define "content" }}
<div id="section-tabs">
	<div id="section-search">
		<form method="GET" action="">
			<input type="text" name="query" placeholder="Search Femtowiki">
			<input class="btn btn-default" type="submit" value="Search">
		</form>
	</div>
	<div id="section-tabs-right">
		{{ if .IsEditMode }}
		<span class="active"><a href="{{ .EditURL }}">Source</a></span>
		<span><a href="{{ .URL }}">Read</a></span>
		{{ else }}
		<span><a href="{{ .EditURL }}">Source</a></span>
		<span class="active"><a href="{{ .URL }}">Read</a></span>
		{{ end }}
	</div>
	<div id="section-tabs-left">
		<span class="active"><a href="">Main</a></span>
		<span><a href="">Discussion</a></span>
	</div>
</div>
<div id="meat">
{{ if .IsEditMode }}
	<form action="/editpage" method="POST">
		<input type="hidden" name="csrf" value="{{ .ctx.CSRFToken }}">
		<input type="hidden" name="t" value="{{ .Title }}">
		<textarea rows="50" name="content">{{ .Content }}</textarea>
		<input type="submit" class="btn btn-default" name="action" value="Update">
		{{ if .IsCRUDGroupMember }}
		<input type="submit" class="btn btn-danger" name="action" value="Delete" onclick="return confirm('Are you sure you want to delete this page?');">
		{{ end }}
	</form>
{{ else }}
	<h1>{{ .Title }}</h1>

	<div class="toc">
	<ol>
	<li><a href="">Truncate the results</a></li>
	<li>
		<a href="">Legal</a>
		<ul>
			<li><a href="">Jane vs. Roe</a></li>
			<li><a href="">Jon vs. Jane</a></li>
			<li><a href="">Wade vs. Roe</a></li>
		</ul>
	</li>
	<li><a href="">Variations</a></li>
	</ol>
	</div>
	{{ .Content }}
{{ end }}
</div>
{{ end }}`