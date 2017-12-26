// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const searchSrc = `
{{ define "content" }}
<div style="padding: 10px;">
	<div style="width: 100%; max-width: 480px; margin-bottom: 25px;">
		<form class="form-inline" method="GET" action="/search">
			<input class="form-control" style="width: 70%;" type="text" name="q" placeholder="Search Femtowiki" value="{{ .SearchTerms }}">
			<input class="btn btn-inline btn-default" style="width: 20%;" type="submit" value="Search">
		</form>
	</div>
	{{ if .Results }}
	{{ range .Results }}
		<div style="margin-top: 15px;">
			<div style="font-size: 18px;"><a href="/pages/{{ .CTitle }}">{{ .Title }}</a></div>
			<div>{{ .Snippet }}</div>
		</div>
	{{ end }}
	{{ else }}
	<p>No pages found</p>
	{{ end }}
</div>
{{ end }}
`