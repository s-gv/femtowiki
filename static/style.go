// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package static

const StyleSrc = `
* {
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
}
html, body {
	margin: 0;
	padding: 0;
}
body {
	font-family: Arial, "Helvetica Neue", Helvetica, sans-serif;
	-webkit-font-smoothing: antialiased;
	text-rendering: optimizeLegibility;
	line-height: 1.58;
	margin-bottom: 50px;
	background-color: #f6f6f6;
}
.clearfix {
	overflow: auto;
	zoom: 1;
}
a {
	text-decoration: none;
}
a:link {
	color: #07C;
}
a:hover, a:active {
	color: #3af;
}
a:visited {
	color: #005999;
}
#container {
	position: relative;
}
#header {
	text-align: right;
	padding: 18px 6px;
	background-color: #333333;
	border-bottom: 5px #08C solid;
}
#header .logo {
	float: left;
}
#header a {
	color: #FFFFFF;
	font-size: 16px;
	padding: 0 10px;
	font-weight: bold;
}
#header a:link, #header a:visited, #header a:hover, #header a:active {
	text-decoration: none;
}
#section-profile {
	text-align: right;
	padding: 10px 10px;
	font-size: 14px;
	margin-bottom: 16px;
}
.section-profile-link {
	padding: 0 10px;
}
#footer {
	font-size: 12px;
}
#footer a {
	margin: 0 4px;
}
@media screen and (min-width:800px) {
	#header a {
		padding: 0 20px;
	}
	#content, #footer {
		margin-left: 220px;
	}
}

/* navbar to the left */
#nav {
	margin: 16px 0;
	font-size: 14px;
}
@media screen and (min-width:800px) {
	#nav {
		position: absolute;
		top: 0px;
		width: 200px;
		margin: 16px 12px;
	}
}
.nav-section {
	padding: 8px 16px;
}
.nav-section h3 {
	margin: 0;
	font-weight: normal;
	font-size: 1em;
}
.nav-section hr {
	height: 1px;
	color: #ddd;
	background-color: #ddd;
	border: none;
	margin: 2px 0;
}
.nav-section ul {
	padding: 0;
	list-style-type: none;
	margin-top: 6px;
	margin-bottom: 12px;
}

/* search section */
#section-tabs {
	font-size: 14px;
	padding: 4px 0;
	position: relative;
	bottom: -1px;
	z-index: 10;
}
#section-search {
	margin-bottom: 12px;
	margin-left: 4px;
	margin-right: 4px;
}
#section-search input[type="text"] {
	padding: 2px;
	border: 1px solid #ccc;
}
#section-search input[type="submit"] {
	padding: 3px 8px;
}
#section-search input[type="text"] {
	width: 75%;
}
#section-search input[type="submit"] {
	width: 22%;
}
@media screen and (min-width:800px) {
	#section-search {
		float: right;
	}
	#section-search input[type="text"] {
		width: 240px;
	}
	#section-search input[type="submit"] {
		width: auto;
	}
}
#section-tabs-right {
	float: right;
}
#section-tabs-left {
}
#section-tabs span {
	padding: 8px 12px;
	border: 1px none #ccc;
}
#section-tabs span.active {
	background-color: white;
	border-style: solid;
	border-bottom: none;
}

/* meat of the page */
#meat {
	position: relative;
	min-height: 480px;
	margin-bottom: 16px;
	padding: 16px;
	background-color: white;
	border: 1px solid #ccc;
}
@media screen and (min-width:800px) {
	#meat {
		padding: 16px 24px;
	}
}
#meat h1, #meat h2 {
	font-family: 'Linux Libertine', 'Georgia', 'Times', serif;
	font-weight: normal;
	border-style: none none solid none;
	border-color: #ccc;
	border-width: 1px;
}
#meat textarea {
	width: 100%;
}
@media screen and (min-width:800px) {
	#meat input[type="submit"] {
		width: auto;
	}
}

/* table of contents box */
.toc {
	background-color: #f6f6f6;
	width: 100%;
	border: 1px solid #ccc;
}
@media screen and (min-width:800px) {
	.toc {
		width: 360px;
	}
}

/* forms */
.form-container {
	padding: 20px 8px;
	width: 100%;
}
@media screen and (min-width:800px) {
	.form-container {
		width: 360px;
	}
}
.form-group {
	margin: 8px 0px;
}
.form-group a, .form-group span {
	font-size: 14px;
}
input[type="text"].form-control, input[type="password"].form-control, input[type="email"].form-control {
	padding: 5px;
	border: 1px solid #ccc;
	border-radius: 4px;
	width: 100%;
}
textarea:focus, input[type="text"]:focus, input[type="password"]:focus, input[type="email"]:focus {
	outline: #0099E6 auto 1px;
	outline-offset: -2px;
}
.btn {
	padding: 6px 12px;
	border-radius: 4px;
	border: none;
	width: 100%;
}
.btn-default {
	background-color: #08c;
	color: #fff;
}
.btn-default:hover {
	background-color: #0099E6;
	cursor: pointer;
}
.btn-default:active {
	background-color: #006699;
	cursor: pointer;
}

/* flash message */
.flash {
	color: red;
}
`