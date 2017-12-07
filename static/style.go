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
	padding: 18px;
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
	padding: 10px 0px;
	font-size: 14px;
	margin-bottom: 16px;
}
#section-profile a {
	padding: 0 10px;
}
#footer {
	font-size: 12px;
}
#footer a {
	margin: 0 4px;
}
@media screen and (min-width:600px) {
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
@media screen and (min-width:600px) {
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

}
.nav-section ul {
	padding: 0;
	list-style-type: none;
}

/* search section */
#section-search {
	font-size: 14px;
	padding: 4px 0;
}
#section-search span.right {
	float: right;
	position: relative;
	top: -12px;
}
#section-search span {
	padding: 8px 12px;
	border: 1px none #ccc;
}
#section-search span.active {
	background-color: white;
	border-style: solid;
	border-bottom: none;
}


/* meat of the page */
#meat {
	min-height: 480px;
	margin-bottom: 16px;
	padding: 16px;
	background-color: white;
	border: 1px solid #ccc;
}

`