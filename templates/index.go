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
			<input type="submit" value="Search">
		</form>
	</div>
	<div id="section-tabs-right">
		<span><a href="">Source</a></span>
		<span class="active"><a href="">Read</a></span>
	</div>
	<div id="section-tabs-left">
		<span class="active"><a href="">Main</a></span>
		<span><a href="">Discussion</a></span>
	</div>
</div>
<div id="meat">
	FemtoWiki content
</div>
{{ end }}`