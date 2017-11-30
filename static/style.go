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
	height: 100%;
}
.clearfix {
	overflow: auto;
	zoom: 1;
}
header {
	padding: 10px;
	text-align: center;
}
footer {
	text-align: center;
}
nav {
	float: left;
	max-width: 180px;
}
article {
	margin-left: 200px;
}
.container {
}
`