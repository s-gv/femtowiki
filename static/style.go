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
.clearfix {
	overflow: auto;
	zoom: 1;
}
#container {
	position: relative;
}
#header {
	text-align: right;
	padding: 10px;
}
#content {
	padding: 10px;
}
#nav {
	padding: 10px;
}
#footer {
	padding: 10px;
}
@media screen and (min-width:600px) {
	#nav {
		position: absolute;
		top: 0px;
		width: 180px;
	}
	#content {
		margin-left: 200px;
		min-height: 480px;
	}
	#footer {
		margin-left: 200px;
	}
}
`