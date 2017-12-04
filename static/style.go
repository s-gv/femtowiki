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
.clearfix {
	overflow: auto;
	zoom: 1;
}
html, body {
	margin: 0;
	padding: 0;
}
#wrap {
	position: relative;
}
#header {
	padding: 10px;
	text-align: center;
}
@media screen and (min-width:600px) {
	#nav {
		position: absolute;
		width: 180px;
	}
	#container {
		position: absolute;
		margin-left: 200px;
	}
	#content {
		min-height: 480px;
	}
}
`