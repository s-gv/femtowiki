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
html, body, #container {
	height: 100%;
	margin: 0;
	padding: 0;
}
#container {
	height: auto;
	min-height: 100%;
}
#main {
	padding-bottom: 3em;
}
.clearfix {
	overflow: auto;
	zoom: 1;
}
#header {
	padding: 10px;
	text-align: center;
}
#footer {
	position: relative;
	clear: both;
	height: 3em;
	margin-top: -3em;
	text-align: center;
}
#nav {
	float: left;
	max-width: 180px;
}
#content {
	margin-left: 200px;
}
`