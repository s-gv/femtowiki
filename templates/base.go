// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const baseSrc = `<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" type="text/css" href="/static/css/femtowiki.css?v=010">
	<title>
		Femtowiki
	</title>
	{{ block "head" . }}{{ end }}
</head>
<body>
	<div id="header">
		<div class="logo"><a href="/">Femtowiki</a></div> <a href="">Home</a> <a href="">Download</a> <a href="">FAQ</a>
	</div>
	<div id="container">
		<div id="content">
			<div id="section-profile"><a href="/signup">Signup</a> <a href="/login">Login</a></div>
			<div id="section-tabs">
				<div id="section-search">
					<form method="GET" action="">
						<input type="text" name="query" placeholder="Search">
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
				{{ block "content" . }}{{ end }}
			</div>
		</div>
		<div id="nav">
			<div class="nav-section">
				<ul>
					<li><a href="">Home</a></li>
					<li><a href="">Download</a></li>
					<li><a href="">FAQ</a></li>
				</ul>
			</div>
			<div class="nav-section">
				<h3>Cities</h3>
				<hr>
				<ul>
					<li><a href="">London</a></li>
					<li><a href="">New York</a></li>
					<li><a href="">Bangalore</a></li>
				</ul>
			</div>
			<div class="nav-section">
				<h3>Languages</h3>
				<hr>
				<ul>
					<li><a href="">English</a></li>
					<li><a href="">Hindi</a></li>
					<li><a href="">Kannada</a></li>
				</ul>
			</div>
		</div>
		<div id="footer">
			<a href="">Privacy</a> <a href="">Terms</a> <a href="">Help</a>
		</div>
	</div>
	<script src="/static/js/femtowiki.js?v=010"></script>
</body>
</html>`